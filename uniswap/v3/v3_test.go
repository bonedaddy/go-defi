package v3

import (
	"context"
	"testing"

	v3factory "github.com/bonedaddy/go-defi/bindings/uniswap/v3/factory"
	"github.com/bonedaddy/go-defi/testenv"
	"github.com/stretchr/testify/require"
)

func TestV3Simulated(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tenv, err := testenv.NewBlockchain(ctx)
	require.NoError(t, err)
	var factory *v3factory.V3factory
	t.Run("DeployFactory", func(t *testing.T) {
		addr, tx, pFactory, err := v3factory.DeployV3factory(tenv.Auth, tenv)
		require.NoError(t, err)
		_, err = tenv.DoWaitDeployed(tx)
		require.NoError(t, err)
		factory = pFactory
		t.Log("pool factory address: ", addr)
	})
	_ = factory
}
