package alipay

import (
	"pay/pkg/helper"
	"pay/pkg/param"
)

func (client *Client) app(payload param.Params) (param.Params, error) {
	var (
		resp    = param.Params{}
		sign    string
		err     error
	)
	payload["method"] = AppPayMethod
	payload["biz_content"] = helper.JsonMarshal(payload["biz_content"])
	if sign, err = client.generateSign(payload); err != nil {
		return nil, err
	}
	payload["sign"] = sign
	resp["app_param"] = payload.ToUrlValue()
	return resp, nil
}
