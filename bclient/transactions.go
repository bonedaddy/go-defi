package bclient

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
)

// EncodeTx is used to marshal a transaction and hex encode it
// suitable for encoding transctions and storing them into a database
func (bc *BClient) EncodeTx(hash string) (string, error) {
	tx, _, err := bc.bc.TransactionByHash(bc.ctx, common.HexToHash(hash))
	if err != nil {
		return "", err
	}
	txBinary, err := tx.MarshalBinary()
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(txBinary), nil
}
