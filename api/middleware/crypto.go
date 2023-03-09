package middleware

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
)

func EncryptoKey() (*rsa.PrivateKey, error) {
	rng := rand.Reader
	privateKey, err := rsa.GenerateKey(rng, 2048)
	if err != nil {
		return nil, errors.New("fail to genartekey")
	}
	fmt.Println(privateKey.PublicKey)
	return privateKey, nil
}
