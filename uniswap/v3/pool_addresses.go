package v3

import (
	"encoding/hex"
	"math/big"

	ucommon "github.com/bonedaddy/go-defi/uniswap/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	// InitCodeHash is the init code's hash pulled from https://github.com/Uniswap/uniswap-v3-sdk/blob/main/src/constants.ts
	InitCodeHash = "01d4d358e07707f4db84b6a7527455b06f95ee89b5d059b4a1298ada7b6c7d67"
)

// GeneratePoolAddress is used to calculate a pool's address
// it leverages create 2 with token0 token1 and the fee being part of the salt
func GeneratePoolAddress(
	factory,
	token0,
	token1 common.Address,
	fee Fee,
) common.Address {
	token0, token1 = ucommon.SortAddressess(token0, token1)

	message := []byte{255}
	message = append(message, factory.Bytes()...)

	addrType, _ := abi.NewType("address", "address", nil)
	uint24Type, _ := abi.NewType("uint24", "uint24", nil)
	arguments := abi.Arguments{
		{Type: addrType},
		{Type: addrType},
		{Type: uint24Type},
	}
	saltEncoded, _ := arguments.Pack(
		token0, token1, fee.Big(),
	)
	saltHash := crypto.Keccak256(saltEncoded)

	message = append(message, saltHash...)

	b, _ := hex.DecodeString(InitCodeHash)
	message = append(message, b...)
	hashed := crypto.Keccak256(message)
	addressBytes := big.NewInt(0).SetBytes(hashed)
	addressBytes = addressBytes.Abs(addressBytes)
	return common.BytesToAddress(addressBytes.Bytes())
}
