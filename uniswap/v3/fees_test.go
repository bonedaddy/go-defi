package uniswapv3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFees(t *testing.T) {
	type args struct {
		percent float64
	}
	tests := []struct {
		name    string
		args    args
		wantFee Fee
		wantErr error
	}{
		{"0.05", args{0.05}, FeePointZeroFivePercent, nil},
		{"0.3", args{0.3}, FeePointThreePercent, nil},
		{"1.0", args{1.0}, FeeOnePercent, nil},
		{"1", args{1}, FeeOnePercent, nil},
		{"InvalidFee-Zero", args{0}, FeeZero, ErrUnsupportedFee},
		{"InvalidFee-Million", args{1000000}, FeeZero, ErrUnsupportedFee},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fee, err := GetFee(tt.args.percent)
			require.Equal(t, err, tt.wantErr)
			require.Equal(t, fee, tt.wantFee)
		})
	}
}
