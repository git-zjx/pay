package wechatv3

import (
	"crypto/rsa"
	"fmt"
	"pay/clients/wechatv3/pkg/sign"
	"pay/pkg/constant"
	payErr "pay/pkg/error"
	"pay/pkg/helper"
	"pay/pkg/param"
	"time"
)

type Config struct {
	AppId               string
	MchId               string
	CertificateSerialNo string
	ApiKey              string
	CertPath            string
	KeyPath             string
	Pkcs12Path          string
	ReturnUrl           string
	NotifyUrl           string
	SignType            string
	Sandbox             bool
}

type Client struct {
	config *Config
}

func NewClient(config *Config) *Client {
	client := new(Client)
	client.config = config
	return client
}

func (client *Client) Pay(method string, request param.Params) (param.Params, error) {
	payload := client.generatePayload(request)
	switch method {
	case constant.MP:
		return client.mp(payload)
	case constant.WAP:
		return client.wap(payload)
	case constant.APP:
		return client.app(payload)
	case constant.MINI:
		return client.mini(payload)
	case constant.POS:
		return client.pos(payload)
	case constant.SCAN:
		return client.scan(payload)
	default:
		return nil, payErr.PayMethodErr
	}
}

func (client *Client) generatePayload(request param.Params) param.Params {
	request["appid"] = client.config.AppId
	request["mchid"] = client.config.MchId
	request["notify_url"] = client.config.NotifyUrl
	request["nonce_str"] = helper.GenerateRandomString(32)
	return request
}

func (client *Client) generateAuthorizationHeader(m interface{}, method, url string, privateKey *rsa.PrivateKey) (string, error) {
	nonce := helper.GenerateRandomString(32)
	timestamp := time.Now().Unix()
	signatureResult, err := sign.Generate(m, method, url, timestamp, nonce, privateKey)
	if err != nil {
		return "", err
	}
	authorization := fmt.Sprintf(AuthorizationFormat, "WECHATPAY2-SHA256-RSA2048", client.config.MchId, nonce, timestamp, client.config.CertificateSerialNo, signatureResult)
	return authorization, nil
}
