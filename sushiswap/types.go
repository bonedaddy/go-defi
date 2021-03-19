package sushiswap

import "math/big"

// Reserve represents a given sushiswap pair reserve
type Reserve struct {
	Reserve0           *big.Int // Token
	Reserve1           *big.Int // QuoteToken
	BlockTimestampLast uint32
}
