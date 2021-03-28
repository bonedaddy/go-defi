package uniswapv3

import (
	"context"

	v3factory "github.com/bonedaddy/go-defi/bindings/uniswap/v3/factory"
	v3nftposmanager "github.com/bonedaddy/go-defi/bindings/uniswap/v3/nftposmanager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

// DeployNFTPositionManager deploys an pool nft position manager
func (c *Client) DeployNFTPositionManager(
	auth *bind.TransactOpts,
	factory,
	weth9,
	tokenDescriptor common.Address,
) (*v3nftposmanager.V3nftposmanager, common.Address, error) {
	addr, tx, contract, err := v3nftposmanager.DeployV3nftposmanager(
		auth,
		c.bc,
		factory,
		weth9,
		tokenDescriptor,
	)
	if err != nil {
		return nil, common.Address{}, err
	}
	_, err = bind.WaitDeployed(c.ctx, c.bc, tx)
	if err != nil {
		return nil, common.Address{}, err
	}
	return contract, addr, nil
}
