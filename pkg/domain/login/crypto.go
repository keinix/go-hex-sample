package login

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/argon2"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var p = &params{
	memory:      64 * 1024,
	iterations:  3,
	parallelism: 2,
	saltLength:  16,
	keyLength:   512,
}

func newPasswordHash(password string) (string, error) {
	salt, err := generateSalt(p.saltLength)
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)
	encodedHash := encodeHash(hash)
	return encodedHash, nil
}

func generateSalt(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("error generating salt: %w", err)
	}
	return b, nil
}

// convert the hash to a format that has the necessary information
// for the hash to be replicated given the same password input
func encodeHash(hash []byte) string {

}

func checkPlainTextMatchesHash(plainText string, hash string) (bool, error) {

}

func newSessionToken() string {

}
