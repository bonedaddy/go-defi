package bclient

import "errors"

var (
	// ErrNotEthClient is an error return when attempting to do type conversions against the
	// blockchain interface to an ethclient, when the underlying blockchain interface is not
	// an ethclient type. Usually will be encountered when using a simulated backend
	ErrNotEthClient = errors.New("underlying blockchain provider is not of type ethclient.Client")
)
