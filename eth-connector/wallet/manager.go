package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/pem"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

type Manager struct {
}

func NewManager() *Manager {
	return new(Manager)
}

// GenerateKey Generate a key pair of public and private
// key, using the ECDSA algorithm ver an P256 elliptic curve
func (m *Manager) GenerateKey() (*ecdsa.PrivateKey, error) {
	ecdsaKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	return ecdsaKey, nil
}

func (m *Manager) GetAddressFromKey(ecdsaKey *ecdsa.PrivateKey) (string, error) {
	return crypto.PubkeyToAddress(ecdsaKey.PublicKey).String(), nil
}

func (m *Manager) PemEncodeKey(privKey *ecdsa.PrivateKey) (string, string, error) {
	x509Encoded := crypto.FromECDSA(privKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub := crypto.FromECDSAPub(&privKey.PublicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncoded), string(pemEncodedPub), nil
}

func (m *Manager) PemDecodePublicKey(publicKeyString string) (*ecdsa.PublicKey, error) {
	pemPublicKey, _ := pem.Decode([]byte(publicKeyString))
	pubKey, err := crypto.UnmarshalPubkey(pemPublicKey.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode pem public key: %s", err.Error())
	}
	return pubKey, nil
}

func (m *Manager) PemDecodePrivateKey(privateKeyString string) (*ecdsa.PrivateKey, error) {
	pemPrivKey, _ := pem.Decode([]byte(privateKeyString))
	privKey, err := crypto.ToECDSA(pemPrivKey.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode pem private key: %s", err.Error())
	}
	return privKey, nil
}
