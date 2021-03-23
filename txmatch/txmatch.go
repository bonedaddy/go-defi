package txmatch

import (
	"log"
	"math/big"
	"strings"

	"github.com/bonedaddy/go-defi/bclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
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

func (m *Matcher) Match(startBlock, endBlock uint64) error {
	ec, err := m.bc.EthClient()
	if err != nil {
		return err
	}
	for i := startBlock; i <= endBlock; i++ {
		ib := new(big.Int).SetUint64(i)
		block, err := ec.BlockByNumber(m.bc.Context(), ib)
		if err != nil {
			log.Println("encountered error fetching block by number: ", block)
			continue
		}
		for _, tx := range block.Transactions() {
			// first check it is to one of the contracts
			if tx.To() == nil {
				// skip this as its a contract deployment transaction
				continue
			}
			if !m.wantContracts[*tx.To()] {
				// skip this as its not to a contract we are interested in
				continue
			}
			log.Println("found interested transaction: ", tx.Hash())
		}
	}
	return nil
}
