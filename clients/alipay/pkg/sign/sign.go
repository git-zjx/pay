package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"hash"
	"pay/clients/alipay/pkg/key"
	"pay/pkg/helper"
	"pay/pkg/param"
)

const (
	RSA  = "RSA"
	RSA2 = "RSA2"
)

func Generate(m interface{}, signType string, privateKey *rsa.PrivateKey) (string, error) {
	var (
		encryptedBytes []byte
		h              hash.Hash
		hasher         crypto.Hash
		err            error
		paramMap       param.Params
	)
	err = helper.JsonUnMarshal([]byte(helper.JsonMarshal(m)), &paramMap)
	if err != nil {
		return "", err
	}
	signData := helper.GenerateQueryStringExceptSign(paramMap)
	switch signType {
	case RSA:
		h = sha1.New()
		hasher = crypto.SHA1
	case RSA2:
		h = sha256.New()
		hasher = crypto.SHA256
	default:
		h = sha256.New()
		hasher = crypto.SHA256
	}

	if _, err := h.Write([]byte(signData)); err != nil {
		return "", err
	}

	encryptedBytes, err = rsa.SignPKCS1v15(rand.Reader, privateKey, hasher, h.Sum(nil))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}

func Verify(m interface{}, signType string, publicKey []byte, sign string) error {
	var (
		h        hash.Hash
		hasher   crypto.Hash
		paramMap param.Params
		err      error
	)
	pKey, err := key.DecodePublicKey(publicKey)
	if err != nil {
		return err
	}
	signBytes, _ := base64.StdEncoding.DecodeString(sign)
	err = helper.JsonUnMarshal([]byte(helper.JsonMarshal(m)), &paramMap)
	if err != nil {
		return err
	}
	signData := helper.GenerateQueryStringExceptSign(paramMap)
	switch signType {
	case RSA:
		hasher = crypto.SHA1
	case RSA2:
		hasher = crypto.SHA256
	default:
		hasher = crypto.SHA256
	}
	h = hasher.New()
	h.Write([]byte(signData))
	return rsa.VerifyPKCS1v15(pKey, hasher, h.Sum(nil), signBytes)
}

