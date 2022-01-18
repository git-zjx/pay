package alipay

import (
	"pay/pkg/helper"
	"pay/pkg/http"
	"pay/pkg/param"
	"strings"
)

func (client *Client) refund(payload param.Params) (param.Params, error) {
	var (
		httpResp = param.Params{}
		resp     = param.Params{}
		url      = client.getUrl()
		sign     string
		err      error
	)
	payload["method"] = RefundMethod
	payload["biz_content"] = helper.JsonMarshal(payload["biz_content"])
	if sign, err = client.generateSign(payload); err != nil {
		return nil, err
	}
	payload["sign"] = sign
	if httpResp, err = http.Post(url, http.TypeUrlencoded, strings.NewReader(payload.ToUrlValue()), nil); err != nil {
		return nil, err
	}
	resp, sign, err = client.getRespAndSignFromHttpResp(httpResp, RefundMethod)
	if err != nil {
		return nil, err
	}
	if err = client.verifySign(resp, sign); err != nil {
		return nil, err
	}
	return resp, nil
}