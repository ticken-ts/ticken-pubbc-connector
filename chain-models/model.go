package chain_models

import (
	"math/big"
)

type TicketStatus string

const (
	// TicketStatusIssued represents the state of the
	// ticket right after it is "issued". Tickets in
	// this state can be scanned
	TicketStatusIssued TicketStatus = "issued"

	// TicketStatusScanned represents the state of the
	// ticket after it is "scanned". Note that this is
	// not  done in the same moment the scanning occurs.
	TicketStatusScanned TicketStatus = "scanned"

	// TicketStatusExpired represents the state of the
	// ticket after the event is finished and the ticket
	// never were scanned
	TicketStatusExpired TicketStatus = "expired"
)

type Ticket struct {
	TokenID *big.Int
	Section string
	Owner   string
	Status  TicketStatus
}
