package node

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Connector struct {
	chainURL string
	EthCLI   *ethclient.Client
	chainID  *big.Int
}

func New(chainUrl string) *Connector {
	return &Connector{chainURL: chainUrl}
}

func (nc *Connector) Connect() error {
	cli, err := ethclient.Dial(nc.chainURL)
	if err != nil {
		return err
	}

	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		return err
	}

	nc.EthCLI = cli
	nc.chainID = chainID
	return nil
}

func (nc *Connector) IsConnected() bool {
	return nc.EthCLI != nil
}

func (nc *Connector) GetContract(scAddr string, scMetadata *bind.MetaData) (*bind.BoundContract, error) {
	parsed, err := scMetadata.GetAbi()
	if err != nil {
		return nil, err
	}

	return bind.NewBoundContract(common.HexToAddress(scAddr), *parsed, nc.EthCLI, nc.EthCLI, nc.EthCLI), nil
}

// GetTransactor convert pk as hex string to a transactor object for contract calls
func (nc *Connector) GetTransactor(privKey string) (*bind.TransactOpts, error) {
	ecdsaKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, err
	}

	return bind.NewKeyedTransactorWithChainID(ecdsaKey, big.NewInt(nc.chainID.Int64()))
}
