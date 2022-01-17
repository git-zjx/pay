package alipay

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"pay/clients/alipay/pkg/key"
	"pay/clients/alipay/pkg/sign"
	"pay/pkg/constant"
	payErr "pay/pkg/error"
	"pay/pkg/helper"
	"pay/pkg/param"
	"time"
)

type Config struct {
	AppId           string
	AppPrivateKey   string
	AlipayPublicKey string
	ReturnUrl       string
	NotifyUrl       string
	SignType        string
	Sandbox         bool
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
	case constant.WEB:
		return client.web(payload)
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

func (client *Client) Find(request param.Params, isRefund bool) (param.Params, error) {
	payload := client.generatePayload(request)
	if isRefund {
		return client.refundQuery(payload)
	}
	return client.query(payload)
}

func (client *Client) Refund(request param.Params) (param.Params, error) {
	payload := client.generatePayload(request)
	return client.refund(payload)
}

func (client *Client) Close(request param.Params) (param.Params, error) {
	payload := client.generatePayload(request)
	return client.close(payload)
}

func (client *Client) Cancel(request param.Params) (param.Params, error) {
	payload := client.generatePayload(request)
	return client.cancel(payload)
}

func (client *Client) Verify(request param.Params) (param.Params, error) {
	var err error
	if err = sign.Verify(request, client.config.SignType, []byte(client.config.AlipayPublicKey), request["sign"].(string)); err != nil {
		return nil, err
	}
	return request, nil
}

func (client *Client) Success() {
	fmt.Println("success")
}

func (client *Client) getUrl() string {
	if client.config.Sandbox {
		return SandboxHost
	}
	return Host
}

func (client *Client) generateSign(params param.Params) (string, error) {
	var (
		priKey  *rsa.PrivateKey
		signRes string
		err     error
	)

	if priKey, err = key.DecodePrivateKey([]byte(key.FormatPrivateKey(client.config.AppPrivateKey))); err != nil {
		return signRes, err
	}
	if signRes, err = sign.Generate(params, client.config.SignType, priKey); err != nil {
		return signRes, err
	}
	return signRes, nil
}

func (client *Client) verifySign(params param.Params, retSign string) error {
	var (
		err error
	)
	if err = sign.Verify(params, client.config.SignType, []byte(client.config.AlipayPublicKey), retSign); err != nil {
		return err
	}
	return nil
}

func (client *Client) isSuccess(data param.Params) error {
	if data["code"].(string) == "10000" {
		return nil
	}
	return errors.New(fmt.Sprintf("%s", data))
}

func (client *Client) getRespAndSignFromHttpResp(httpResp param.Params, method string) (param.Params, string, error) {
	var (
		resp param.Params
		err  error
	)
	data, ok := httpResp[helper.BuildResponseKeyFromMethod(method)]
	if !ok {
		return nil, "", payErr.PayReturnParamFormatErr
	}
	resp = data.(map[string]interface{})
	if err = client.isSuccess(resp); err != nil {
		return nil, "", err
	}
	retSign, ok := httpResp["sign"]
	if !ok {
		return nil, "", payErr.PayReturnParamNotHaveSignErr
	}
	return resp, retSign.(string), nil
}

func (client *Client) generatePayload(request param.Params) param.Params {
	var payload = param.Params{}
	payload["app_id"] = client.config.AppId
	payload["format"] = "JSON"
	payload["charset"] = "utf-8"
	payload["sign_type"] = client.config.SignType
	payload["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	payload["version"] = "1.0"
	payload["return_url"] = client.config.ReturnUrl
	payload["notify_url"] = client.config.NotifyUrl
	payload["biz_content"] = request
	return payload
}
