package scclient

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ticken-ts/ticken-pubbc-connector/ethereum/node"
)

type Submiter struct {
	nc         *node.Connector
	sc         *bind.BoundContract
	transactor *bind.TransactOpts
}

func NewSubmiter(nc *node.Connector, identity string, scAddr string, scMetadata *bind.MetaData) (*Submiter, error) {
	sc, err := nc.GetContract(scAddr, scMetadata)
	if err != nil {
		return nil, err
	}

	transactor, err := nc.GetTransactor(identity)
	if err != nil {
		return nil, err
	}

	return &Submiter{nc: nc, sc: sc, transactor: transactor}, nil
}

func (submiter *Submiter) SubmitTx(method string, params ...interface{}) (string, error) {
	tx, err := submiter.sc.Transact(submiter.transactor, method, params)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	// Wait for transaction to be mined
	_, err = bind.WaitMined(submiter.transactor.Context, submiter.nc.EthCLI, tx)
	if err != nil {
		return "", err
	}

	return tx.Hash().String(), nil
}
