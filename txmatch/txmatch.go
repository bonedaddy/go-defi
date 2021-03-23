package txmatch

import (
	"math/big"
	"os"
	"strings"

	"github.com/bonedaddy/go-defi/bclient"
	"github.com/bonedaddy/go-defi/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

// Matcher provides transaction matching capabilities allowing for the filtration
// of transactions to specific contracts or addresses using ABI's as the method
// of filtering out transactions. Suitable for building a list of addresses
// that have interacted with particular contracts
type Matcher struct {
	l             *zap.Logger
	bc            *bclient.BClient
	wantMethods   map[string]bool
	wantContracts map[common.Address]bool
	abi           abi.ABI
}

// NewMatcher returns a new transaction matcher
func NewMatcher(
	logger *zap.Logger,
	bc *bclient.BClient,
	abistr string,
	methods,
	contracts []string,
) (*Matcher, error) {
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
		l:             logger.Named("txmatch"),
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
		if utils.LogContextDone(m.l, m.bc.Context()) {
			m.l.Info("exiting at block number", zap.Uint64("block.number", i))
			return nil
		}
		ib := new(big.Int).SetUint64(i)
		block, err := ec.BlockByNumber(m.bc.Context(), ib)
		if err != nil {
			m.l.Error("ecountered error fetching block by number", zap.Error(err), zap.Uint64("block.number", i))
			continue
		}
		for _, tx := range block.Transactions() {
			if utils.LogContextDone(m.l, m.bc.Context()) {
				m.l.Info("exiting at block number", zap.Uint64("block.number", i))
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
			// for some reason when using common.Address in a map
			// the al zero adddress will be marked as true
			// and pass the check above m.wantContracts so explicitly ignore it
			// if tx.To().String() == common.HexToAddress("").String() {
			//	continue
			//}
			if len(tx.Data()) < 3 {
				// skip as this is not a contract call we are interested in
				continue
			}
			m.l.Info("found interested transaction", zap.String("tx.hash", tx.Hash().String()))
			// validate the method is as expected
			method, err := m.abi.MethodById(tx.Data()[0:4])
			if err != nil {
				// skip it
				continue
			}
			if !m.wantMethods[method.Name] {
				// skip
				continue
			}
			msg, err := tx.AsMessage(types.NewEIP2930Signer(chid))
			if err != nil {
				m.l.Error("failed to parse tx as message", zap.Error(err), zap.String("tx.hash", tx.Hash().String()))
				continue
			}
			_, err = fh.WriteString(msg.From().String() + "\n")
			if err != nil {
				m.l.Error("failed to append address to file", zap.Error(err))
			}
		}
	}
	return nil
}
