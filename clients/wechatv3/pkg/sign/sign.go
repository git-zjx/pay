package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
)

func Generate(message string, privateKey *rsa.PrivateKey) (string, error) {
	h := crypto.Hash.New(crypto.SHA256)
	_, err := h.Write([]byte(message))
	if err != nil {
		return "", err
	}
	hashed := h.Sum(nil)
	signatureByte, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signatureByte), nil
}

func Verify() {

}
