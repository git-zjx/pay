package alipay

import (
	"pay/pkg/helper"
	"pay/pkg/exhttp"
	"pay/pkg/param"
	"strings"
)

func (client *Client) pos(payload param.Params) (param.Params, error){
	var (
		httpResp = param.Params{}
		resp     = param.Params{}
		url      = client.getUrl()
		sign     string
		err      error
	)
	bizContent := payload["biz_content"].(param.Params)
	bizContent["product_code"] = "FACE_TO_FACE_PAYMENT"
	bizContent["scene"] = "bar_code"
	payload["method"] = PayMethod
	payload["biz_content"] = helper.JsonMarshal(bizContent)
	if sign, err = client.generateSign(payload); err != nil {
		return nil, err
	}
	payload["sign"] = sign
	if httpResp, err = exhttp.Post(url, exhttp.TypeUrlencoded, strings.NewReader(payload.ToUrlValue()), nil); err != nil {
		return nil, err
	}
	resp, sign, err = client.getRespAndSignFromHttpResp(httpResp, PayMethod)
	if err != nil {
		return nil, err
	}
	if err = client.verifySign(resp, sign); err != nil {
		return nil, err
	}
	return resp, nil
}