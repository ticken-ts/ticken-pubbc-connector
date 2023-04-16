package ticken_pubbc_connector

import (
	chainmodels "github.com/ticken-ts/ticken-pubbc-connector/chain-models"
	"math/big"
)

type NodeConnector interface {
	Connect() error
	IsConnected() bool
}

type Admin interface {
	// DeployEventContract is a blocking function that deploys
	// the event smart contract and returns the contract address
	DeployEventContract() (string, error)

	// CreateWallet creates a wallet for a new user in the
	// public blockchain. It returns the private key and the
	// wallet address derived
	CreateWallet() (string, string, error)

	// GetWalletForKey generates the wallet address associated to
	// the private key passed by parameter. This method is idempotent,
	// so, no matter how many times we pass the private key, it will
	// always return the same wallet address
	GetWalletForKey(walletPrivKey string) (string, error)
}

type Caller interface {
	// MintTicket is a blocking function that generates a ticket
	// in the public blockchain and assign the buyer as owner
	// It returns the transaction ID that generated the ticket
	MintTicket(scAddr string, buyerAddr string, section string, tokenID *big.Int) (string, error)

	// GetTickets returns all tickets owned by the given user
	GetTickets(scAddr string, userAddr string) ([]*chainmodels.Ticket, error)

	// GetTicket return the ticket that match the ticket ID that was minted
	// in the contract with address scAddr
	GetTicket(scAddr string, tokenID *big.Int) (*chainmodels.Ticket, error)

	// TransferTicket transfer a ticket with token id "tokenID" from the owner
	// "fromAddr" to the new owner "toAddr". This method can be only invoked by
	// the contract address
	TransferTicket(scAddr string, tokenID *big.Int, fromAddr string, toAddr string) (string, error)
}
