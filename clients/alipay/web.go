package alipay

import (
	"pay/pkg/helper"
	"pay/pkg/param"
)

func (client *Client) web(payload param.Params) (param.Params, error) {
	var (
		resp    = param.Params{}
		url     = client.getUrl()
		sign    string
		err     error
	)
	bizContent := payload["biz_content"].(param.Params)
	bizContent["product_code"] = "FAST_INSTANT_TRADE_PAY"
	payload["method"] = PagePayMethod
	payload["biz_content"] = helper.JsonMarshal(bizContent)
	if sign, err = client.generateSign(payload); err != nil {
		return nil, err
	}
	payload["sign"] = sign
	resp["web_url"] = url + "?" + payload.ToUrlValue()
	return resp, nil
}
