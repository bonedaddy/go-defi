package common

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func SortAddressess(tkn0, tkn1 common.Address) (common.Address, common.Address) {
	token0Rep := big.NewInt(0).SetBytes(tkn0.Bytes())
	token1Rep := big.NewInt(0).SetBytes(tkn1.Bytes())

	if token0Rep.Cmp(token1Rep) > 0 {
		tkn0, tkn1 = tkn1, tkn0
	}

	return tkn0, tkn1
}
