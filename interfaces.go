package ticken_pubbc_connector

import chainmodels "github.com/ticken-ts/ticken-pubbc-connector/chain-models"

type NodeConnector interface {
	Connect() error
	IsConnected() bool
}

type Admin interface {
	DeployEventContract() (string, error)
	CreateWallet() (string, string, error)
}

type Caller interface {
	MintTicket(buyerAddr string, section string) (string, error)
	GetTickets(buyerAddr string, section string) ([]chainmodels.Ticket, error)
}
