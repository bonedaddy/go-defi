package bclient

import (
	"context"

	"github.com/bonedaddy/go-defi/config"
	"github.com/bonedaddy/go-defi/sushiswap"
	"github.com/bonedaddy/go-indexed/uniswap"
	"github.com/ethereum/go-ethereum/ethclient"
)

// BClient wraps ethclient and provides helper functions for commonly used functionality
type BClient struct {
	ctx    context.Context
	cancel context.CancelFunc
	ec     *ethclient.Client
	uc     *uniswap.Client
	sc     *sushiswap.Client
}

// New returns a new bclient implementation without any authentication
// as well as wrappers around the uniswap and sushiswap clients. All underlying
// types are exposed in addition to helper functions
func NewClient(ctx context.Context, cfg *config.Blockchain) (*BClient, error) {
	ctx, cancel := context.WithCancel(ctx)
	var (
		ec  *ethclient.Client
		err error
	)
	switch cfg.ProviderType {
	case "infura":
		transport := ""
		path := ""
		if cfg.Websockets {
			transport = "wss://"
			path = "/ws/v3"
		} else {
			transport = "http://"
			path = "/v3"
		}
		url := transport + cfg.Network + ".infura.io" + path + "/" + cfg.APIKey
		ec, err = ethclient.DialContext(ctx, url)
	case "direct":
		ec, err = ethclient.DialContext(ctx, cfg.RPC)
	}
	if err != nil {
		cancel()
		return nil, err
	}
	return &BClient{
		ctx:    ctx,
		cancel: cancel,
		ec:     ec,
		uc:     uniswap.NewClient(ec),
		sc:     sushiswap.NewClient(ec),
	}, nil
}

// CurrentBlock returns the current block known by the ethereum client
func (bc *BClient) CurrentBlock() (uint64, error) { return bc.ec.BlockNumber(bc.ctx) }

// EthClient returns the underlying ethclient
func (bc *BClient) EthClient() *ethclient.Client { return bc.ec }

// Sushiswap returns the underlying sushiswap client
func (bc *BClient) Sushiswap() *sushiswap.Client { return bc.sc }

// Uniswap returns the underlying uniswap client
func (bc *BClient) Uniswap() *uniswap.Client { return bc.uc }

func (bc *BClient) Context() context.Context { return bc.ctx }

// Close terminates the blockchain connection
func (bc *BClient) Close() {
	bc.ec.Close()
}
