package wechatv2

import "pay/pkg/param"

func (client *Client) scan(payload param.Params) (param.Params, error) {
	payload["trade_type"] = "NATIVE"
	return client.prePay(payload)
}
