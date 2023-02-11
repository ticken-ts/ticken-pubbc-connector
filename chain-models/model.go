package chain_models

import "math/big"

type Ticket struct {
	TokenID *big.Int
	Section string
	Scanned bool
}
