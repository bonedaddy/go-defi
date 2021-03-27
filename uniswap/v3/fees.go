package uniswapv3

import (
	"errors"
	"math/big"
)

// Fee is a typed int64 which denotes the supported
// fee settings paid by swappers denoted in hundredths of a basis point (0.0001%)
//
// Supprted feeds are 500 (0.05%), 3000 (0.3%), 10000 (1.0%)
type Fee int64

// Big returns Fee as type *big.Int
func (f Fee) Big() *big.Int {
	return big.NewInt(int64(f))
}

// Percent returns the fee percentage as a float64 type
func (f Fee) Percent() float64 {
	switch f {
	case FeePointZeroFivePercent:
		return 0.05
	case FeePointThreePercent:
		return 0.3
	case FeeOnePercent:
		return 1.0
	case FeeZero:
		fallthrough
	default:
		return 0
	}
}

var (
	// FeePointZeroFivePercent denotes a fee of 0.05%
	FeePointZeroFivePercent = Fee(500)
	// FeePointThreePercent denotes a fee of 0.3%
	FeePointThreePercent = Fee(3000)
	// FeeOnePercent denotes a fee of 1.0%
	FeeOnePercent = Fee(10000)
	// FeeZero denotes a fee of 0% which is currently unsupported
	// and is the fee returned during errors
	FeeZero = Fee(0)

	// fees as reported by https://github.com/Uniswap/uniswap-v3-sdk/blob/main/src/constants.ts#L9-L13
	// LowFee see constants linke above
	LowFee = FeePointZeroFivePercent
	// MediumFee see constants link above
	MediumFee = FeePointThreePercent
	// HighFee see constnats link above
	HighFee = FeeOnePercent

	// ErrUnsupportedFee is an error returned when attempting to use an unsupported fee setting
	ErrUnsupportedFee = errors.New("unsupported fee setting")
)

// GetFee returns a Fee type from the desired percent
// returning an error if the fee is unsupported
func GetFee(percent float64) (Fee, error) {
	switch percent {
	case 0.05:
		return FeePointZeroFivePercent, nil
	case 0.3:
		return FeePointThreePercent, nil
	case 1.0:
		return FeeOnePercent, nil
	default:
		return FeeZero, ErrUnsupportedFee
	}
}
