package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
)

const (
	privateKeyLen = 64
	publicKeyLen  = 32
	seedLen       = 32
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

func (pk *PrivateKey) Bytes() []byte {
	return pk.key
}

func (pk *PrivateKey) Sign(data []byte) ([]byte, error) {
	return ed25519.Sign(pk.key, data), nil
}

func (pk *PrivateKey) Public() *PublicKey {
	return &PublicKey{pk.key.Public().(ed25519.PublicKey)}
}

func GenerateKey() (*PrivateKey, error) {
	seed := make([]byte, seedLen)
	_, err := rand.Read(seed)
	if err != nil {
		panic(err)
	}
	private := ed25519.NewKeyFromSeed(seed)
	return &PrivateKey{private}, nil
}

type PublicKey struct {
	key ed25519.PublicKey
}

func (pk *PublicKey) Bytes() []byte {
	return pk.key
}
