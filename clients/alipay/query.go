package alipay

import (
	"pay/pkg/helper"
	"pay/pkg/exhttp"
	"pay/pkg/param"
	"strings"
)

func (client *Client) query(payload param.Params) (param.Params, error) {
	var (
		httpResp = param.Params{}
		resp     = param.Params{}
		url      = client.getUrl()
		sign     string
		err      error
	)
	payload["method"] = QueryMethod
	payload["biz_content"] = helper.JsonMarshal(payload["biz_content"])
	if sign, err = client.generateSign(payload); err != nil {
		return nil, err
	}
	payload["sign"] = sign
	if httpResp, err = exhttp.Post(url, exhttp.TypeUrlencoded, strings.NewReader(payload.ToUrlValue()), nil); err != nil {
		return nil, err
	}
	resp, sign, err = client.getRespAndSignFromHttpResp(httpResp, QueryMethod)
	if err != nil {
		return nil, err
	}
	if err = client.verifySign(resp, sign); err != nil {
		return nil, err
	}
	return resp, nil
}

func (client *Client) refundQuery(payload param.Params) (param.Params, error) {
	var (
		httpResp = param.Params{}
		resp     = param.Params{}
		url      = client.getUrl()
		sign     string
		err      error
	)
	payload["method"] = RefundQueryMethod
	payload["biz_content"] = helper.JsonMarshal(payload["biz_content"])
	if sign, err = client.generateSign(payload); err != nil {
		return nil, err
	}
	payload["sign"] = sign
	if httpResp, err = exhttp.Post(url, exhttp.TypeUrlencoded, strings.NewReader(payload.ToUrlValue()), nil); err != nil {
		return nil, err
	}
	resp, sign, err = client.getRespAndSignFromHttpResp(httpResp, RefundQueryMethod)
	if err != nil {
		return nil, err
	}
	if err = client.verifySign(resp, sign); err != nil {
		return nil, err
	}
	return resp, nil
}
