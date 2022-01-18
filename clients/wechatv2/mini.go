package wechatv2

import "pay/pkg/param"

func (client *Client) mini(payload param.Params) (param.Params, error) {
	return client.mp(payload)
}
