package alipay

import (
	"pay/pkg/helper"
	"pay/pkg/param"
)

func (client *Client) wap(payload param.Params) (param.Params, error){
	var (
		resp     = param.Params{}
		url      = client.getUrl()
		sign     string
		err      error
	)
	bizContent := payload["biz_content"].(param.Params)
	bizContent["product_code"] = "QUICK_WAP_PAY"
	payload["method"] = WapPayMethod
	payload["biz_content"] = helper.JsonMarshal(bizContent)
	if sign, err = client.generateSign(payload); err != nil {
		return nil, err
	}
	payload["sign"] = sign
	resp["wap_url"] = url + "?" + payload.ToUrlValue()
	return resp, nil
}