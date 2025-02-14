package crypto

import (
	"crypto/ed25519"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrivateKeyLen(t *testing.T) {
	privateKey, err := GenerateKey()
	assert.NoError(t, err, "Error generating key")
	assert.Len(t, privateKey.key, privateKeyLen, "Wrong private key length")
}

func TestPublicKeyLen(t *testing.T) {
	privateKey, err := GenerateKey()
	assert.NoError(t, err, "Error generating key")

	publicKey := privateKey.Public()
	assert.Len(t, publicKey.key, publicKeyLen, "Wrong public key length")
}

func TestPrivateKeyBytes(t *testing.T) {
	privateKey, _ := GenerateKey()
	bytes := privateKey.Bytes()

	assert.EqualValues(t, bytes, []byte(privateKey.key), "Private key bytes must be equal")
}

func TestPublicKeyBytes(t *testing.T) {
	privateKey, _ := GenerateKey()
	publicKey := privateKey.Public()
	bytes := publicKey.Bytes()

	assert.EqualValues(t, bytes, []byte(publicKey.key), "Public key bytes must be equal")
}

func TestSignature(t *testing.T) {
	privateKey, _ := GenerateKey()
	publicKey := privateKey.Public()

	message := []byte("Hello, World!")
	signature, err := privateKey.Sign(message)
	assert.NoError(t, err, "Error signing message")
	assert.NotEmpty(t, signature, "Empty signature")

	isValid := ed25519.Verify(publicKey.key, message, signature)
	assert.True(t, isValid, "Signature must be valid")

	wrongMessage := []byte("Hello, Go!")
	isValid = ed25519.Verify(publicKey.key, wrongMessage, signature)
	assert.False(t, isValid, "Signature must be invalid")
}
