package scclient

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ticken-ts/ticken-pubbc-connector/ethereum/node"
)

type Admin struct {
	nc       *node.Connector
	identity string
}

func NewAdmin(nc *node.Connector, identity string) (*Admin, error) {
	return &Admin{nc: nc, identity: identity}, nil
}

func (a *Admin) Deploy(sc *bind.MetaData) (common.Address, *types.Transaction, error) {
	parsedABI, err := sc.GetAbi()
	if err != nil {
		return common.Address{}, nil, err
	}
	if parsedABI == nil {
		return common.Address{}, nil, errors.New("GetABI returned nil")
	}

	transactor, err := a.nc.GetTransactor(a.identity)
	if err != nil {
		return common.Address{}, nil, err
	}

	address, tx, _, err := bind.DeployContract(transactor, *parsedABI, common.FromHex(sc.Bin), a.nc.EthCLI)
	if err != nil {
		return common.Address{}, nil, err
	}

	return address, tx, nil
}
