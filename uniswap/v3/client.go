package v3

import (
	"context"

	v3factory "github.com/bonedaddy/go-defi/bindings/uniswap/v3/factory"
	"github.com/ethereum/go-ethereum/common"

	"github.com/bonedaddy/go-defi/utils"
)

type Client struct {
	ctx    context.Context
	cancel context.CancelFunc
	bc     utils.Blockchain
}

// New returns a new uniswap v3 client
func New(ctx context.Context, bc utils.Blockchain) *Client {
	ctx, cancel := context.WithCancel(ctx)
	return &Client{ctx: ctx, cancel: cancel, bc: bc}
}

// Close terminates the context associated with this client
func (c *Client) Close() {
	c.cancel()
}

// Factory returns a uniswapv3 factory contract binding
func (c *Client) Factory(addr string) (*v3factory.V3factory, error) {
	return v3factory.NewV3factory(common.HexToAddress(addr), c.bc)
}
