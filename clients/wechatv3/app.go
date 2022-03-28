package wechatv3

import (
	"crypto/rsa"
	"fmt"
	netHttp "net/http"
	"pay/clients/wechatv3/pkg/http"
	 "pay/pkg/exerror"
	"pay/pkg/helper"
	"pay/pkg/param"
	"time"
)

func (client *Client) app(payload param.Params) (param.Params, error) {
	var (
		resp          = param.Params{}
		prePayResp    = param.Params{}
		privateKey    *rsa.PrivateKey
		prepayId      string
		authorization string
		ok            bool
		sign          string
		err           error
	)
	if privateKey, err = client.generatePrivateKey(); err != nil {
		return nil, err
	}
	if authorization, err = client.generateAuthorizationHeader(payload, netHttp.MethodPost, AppPayMethod, privateKey); err != nil {
		return nil, err
	}
	if prePayResp, err = http.Request(netHttp.MethodPost, client.getUrl(AppPayMethod), authorization, payload); err != nil {
		return nil, err
	}
	if err = client.isSuccess(prePayResp); err != nil {
		return nil, err
	}
	if _, ok = prePayResp["prepay_id"]; !ok {
		return nil, exerror.PrepayIdNotFoundErr
	}
	if prepayId, ok = prePayResp["prepay_id"].(string); !ok {
		return nil, exerror.PrepayIdFormatErr
	}
	resp["appid"] = client.config.AppId
	resp["partnerid"] = client.config.MchId
	resp["prepayid"] = prepayId
	resp["package"] = "Sign=WXPay"
	resp["noncestr"] = helper.GenerateRandomString(32)
	resp["timestamp"] = time.Now().Unix()
	if sign, err = client.generateSign(fmt.Sprintf("%s\n%d\n%s\n%s\n", resp["appid"], resp["timestamp"], resp["noncestr"], resp["prepayid"])); err != nil {
		return nil, err
	}
	resp["sign"] = sign
	return resp, nil
}
