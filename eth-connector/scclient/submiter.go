package scclient

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ticken-ts/ticken-pubbc-connector/eth-connector/node"
)

type Submiter struct {
	nc         *node.Connector
	transactor *bind.TransactOpts
	scMetadata *bind.MetaData
}

func NewSubmiter(nc *node.Connector, identity string, scMetadata *bind.MetaData) (*Submiter, error) {
	transactor, err := nc.GetTransactor(identity)
	if err != nil {
		return nil, err
	}

	return &Submiter{nc: nc, scMetadata: scMetadata, transactor: transactor}, nil
}

func (submiter *Submiter) SubmitTx(scAddr string, method string, params ...interface{}) (string, error) {
	sc, err := submiter.nc.GetContract(scAddr, submiter.scMetadata)
	if err != nil {
		return "", err
	}

	tx, err := sc.Transact(submiter.transactor, method, params)
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
