package testenv

import (
	"math/big"

	"github.com/bonedaddy/bdsm/bindings/erc20factory"
	v3factory "github.com/bonedaddy/go-defi/bindings/uniswap/v3/factory"
	v3swaprouter "github.com/bonedaddy/go-defi/bindings/uniswap/v3/swaprouter"
	"github.com/bonedaddy/go-defi/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// UniswapV3Testenv provides an abstraction overtop of Testenv
// that is specialized towards a uniswap v3 environment
type UniswapV3Testenv struct {
	*Testenv
	PoolFactory       *v3factory.V3factory
	SwapRouter        *v3swaprouter.V3swaprouter
	WETH9             common.Address
	FactoryAddress    common.Address
	SwapRouterAddress common.Address
}

func (tenv *Testenv) DeployUniswapV3() (*UniswapV3Testenv, error) {
	poolFactory, tx, factory, err := v3factory.DeployV3factory(tenv.Auth, tenv)
	if err != nil {
		return nil, err
	}
	_, err = tenv.DoWaitDeployed(tx)
	if err != nil {
		return nil, err
	}
	weth9, err := tenv.DeployERC20(tenv.Auth.From, utils.ToWei(100000000.0, 18), "WETH9", "WETH9", 18)
	if err != nil {
		return nil, err
	}
	utestenv := &UniswapV3Testenv{
		Testenv:        tenv,
		PoolFactory:    factory,
		FactoryAddress: poolFactory,
		WETH9:          weth9,
	}
	swapRouter, tx, router, err := v3swaprouter.DeployV3swaprouter(tenv.Auth, tenv, poolFactory, weth9)
	if err != nil {
		return nil, err
	}
	_, err = tenv.DoWaitDeployed(tx)
	if err != nil {
		return nil, err
	}
	utestenv.SwapRouter = router
	utestenv.SwapRouterAddress = swapRouter
	return utestenv, nil
}

// V3Factory deploys a uniswap v3 factory
func (tenv *Testenv) V3Factory() (*v3factory.V3factory, common.Address, error) {
	addr, tx, contract, err := v3factory.DeployV3factory(tenv.Auth, tenv)
	if err != nil {
		return nil, common.Address{}, err
	}
	_, err = bind.WaitDeployed(tenv.ctx, tenv, tx)
	if err != nil {
		return nil, common.Address{}, err
	}
	return contract, addr, nil
}

func (tenv *Testenv) DeployERC20(
	owner common.Address,
	supply *big.Int,
	name string,
	symbol string,
	decimals uint8,
) (common.Address, error) {
	_, tx, factory, err := erc20factory.DeployErc20factory(tenv.Auth, tenv)
	if err != nil {
		return common.Address{}, err
	}
	_, err = tenv.DoWaitDeployed(tx)
	if err != nil {
		return common.Address{}, err
	}
	tx, err = factory.DeployToken(
		tenv.Auth,
		owner,
		supply,
		decimals,
		name,
		symbol,
	)
	if err != nil {
		return common.Address{}, err
	}
	if err := tenv.DoWaitMined(tx); err != nil {
		return common.Address{}, err
	}
	return factory.SymbolToAddr(nil, symbol)
}
