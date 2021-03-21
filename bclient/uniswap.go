package bclient

import (
	"math/big"
)

// EthDaiPrice returns the price of ETH in terms of DAI
func (bc *BClient) EthDaiPrice() (*big.Int, error) {
	reserves, err := bc.uc.GetReserves(WETHTokenAddress, DAITokenAddress)
	if err != nil {
		return nil, err
	}
	return new(big.Int).Div(reserves.Reserve1, reserves.Reserve0), nil
}
