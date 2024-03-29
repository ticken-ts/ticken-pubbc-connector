package eth_connector

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ticken-ts/ticken-pubbc-connector/eth-connector/node"
	"github.com/ticken-ts/ticken-pubbc-connector/eth-connector/scclient"
	"github.com/ticken-ts/ticken-pubbc-connector/eth-connector/wallet"
)

type Admin struct {
	wm      *wallet.Manager
	scAdmin *scclient.Admin
	nc      node.Connector

	scMetadata *bind.MetaData
}

func NewAdmin(nc *node.Connector, identity string) (*Admin, error) {
	if !nc.IsConnected() {
		return nil, fmt.Errorf("node connector is not connected")
	}

	scAdmin, err := scclient.NewAdmin(nc, identity)
	if err != nil {
		return nil, err
	}

	scMetadata, err := ReadMetadata()
	if err != nil {
		return nil, err
	}

	return &Admin{
		wm:         wallet.NewManager(),
		scAdmin:    scAdmin,
		scMetadata: scMetadata,
	}, nil
}

func (admin *Admin) DeployEventContract() (string, error) {
	scAddr, _, err := admin.scAdmin.Deploy(admin.scMetadata)
	if err != nil {
		return "", err
	}

	return scAddr.String(), nil
}

func (admin *Admin) CreateWallet() (string, string, string, error) {
	key, err := admin.wm.GenerateKey()
	if err != nil {
		return "", "", "", err
	}

	walletPrivKeyPem, walletPubKeyPem, err := admin.wm.PemEncodeKey(key)
	if err != nil {
		return "", "", "", err
	}

	walletAddress, err := admin.wm.GetAddressFromKey(key)
	if err != nil {
		return "", "", "", err
	}

	return walletPrivKeyPem, walletPubKeyPem, walletAddress, nil
}

func (admin *Admin) GetWalletForKey(walletPrivKeyPem string) (string, string, string, error) {
	key, err := admin.wm.PemDecodePrivateKey(walletPrivKeyPem)
	if err != nil {
		return "", "", "", err
	}

	walletPrivKeyPem, walletPubKeyPem, err := admin.wm.PemEncodeKey(key)
	if err != nil {
		return "", "", "", err
	}

	walletAddress, err := admin.wm.GetAddressFromKey(key)
	if err != nil {
		return "", "", "", err
	}

	return walletPrivKeyPem, walletPubKeyPem, walletAddress, nil
}
