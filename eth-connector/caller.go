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

func NewCaller(nc *node.Connector, identity string) (*Caller, error) {
	if !nc.IsConnected() {
		return nil, fmt.Errorf("node connector is not connected")
	}

	submiter, err := scclient.NewSubmiter(nc, identity, scEventMetadata)
	if err != nil {
		return nil, err
	}

	querier, err := scclient.NewQuerier(nc, scEventMetadata)
	if err != nil {
		return nil, err
	}

	return &Caller{submiter: submiter, querier: querier}, nil
}

func (cc *Caller) MintTicket(scAddr string, buyerAddr string, section string, tokenID *big.Int) (string, error) {
	return cc.submiter.SubmitTx(scAddr, "safeMint", common.HexToAddress(buyerAddr), section, tokenID)
}

func (cc *Caller) GetTickets(scAddr string, userAddr string) ([]chainmodels.Ticket, error) {
	res, err := cc.querier.Query(scAddr, "getTicketsByOwner", common.HexToAddress(userAddr))
	if err != nil {
		return nil, err
	}

	tickets := *abi.ConvertType(res[0], new([]chainmodels.Ticket)).(*[]chainmodels.Ticket)

	return tickets, err
}

func (cc *Caller) GetTicket(scAddr string, tokenID *big.Int) (*chainmodels.Ticket, error) {
	res, err := cc.querier.Query(scAddr, "getTicket", tokenID)
	if err != nil {
		return nil, err
	}

	ticket := abi.ConvertType(res[0], new(chainmodels.Ticket)).(chainmodels.Ticket)

	return &ticket, err
}

func (cc *Caller) TransferTicket(scAddr string, tokenID *big.Int, fromAddr string, toAddr string) (string, error) {
	return cc.submiter.SubmitTx(
		scAddr,
		"safeTransferFrom",
		common.HexToAddress(fromAddr),
		common.HexToAddress(toAddr),
		tokenID,
	)
}
