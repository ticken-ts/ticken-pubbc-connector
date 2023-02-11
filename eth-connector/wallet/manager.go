package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
)

type Manager struct {
}

func NewManager() *Manager {
	return &Manager{}
}

// GeneratePrivateKey Generate private key
func (m *Manager) GeneratePrivateKey() (string, error) {
	pk, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(crypto.FromECDSA(pk)), nil
}

// GetAddressFromPrivateKey Get address from private key
func (m *Manager) GetAddressFromPrivateKey(privKey string) (string, error) {
	ecdsaKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return "", err
	}

	return crypto.PubkeyToAddress(ecdsaKey.PublicKey).String(), nil
}
