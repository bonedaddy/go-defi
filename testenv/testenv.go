package testenv

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/bonedaddy/go-defi/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type Testenv struct {
	ctx    context.Context
	cancel context.CancelFunc
	pk     *ecdsa.PrivateKey
	Auth   *bind.TransactOpts
	*backends.SimulatedBackend
}

// NewBlockchain is used to generate a simulated blockchain
func NewBlockchain(ctx context.Context) (*Testenv, error) {
	ctx, cancel := context.WithCancel(ctx)
	auth, pk, err := utils.NewAccount()
	if err != nil {
		cancel()
		return nil, err
	}
	// https://medium.com/coinmonks/unit-testing-solidity-contracts-on-ethereum-with-go-3cc924091281
	balance := new(big.Int).Mul(big.NewInt(999999999999999), big.NewInt(999999999999999))
	gAlloc := map[common.Address]core.GenesisAccount{
		auth.From: {Balance: balance},
	}
	sim := backends.NewSimulatedBackend(gAlloc, 8000000)
	return &Testenv{ctx: ctx, cancel: cancel, pk: pk, Auth: auth, SimulatedBackend: sim}, nil
}

func (t *Testenv) DoWaitMined(tx *types.Transaction, printArgs ...string) error {
	t.Commit()
	rcpt, err := bind.WaitMined(t.ctx, t, tx)
	if len(printArgs) > 0 {
		log.Println("gas used by transaction: ", rcpt.CumulativeGasUsed, printArgs)
	}
	return errors.Wrap(err, "wait mined")
}

func (t *Testenv) DoWaitDeployed(tx *types.Transaction, printArgs ...string) (common.Address, error) {
	t.Commit()
	addr, err := bind.WaitDeployed(t.ctx, t, tx)
	if err != nil {
		return common.Address{}, err
	}
	rcpt, err := t.TransactionReceipt(t.ctx, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}
	if len(printArgs) > 0 {
		log.Println("gas used by deployment transaction: ", rcpt.CumulativeGasUsed, printArgs)
	}
	return addr, nil
}

func (t *Testenv) SendETH(recipient common.Address, value *big.Int) error {
	fromAddress := t.Auth.From
	nonce, err := t.PendingNonceAt(t.ctx, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasLimit := uint64(21000) // in units
	gasPrice, err := t.SuggestGasPrice(t.ctx)
	if err != nil {
		log.Fatal(err)
	}

	var data []byte
	tx := types.NewTransaction(nonce, recipient, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, t.pk)
	if err != nil {
		log.Fatal(err)
	}

	if err := t.SendTransaction(t.ctx, signedTx); err != nil {
		return errors.Wrap(err, "send transaction")
	}
	t.DoWaitMined(signedTx, "sendeth", signedTx.Hash().Hex())
	return nil
}

// BlockNumber returns the most recent block number
func (t *Testenv) BlockNumber(ctx context.Context) (uint64, error) {
	return t.Blockchain().CurrentBlock().NumberU64(), nil
}

func (t *Testenv) Context() context.Context {
	return t.ctx
}

func (t *Testenv) Cancel() {
	t.cancel()
}

func (t *Testenv) PK() *ecdsa.PrivateKey {
	return t.pk
}
