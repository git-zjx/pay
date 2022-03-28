package sign

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	 "pay/pkg/exerror"
	"pay/pkg/helper"
	"pay/pkg/param"
	"strings"
)

const (
	MD5    = "MD5"
	SHA256 = "HMAC-SHA256"
)

func Generate(m interface{}, signType string, apiKey string) (string, error) {
	var (
		paramMap param.Params
		h        hash.Hash
		err      error
	)
	err = helper.JsonUnMarshal([]byte(helper.JsonMarshal(m)), &paramMap)
	if err != nil {
		return "", err
	}
	signData := helper.GenerateQueryStringExceptSign(paramMap)

	switch signType {
	case MD5:
		signData = signData + apiKey
		h = md5.New()
	case SHA256:
		h = hmac.New(sha256.New, []byte(apiKey))
	default:
		signData = signData + apiKey
		h = md5.New()
	}
	if _, err := h.Write([]byte(signData)); err != nil {
		return "", err
	}
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil))), nil
}

func Verify(m interface{}, signType string, apiKey string, retSign string) error {
	sign, err := Generate(m, signType, apiKey)
	if err != nil {
		return err
	}
	if sign != retSign {
		return exerror.SignNotMatchErr
	}
	return nil
}
