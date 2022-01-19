package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"pay/pkg/helper"
)

const SignatureMessageFormat = "%s\n%s\n%d\n%s\n%s\n"

func Generate(m interface{}, method, url string, timestamp int64, nonce string, privateKey *rsa.PrivateKey) (string, error) {
	signBody := helper.JsonMarshal(m)
	message := fmt.Sprintf(SignatureMessageFormat, method, url, timestamp, nonce, signBody)
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
