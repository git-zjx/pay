package wechatv3

import (
	"crypto/rsa"
	"fmt"
	netHttp "net/http"
	"pay/clients/wechatv3/pkg/http"
	"pay/pkg/param"
)

func (client *Client) close(payload param.Params) (param.Params, error) {
	var (
		resp          = param.Params{}
		privateKey    *rsa.PrivateKey
		authorization string
		err           error
	)
	if privateKey, err = client.generatePrivateKey(); err != nil {
		return nil, err
	}
	if authorization, err = client.generateAuthorizationHeader(payload, netHttp.MethodPost, AppPayMethod, privateKey); err != nil {
		return nil, err
	}
	if resp, err = http.Request(netHttp.MethodPost, client.getUrl(fmt.Sprintf(CloseMethod, payload["out_trade_no"])), authorization, payload); err != nil {
		return nil, err
	}
	return resp, nil
}
