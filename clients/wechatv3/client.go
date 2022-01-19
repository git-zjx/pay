package wechatv3

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"pay/clients/wechatv3/pkg/sign"
	"pay/pkg/constant"
	payErr "pay/pkg/error"
	"pay/pkg/helper"
	"pay/pkg/param"
	"time"
)

type Config struct {
	AppId     string
	MchId     string
	CertNo    string
	ApiKey    string
	CertPath  string
	KeyPath   string
	ReturnUrl string
	NotifyUrl string
	SignType  string
	Sandbox   bool
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

func (client *Client) getUrl(endPoint string) string {
	if client.config.Sandbox {
		return SandboxHost + endPoint
	}
	return Host + endPoint
}

func (client *Client) isSuccess(data param.Params) error {
	if _, ok := data["code"]; !ok {
		return nil
	}
	return errors.New(fmt.Sprintf("%s", data))
}

func (client *Client) generateSign(message string) (string, error) {
	privateKey, err := client.generatePrivateKey()
	if err != nil {
		return "", err
	}
	return sign.Generate(message, privateKey)
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
	message := fmt.Sprintf(SignatureMessageFormat, method, url, timestamp, nonce, helper.JsonMarshal(m))
	signatureResult, err := sign.Generate(message, privateKey)
	if err != nil {
		return "", err
	}
	authorization := fmt.Sprintf(AuthorizationFormat, AuthorizationType, client.config.MchId, nonce, timestamp, client.config.CertNo, signatureResult)
	return authorization, nil
}

func (client *Client) generatePrivateKey() (*rsa.PrivateKey, error) {
	privateKeyBytes, err := ioutil.ReadFile(client.config.KeyPath)
	if err != nil {
		return nil, fmt.Errorf("read private pem file err:%s", err.Error())
	}
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		return nil, fmt.Errorf("decode private key err")
	}
	if block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("the kind of PEM should be PRVATE KEY")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse private key err:%s", err.Error())
	}
	privateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("%s is not rsa private key", string(privateKeyBytes))
	}
	return privateKey, nil
}
