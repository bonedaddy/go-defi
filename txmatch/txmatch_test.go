package txmatch

import (
	"context"
	"encoding/hex"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/bonedaddy/go-defi/bclient"
	"github.com/bonedaddy/go-defi/bindings/erc20"
	"github.com/bonedaddy/go-defi/config"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

var (

	// NOTE: these transactions and addresses dont imply anything about fitness or correctness of the contracts
	// they are simply here for testing purposes. dont interact with any contracts included unless thoroughly researched
	// encodedTx is a hex encoded marshalled transaction
	encodedTx    = "f8aa04851d91ca360083012d08945ade7ae8660293f2ebfcefaba91d141d72d221e880b844a9059cbb0000000000000000000000001cdb2c26a767ba197fdbbd9fa3488a6032811a7b00000000000000000000000000000000000000000001a84079ae6836e61e8e6625a0c488c44dcd992a3aa9dd9d3e9211b4cd55132c226fe6f743c729a07671a917c5a00ca637f1e43190a2bdff3dbdeec4d2151bf4941a3b4826daca28cdb0a7087499"
	contractAddr = "0x5ade7ae8660293f2ebfcefaba91d141d72d221e8"
)

func TestTxMatch(t *testing.T) {
	t.Cleanup(func() { os.Remove("wantfile.txt") })
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_cfg := *config.ExampleConfig
	cfg := &_cfg
	cfg.ProviderType = "infura"
	cfg.APIKey = os.Getenv("INFURA_API_KEY")
	logger, err := cfg.ZapLogger()
	require.NoError(t, err)
	ec, err := cfg.EthClient(ctx)
	require.NoError(t, err)
	bc, err := bclient.NewClient(ctx, ec)
	require.NoError(t, err)
	parsedABI, err := abi.JSON(strings.NewReader(erc20.Erc20ABI))
	require.NoError(t, err)
	decodedTx, err := hex.DecodeString(encodedTx)
	require.NoError(t, err)
	tx := new(types.Transaction)
	require.NoError(t, tx.UnmarshalBinary(decodedTx))
	_ = parsedABI
	wantMethods := []string{"transfer", "transferFrom"}
	matcher, err := NewMatcher(logger, bc, erc20.Erc20ABI, wantMethods, []string{contractAddr})
	require.NoError(t, err)
	startBlock := uint64(12082591)
	endBlock := uint64(12082599)
	require.NoError(t, matcher.Match("wantfile.txt", startBlock, endBlock))
	data, err := ioutil.ReadFile("wantfile.txt")
	require.NoError(t, err)
	require.Equal(t, string(data), "0xA96F2A820D3d7d5Df3C6c638d40C295a0CF255D1\n")
}
