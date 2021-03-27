package v3

import (
	"context"
	"math/big"
	"testing"

	"github.com/bonedaddy/bdsm/bindings/erc20factory"
	v3factory "github.com/bonedaddy/go-defi/bindings/uniswap/v3/factory"
	"github.com/bonedaddy/go-defi/testenv"
	"github.com/bonedaddy/go-defi/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestV3Simulated(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tenv, err := testenv.NewBlockchain(ctx)
	require.NoError(t, err)
	var poolFactory *v3factory.V3factory
	t.Run("DeployUniswapV3Factory", func(t *testing.T) {
		addr, tx, pFactory, err := v3factory.DeployV3factory(tenv.Auth, tenv)
		require.NoError(t, err)
		_, err = tenv.DoWaitDeployed(tx)
		require.NoError(t, err)
		poolFactory = pFactory
		t.Log("pool factory address: ", addr)
	})
	var erc20Factory *erc20factory.Erc20factory
	t.Run("DeployERC20Factory", func(t *testing.T) {
		addr, tx, eFactory, err := erc20factory.DeployErc20factory(tenv.Auth, tenv)
		require.NoError(t, err)
		_, err = tenv.DoWaitDeployed(tx)
		require.NoError(t, err)
		erc20Factory = eFactory
		t.Log("erc20 factory address: ", addr)
	})
	// token symbols -> addresses
	tokenMap := make(map[string]common.Address)
	var tokenA1, tokenB1, tokenA2, tokenB2 common.Address
	t.Run("CreatingFakeTokens", func(t *testing.T) {
		type args struct {
			owner    common.Address
			supply   *big.Int
			name     string
			symbol   string
			decimals uint8
		}
		// mint tokens, with 100M supply of each
		tests := []struct {
			name string
			args args
		}{
			{"TokenA1", args{tenv.Auth.From, utils.ToWei(100000000.0, 18), "TokenA1", "TA1", 18}},
			{"TokenB1", args{tenv.Auth.From, utils.ToWei(100000000.0, 18), "TokenB1", "TB1", 18}},
			{"TokenA2", args{tenv.Auth.From, utils.ToWei(100000000.0, 18), "TokenA2", "TA2", 18}},
			{"TokenB2", args{tenv.Auth.From, utils.ToWei(100000000.0, 18), "TokenB2", "TB2", 18}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tx, err := erc20Factory.DeployToken(
					tenv.Auth,
					tt.args.owner,
					tt.args.supply,
					tt.args.decimals,
					tt.args.name,
					tt.args.symbol,
				)
				require.NoError(t, err)
				require.NoError(t, tenv.DoWaitMined(tx))
				tokenMap[tt.args.symbol], err = erc20Factory.SymbolToAddr(nil, tt.args.symbol)
				require.NoError(t, err)
				require.NotEqual(t, tokenMap[tt.args.symbol], common.HexToAddress(""))
				t.Logf("%s address %s\n", tt.args.symbol, tokenMap[tt.args.symbol])
				switch tt.args.symbol {
				case "TA1":
					tokenA1 = tokenMap[tt.args.symbol]
				case "TB1":
					tokenB1 = tokenMap[tt.args.symbol]
				case "TA2":
					tokenA2 = tokenMap[tt.args.symbol]
				case "TB2":
					tokenB2 = tokenMap[tt.args.symbol]
				}
			})
		}
	})
	// for naming purposes a pool is identified by:
	// token0:token1-fee
	_ = poolFactory
	_ = erc20Factory
	_ = tokenA1
	_ = tokenB1
	_ = tokenA2
	_ = tokenB2
}
