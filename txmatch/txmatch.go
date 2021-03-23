package txmatch

import (
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/bonedaddy/go-defi/bclient"
	"github.com/bonedaddy/go-defi/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Matcher struct {
	bc            *bclient.BClient
	wantMethods   map[string]bool
	wantContracts map[common.Address]bool
	abi           abi.ABI
}

// NewMatcher returns a new transaction matcher
func NewMatcher(bc *bclient.BClient, abistr string, methods, contracts []string) (*Matcher, error) {
	parsedABI, err := abi.JSON(strings.NewReader(abistr))
	if err != nil {
		return nil, err
	}
	var (
		wantMethods   = make(map[string]bool)
		wantContracts = make(map[common.Address]bool)
	)
	for _, method := range methods {
		wantMethods[method] = true
	}
	for _, contract := range contracts {
		wantContracts[common.HexToAddress(contract)] = true
	}
	return &Matcher{
		abi:           parsedABI,
		bc:            bc,
		wantMethods:   wantMethods,
		wantContracts: wantContracts,
	}, nil
}

func (m *Matcher) Match(outFile string, startBlock, endBlock uint64) error {
	fh, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer fh.Close()
	chid, err := m.bc.ChainID()
	if err != nil {
		return err
	}
	ec, err := m.bc.EthClient()
	if err != nil {
		return err
	}
	for i := startBlock; i <= endBlock; i++ {
		if utils.ContextDone(m.bc.Context()) {
			log.Println("context cancelled, exiting")
			return nil
		}
		ib := new(big.Int).SetUint64(i)
		block, err := ec.BlockByNumber(m.bc.Context(), ib)
		if err != nil {
			log.Println("encountered error fetching block by number: ", block)
			continue
		}
		for _, tx := range block.Transactions() {
			if utils.ContextDone(m.bc.Context()) {
				log.Println("context cancelled, exiting")
				return nil
			}
			// first check it is to one of the contracts
			if tx.To() == nil {
				// skip this as its a contract deployment transaction
				continue
			}
			// validate it is to an interested contract
			if !m.wantContracts[*tx.To()] {
				// skip this as its not to a contract we are interested in
				continue
			}
			log.Println("found interested transaction: ", tx.Hash())
			// validate the method is as expected
			method, err := m.abi.MethodById(tx.Data()[0:4])
			if err != nil {
				// skip it
				continue
			}
			if !m.wantMethods[method.Name] {
				log.Println("Invalid method name")
				// skip
				continue
			}
			msg, err := tx.AsMessage(types.NewEIP2930Signer(chid))
			if err != nil {
				log.Println("failed to parse tx as message: ", err)
				continue
			}
			_, err = fh.WriteString(msg.From().String() + "\n")
			if err != nil {
				log.Println("failed to append address to file: ", err)
			}
		}
	}
	return nil
}
