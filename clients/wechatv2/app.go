package wechatv2

import (
	"pay/pkg/helper"
	"pay/pkg/param"
	"time"
)

func (client *Client) app(payload param.Params) (param.Params, error) {
	var (
		resp          = param.Params{}
		prePayResp = param.Params{}
		sign          string
		err           error
	)
	payload["trade_type"] = "APP"
	if prePayResp, err = client.prePay(payload); err != nil {
		return nil, err
	}
	if err = client.isSuccess(prePayResp); err != nil {
		return nil, err
	}
	resp["appid"] = client.config.AppId
	resp["partnerid"] = client.config.MchId
	resp["prepayid"] = prePayResp["prepay_id"]
	resp["package"] = "Sign=WXPay"
	resp["noncestr"] = helper.GenerateRandomString(32)
	resp["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	if sign, err = client.generateSign(resp); err != nil {
		return nil, err
	}
	resp["sign"] = sign
	return resp, nil
}
