package chain_models

import (
	"math/big"
)

type Ticket struct {
	TokenID *big.Int
	Section string
	Status  string
	Owner   string
}
