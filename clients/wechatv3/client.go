package wechatv3

import (
	"pay/pkg/constant"
	payErr "pay/pkg/error"
	"pay/pkg/helper"
	"pay/pkg/param"
)

type Config struct {
	AppId      string
	MchId      string
	ApiKey     string
	CertPath   string
	KeyPath    string
	Pkcs12Path string
	ReturnUrl  string
	NotifyUrl  string
	SignType   string
	Sandbox    bool
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