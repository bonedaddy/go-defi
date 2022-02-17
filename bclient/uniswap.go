package bclient

import (
	"math/big"

	"github.com/pkg/errors"
)

// EthDaiPrice returns the price of ETH in terms of DAI
func (bc *BClient) EthDaiPrice() (*big.Int, error) {
	reserves, err := bc.uc.GetReserves(WETHTokenAddress, DAITokenAddress)
	if err != nil {
		return nil, errors.Wrap(err, "get reserves")
	}
	return new(big.Int).Div(reserves.Reserve1, reserves.Reserve0), nil
}
