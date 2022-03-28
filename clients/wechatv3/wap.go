package wechatv3

import (
	"crypto/rsa"
	netHttp "net/http"
	"pay/clients/wechatv3/pkg/http"
	 "pay/pkg/exerror"
	"pay/pkg/param"
)

func (client *Client) wap(payload param.Params) (param.Params, error) {
	var (
		resp          = param.Params{}
		prePayResp    = param.Params{}
		privateKey    *rsa.PrivateKey
		h5Url         string
		authorization string
		ok            bool
		err           error
	)
	if privateKey, err = client.generatePrivateKey(); err != nil {
		return nil, err
	}
	if authorization, err = client.generateAuthorizationHeader(payload, netHttp.MethodPost, WapPayMethod, privateKey); err != nil {
		return nil, err
	}
	if prePayResp, err = http.Request(netHttp.MethodPost, client.getUrl(WapPayMethod), authorization, payload); err != nil {
		return nil, err
	}
	if err = client.isSuccess(prePayResp); err != nil {
		return nil, err
	}
	if _, ok = prePayResp["h5_url"]; !ok {
		return nil, exerror.PrepayIdNotFoundErr
	}
	if h5Url, ok = prePayResp["h5_url"].(string); !ok {
		return nil, exerror.PrepayIdFormatErr
	}
	resp["h5_url"] = h5Url
	return resp, nil
}
