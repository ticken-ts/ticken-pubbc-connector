package scclient

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ticken-ts/ticken-pubbc-connector/eth-connector/node"
)

type Querier struct {
	nc         *node.Connector
	scMetadata *bind.MetaData
}

func NewQuerier(nc *node.Connector, scMetadata *bind.MetaData) (*Querier, error) {
	return &Querier{nc: nc, scMetadata: scMetadata}, nil
}

func (querier *Querier) Query(scAddr string, method string, params ...interface{}) ([]interface{}, error) {
	sc, err := querier.nc.GetContract(scAddr, querier.scMetadata)
	if err != nil {
		return nil, err
	}
	var out []interface{}
	if err := sc.Call(nil, &out, method, params); err != nil {
		return nil, err
	}
	return out, nil
}
