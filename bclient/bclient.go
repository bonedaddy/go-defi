package bclient

import (
	"context"

	"github.com/bonedaddy/go-defi/sushiswap"
	"github.com/bonedaddy/go-defi/uniswap"
	"github.com/bonedaddy/go-defi/utils"
	"github.com/ethereum/go-ethereum/ethclient"
)

// BClient wraps ethclient and provides helper functions for commonly used functionality
type BClient struct {
	ctx    context.Context
	cancel context.CancelFunc
	bc     utils.Blockchain
	uc     *uniswap.Client
	sc     *sushiswap.Client
}

// New returns a new bclient implementation without any authentication
// as well as wrappers around the uniswap and sushiswap clients. All underlying
// types are exposed in addition to helper functions
func NewClient(ctx context.Context, bc utils.Blockchain) (*BClient, error) {
	ctx, cancel := context.WithCancel(ctx)

	return &BClient{
		ctx:    ctx,
		cancel: cancel,
		bc:     bc,
		uc:     uniswap.NewClient(bc),
		sc:     sushiswap.NewClient(bc),
	}, nil
}

// EthClient attempts to conver the Blockchain interface to an ethclient type
// returning an error if unable to type convert the interface. This likely indicates
// that a simulated backend is being used
func (bc *BClient) EthClient() (*ethclient.Client, error) {
	ec, ok := bc.bc.(*ethclient.Client)
	if !ok {
		return nil, ErrNotEthClient
	}
	return ec, nil
}

// CurrentBlock returns the current block known by the ethereum client
func (bc *BClient) CurrentBlock() (uint64, error) {
	ec, err := bc.EthClient()
	if err != nil {
		return 0, err
	}
	return ec.BlockNumber(bc.ctx)
}

// Sushiswap returns the underlying sushiswap client
func (bc *BClient) Sushiswap() *sushiswap.Client { return bc.sc }

// Uniswap returns the underlying uniswap client
func (bc *BClient) Uniswap() *uniswap.Client { return bc.uc }

func (bc *BClient) Context() context.Context { return bc.ctx }

// Close terminates the blockchain connection
func (bc *BClient) Close() {
	// error is nil which means we can close the client
	if ec, err := bc.EthClient(); err == nil {
		ec.Close()
	}
}
