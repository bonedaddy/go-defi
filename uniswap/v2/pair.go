// Copyright (C) 2021 Alexandre Trottier (Bonedaddy)
/* Mysterium network payment library.
 *
 * Copyright (C) 2020 BlockDev AG
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// Package uniswap contains the required bindings as well as corresponding wrappers for those bindings for interacting with the uniswap smart contracts.
package uniswap

import (
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// FactoryAddress points to the uniswap factory.
var FactoryAddress = common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")

// Router02Address points to the uniswap v2 02 router.
var Router02Address = common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")

const pairAddressSuffix = "96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f"

// GeneratePairAddress generates a pair address for the given tokens
func GeneratePairAddress(token0, token1 common.Address) common.Address {
	// addresses need to be sorted in an ascending order for proper behaviour
	token0, token1 = sortAddressess(token0, token1)

	// 255 is required as a prefix for this to work
	// see: https://uniswap.org/docs/v2/javascript-SDK/getting-pair-addresses/
	message := []byte{255}

	message = append(message, FactoryAddress.Bytes()...)

	addrSum := token0.Bytes()
	addrSum = append(addrSum, token1.Bytes()...)

	message = append(message, crypto.Keccak256(addrSum)...)

	b, _ := hex.DecodeString(pairAddressSuffix)
	message = append(message, b...)
	hashed := crypto.Keccak256(message)
	addressBytes := big.NewInt(0).SetBytes(hashed)
	addressBytes = addressBytes.Abs(addressBytes)
	return common.BytesToAddress(addressBytes.Bytes())
}

// Quote gets the exchange quote for a given amount of tokens and the respective pairs reserves.
func Quote(amount, reserve0, reserve1 *big.Int) *big.Int {
	if reserve1.Cmp(big.NewInt(0)) <= 0 ||
		reserve0.Cmp(big.NewInt(0)) <= 0 ||
		amount.Cmp(big.NewInt(0)) <= 0 {

		return new(big.Int)
	}

	// amountB = amountA.mul(reserveB) / reserveA;
	multiplied := new(big.Int).Mul(amount, reserve1)
	res := new(big.Int).Div(multiplied, reserve0)
	return res
}

func sortAddressess(tkn0, tkn1 common.Address) (common.Address, common.Address) {
	token0Rep := big.NewInt(0).SetBytes(tkn0.Bytes())
	token1Rep := big.NewInt(0).SetBytes(tkn1.Bytes())

	if token0Rep.Cmp(token1Rep) > 0 {
		tkn0, tkn1 = tkn1, tkn0
	}

	return tkn0, tkn1
}

// Pair represents a token pair.
type Pair struct {
	Token0, Token1 common.Address
}

// GetPathPairs takes in the given token path and returns the corresponding pairs.
func GetPathPairs(tokens []common.Address) []Pair {
	pairs := make([]Pair, 0)

	for i := 1; i < len(tokens); i++ {
		pair := Pair{
			Token0: tokens[i-1],
			Token1: tokens[i],
		}

		pairs = append(pairs, pair)
	}

	return pairs
}
