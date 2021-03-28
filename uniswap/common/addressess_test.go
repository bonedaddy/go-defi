package common

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

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
			got, got1 := SortAddressess(tt.args.tkn0, tt.args.tkn1)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortAddressess() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("sortAddressess() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
