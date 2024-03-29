package eth_connector

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	chainmodels "github.com/ticken-ts/ticken-pubbc-connector/chain-models"
	"github.com/ticken-ts/ticken-pubbc-connector/eth-connector/node"
	"github.com/ticken-ts/ticken-pubbc-connector/eth-connector/scclient"
	"math/big"
)

type Caller struct {
	submiter *scclient.Submiter
	querier  *scclient.Querier

	nc node.Connector
}

type EthTicket struct {
	Owner   common.Address
	Status  uint8
	Section string
	TokenID *big.Int
}

func NewCaller(nc *node.Connector, identity string) (*Caller, error) {
	if !nc.IsConnected() {
		return nil, fmt.Errorf("node connector is not connected")
	}

	scMetadata, err := ReadMetadata()
	if err != nil {
		return nil, err
	}

	submiter, err := scclient.NewSubmiter(nc, identity, scMetadata)
	if err != nil {
		return nil, err
	}

	querier, err := scclient.NewQuerier(nc, scMetadata)
	if err != nil {
		return nil, err
	}

	return &Caller{submiter: submiter, querier: querier}, nil
}

func (cc *Caller) MintTicket(scAddr string, buyerAddr string, section string, tokenID *big.Int) (string, error) {
	return cc.submiter.SubmitTx(scAddr, "mintTicket", common.HexToAddress(buyerAddr), tokenID, section)
}

func (cc *Caller) GetTickets(scAddr string, userAddr string) ([]*chainmodels.Ticket, error) {
	res, err := cc.querier.Query(scAddr, "getTicketsOwnedBy", common.HexToAddress(userAddr))
	if err != nil {
		return nil, err
	}

	ethTickets := *abi.ConvertType(res[0], new([]EthTicket)).(*[]EthTicket)

	var tickets = make([]*chainmodels.Ticket, 0)
	for _, ethTicket := range ethTickets {
		tickets = append(tickets, ethTicketToTicket(&ethTicket))
	}

	return tickets, err
}

func (cc *Caller) GetTicket(scAddr string, tokenID *big.Int) (*chainmodels.Ticket, error) {
	res, err := cc.querier.Query(scAddr, "tickets", tokenID)
	if err != nil {
		return nil, err
	}

	ethTicket := EthTicket{
		Owner:   res[0].(common.Address),
		Status:  res[1].(uint8),
		Section: res[2].(string),
		TokenID: res[3].(*big.Int),
	}

	return ethTicketToTicket(&ethTicket), err
}

func (cc *Caller) TransferTicket(scAddr string, tokenID *big.Int, fromAddr string, toAddr string) (string, error) {
	return cc.submiter.SubmitTx(
		scAddr,
		"transferTicket",
		common.HexToAddress(fromAddr),
		common.HexToAddress(toAddr),
		tokenID,
	)
}

func (cc *Caller) RaiseAnchors(scAddr string) (string, error) {
	return cc.submiter.SubmitTx(scAddr, "raiseAnchors")
}

func ethTicketToTicket(ethTicket *EthTicket) *chainmodels.Ticket {
	var status chainmodels.TicketStatus

	switch ethTicket.Status {
	case 0:
		status = chainmodels.TicketStatusIssued
		break
	case 1:
		status = chainmodels.TicketStatusScanned
		break
	case 2:
		status = chainmodels.TicketStatusExpired
		break
	}

	return &chainmodels.Ticket{
		Owner:   ethTicket.Owner.Hex(),
		TokenID: ethTicket.TokenID,
		Section: ethTicket.Section,
		Status:  status,
	}
}
