package ethereum

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	chainmodels "github.com/ticken-ts/ticken-pubbc-connector/chain-models"
	"github.com/ticken-ts/ticken-pubbc-connector/ethereum/node"
	"github.com/ticken-ts/ticken-pubbc-connector/ethereum/scclient"
)

type Caller struct {
	submiter *scclient.Submiter
	querier  *scclient.Querier

	nc node.Connector
}

func NewCaller(nc *node.Connector, scAddr string, identity string) (*Caller, error) {
	if !nc.IsConnected() {
		return nil, fmt.Errorf("node connector is not connected")
	}

	submiter, err := scclient.NewSubmiter(nc, identity, scAddr, nil)
	if err != nil {
		return nil, err
	}

	querier, err := scclient.NewQuerier(nc, scAddr, nil)
	if err != nil {
		return nil, err
	}

	return &Caller{submiter: submiter, querier: querier}, nil
}

func (cc *Caller) MintTicket(buyerAddr string, section string) (string, error) {
	return cc.submiter.SubmitTx("safeMint", common.HexToAddress(buyerAddr), section)
}

func (cc *Caller) GetTickets(buyerAddr string, section string) ([]chainmodels.Ticket, error) {
	res, err := cc.querier.Query("getTicketsByOwner", common.HexToAddress(buyerAddr), section)
	if err != nil {
		return nil, err
	}

	tickets := *abi.ConvertType(res[0], new([]chainmodels.Ticket)).(*[]chainmodels.Ticket)

	return tickets, err
}
