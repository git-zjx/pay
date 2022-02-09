package wechatv3

import (
	"crypto/rsa"
	netHttp "net/http"
	"pay/clients/wechatv3/pkg/http"
	payErr "pay/pkg/error"
	"pay/pkg/param"
)

func (client *Client) scan(payload param.Params) (param.Params, error) {
	var (
		resp          = param.Params{}
		prePayResp    = param.Params{}
		privateKey    *rsa.PrivateKey
		codeUrl       string
		authorization string
		ok            bool
		err           error
	)
	if privateKey, err = client.generatePrivateKey(); err != nil {
		return nil, err
	}
	if authorization, err = client.generateAuthorizationHeader(payload, netHttp.MethodPost, ScanPayMethod, privateKey); err != nil {
		return nil, err
	}
	if prePayResp, err = http.Request(netHttp.MethodPost, client.getUrl(ScanPayMethod), authorization, payload); err != nil {
		return nil, err
	}
	if err = client.isSuccess(prePayResp); err != nil {
		return nil, err
	}
	if _, ok = prePayResp["code_url"]; !ok {
		return nil, payErr.PrepayIdNotFoundErr
	}
	if codeUrl, ok = prePayResp["code_url"].(string); !ok {
		return nil, payErr.PrepayIdFormatErr
	}
	resp["code_url"] = codeUrl
	return resp, nil
}
