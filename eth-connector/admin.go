package eth_connector

import (
	"fmt"
	"github.com/ticken-ts/ticken-pubbc-connector/eth-connector/node"
	"github.com/ticken-ts/ticken-pubbc-connector/eth-connector/scclient"
	"github.com/ticken-ts/ticken-pubbc-connector/eth-connector/wallet"
)

type Admin struct {
	wm      *wallet.Manager
	scAdmin *scclient.Admin
	nc      node.Connector
}

func NewAdmin(nc *node.Connector, identity string) (*Admin, error) {
	if !nc.IsConnected() {
		return nil, fmt.Errorf("node connector is not connected")
	}

	scAdmin, err := scclient.NewAdmin(nc, identity)
	if err != nil {
		return nil, err
	}

	return &Admin{wm: wallet.NewManager(), scAdmin: scAdmin}, nil
}

func (admin *Admin) DeployEventContract() (string, error) {
	scAddr, _, err := admin.scAdmin.Deploy(scEventMetadata)
	if err != nil {
		return "", err
	}
	return scAddr.String(), nil
}

func (admin *Admin) CreateWallet() (string, string, error) {
	walletPrivKey, err := admin.wm.GeneratePrivateKey()
	if err != nil {
		return "", "", err
	}

	walletAddress, err := admin.wm.GetAddressFromPrivateKey(walletPrivKey)
	if err != nil {
		return "", "", err
	}

	return walletPrivKey, walletAddress, nil
}
