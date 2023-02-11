package scclient

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ticken-ts/ticken-pubbc-connector/ethereum/node"
)

type Querier struct {
	nc *node.Connector
	sc *bind.BoundContract
}

func NewQuerier(nc *node.Connector, scAddr string, scMetadata *bind.MetaData) (*Querier, error) {
	sc, err := nc.GetContract(scAddr, scMetadata)
	if err != nil {
		return nil, err
	}
	return &Querier{nc: nc, sc: sc}, nil
}

func (querier *Querier) Query(method string, params ...interface{}) ([]interface{}, error) {
	var out []interface{}
	if err := querier.sc.Call(nil, &out, method, params); err != nil {
		return nil, err
	}
	return out, nil
}
