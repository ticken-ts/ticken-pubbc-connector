package ticken_pubbc_connector

import chainmodels "github.com/ticken-ts/ticken-pubbc-connector/chain-models"

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
}

type Caller interface {
	// MintTicket is a blocking function that generates a ticket
	// in the public blockchain and assign the buyer as owner
	// It returns the transaction ID that generated the ticket
	MintTicket(buyerAddr string, section string) (string, error)

	// GetTickets returns all tickets owned by the given user
	GetTickets(buyerAddr string, section string) ([]chainmodels.Ticket, error)
}
