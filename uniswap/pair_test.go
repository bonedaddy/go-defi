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

package uniswap

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestGeneratePairAddress(t *testing.T) {
	type args struct {
		token0 common.Address
		token1 common.Address
	}
	tests := []struct {
		name string
		args args
		want common.Address
	}{
		{
			name: "calculate address correctly if addresses sorted",
			args: args{
				token0: common.HexToAddress("0xc1C0472c0C80bCcDC7F5D01A376Bd97a734B8815"),
				token1: common.HexToAddress("0xEfe7FeA4ABf53Df50AbC9D2F7F2079911765E9D1"),
			},
			want: common.HexToAddress("0x45e7e04f70A3dCF8108dC90F67b82A565e492b9b"),
		},
		{
			name: "calculate address correctly if addresses are not sorted",
			args: args{
				token0: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
				token1: common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"),
			},
			want: common.HexToAddress("0xAE461cA67B15dc8dc81CE7615e0320dA1A9aB8D5"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GeneratePairAddress(tt.args.token0, tt.args.token1)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeneratePairAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortAddressess(t *testing.T) {
	type args struct {
		tkn0 common.Address
		tkn1 common.Address
	}
	tests := []struct {
		name  string
		args  args
		want  common.Address
		want1 common.Address
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := sortAddressess(tt.args.tkn0, tt.args.tkn1)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortAddressess() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("sortAddressess() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestQuote(t *testing.T) {
	type args struct {
		amount   *big.Int
		reserve0 func() *big.Int
		reserve1 func() *big.Int
	}
	tests := []struct {
		name string
		args args
		want func() *big.Int
	}{
		{
			name: "calculates the quote correctly 1",
			args: args{
				reserve0: func() *big.Int { return big.NewInt(20162514) },
				reserve1: func() *big.Int {
					res, _ := big.NewInt(0).SetString("496000000000000000000", 10)
					return res
				},
				amount: big.NewInt(1),
			},
			want: func() *big.Int { return big.NewInt(24600106911271) },
		},
		{
			name: "calculates the quote correctly 2",
			args: args{
				reserve1: func() *big.Int { return big.NewInt(20162514) },
				reserve0: func() *big.Int {
					res, _ := big.NewInt(0).SetString("496000000000000000000", 10)
					return res
				},
				amount: big.NewInt(1),
			},
			want: func() *big.Int { return big.NewInt(0) },
		},
		{
			name: "calculates the quote correctly 3",
			args: args{
				reserve1: func() *big.Int { return big.NewInt(20162514) },
				reserve0: func() *big.Int {
					res, _ := big.NewInt(0).SetString("496000000000000000000", 10)
					return res
				},
				amount: big.NewInt(10000000000000000),
			},
			want: func() *big.Int { return big.NewInt(406) },
		},
		{
			name: "calculates the quote correctly 4",
			args: args{
				reserve1: func() *big.Int {
					res, _ := big.NewInt(0).SetString("500000000000000000000", 10)
					return res
				},
				reserve0: func() *big.Int {
					res, _ := big.NewInt(0).SetString("20000000", 10)
					return res
				},
				amount: big.NewInt(10000000000000000),
			},
			want: func() *big.Int {
				res, _ := big.NewInt(0).SetString("250000000000000000000000000000", 10)
				return res
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Quote(tt.args.amount, tt.args.reserve0(), tt.args.reserve1()); !reflect.DeepEqual(got, tt.want()) {
				t.Errorf("Quote() = %v, want %v", got, tt.want())
			}
		})
	}
}

func TestGetPathPairs(t *testing.T) {
	type args struct {
		tokens []common.Address
	}
	tests := []struct {
		name string
		args args
		want []Pair
	}{
		{
			name: "no pairs if single address",
			args: args{
				tokens: []common.Address{common.HexToAddress("0x0")},
			},
			want: []Pair{},
		},
		{
			name: "no pairs no addresses",
			args: args{
				tokens: []common.Address{},
			},
			want: []Pair{},
		},
		{
			name: "single pair if two addresses",
			args: args{
				tokens: []common.Address{
					common.HexToAddress("0x0"),
					common.HexToAddress("0x1"),
				},
			},
			want: []Pair{
				{
					Token0: common.HexToAddress("0x0"),
					Token1: common.HexToAddress("0x1"),
				},
			},
		},
		{
			name: "two pair if three addresses",
			args: args{
				tokens: []common.Address{
					common.HexToAddress("0x0"),
					common.HexToAddress("0x1"),
					common.HexToAddress("0x2"),
				},
			},
			want: []Pair{
				{
					Token0: common.HexToAddress("0x0"),
					Token1: common.HexToAddress("0x1"),
				},
				{
					Token0: common.HexToAddress("0x1"),
					Token1: common.HexToAddress("0x2"),
				},
			},
		},
		{
			name: "three pair if four addresses",
			args: args{
				tokens: []common.Address{
					common.HexToAddress("0x0"),
					common.HexToAddress("0x1"),
					common.HexToAddress("0x2"),
					common.HexToAddress("0x3"),
				},
			},
			want: []Pair{
				{
					Token0: common.HexToAddress("0x0"),
					Token1: common.HexToAddress("0x1"),
				},
				{
					Token0: common.HexToAddress("0x1"),
					Token1: common.HexToAddress("0x2"),
				},
				{
					Token0: common.HexToAddress("0x2"),
					Token1: common.HexToAddress("0x3"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPathPairs(tt.args.tokens); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPathPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
