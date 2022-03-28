package alipay

import (
	"pay/pkg/exhttp"
	"pay/pkg/helper"
	"pay/pkg/param"
	"strings"
)

func (client *Client) close(payload param.Params) (param.Params, error) {
	var (
		httpResp = param.Params{}
		resp     = param.Params{}
		url      = client.getUrl()
		sign     string
		err      error
	)
	payload["method"] = CloseMethod
	payload["biz_content"] = helper.JsonMarshal(payload["biz_content"])
	if sign, err = client.generateSign(payload); err != nil {
		return nil, err
	}
	payload["sign"] = sign
	if httpResp, err = exhttp.Post(url, exhttp.TypeUrlencoded, strings.NewReader(payload.ToUrlValue()), nil); err != nil {
		return nil, err
	}
	resp, sign, err = client.getRespAndSignFromHttpResp(httpResp, CloseMethod)
	if err != nil {
		return nil, err
	}
	if err = client.verifySign(resp, sign); err != nil {
		return nil, err
	}
	return resp, nil
}

func (client *Client) cancel(payload param.Params) (param.Params, error) {
	var (
		httpResp = param.Params{}
		resp     = param.Params{}
		url      = client.getUrl()
		sign     string
		err      error
	)
	payload["method"] = CancelMethod
	payload["biz_content"] = helper.JsonMarshal(payload["biz_content"])
	if sign, err = client.generateSign(payload); err != nil {
		return nil, err
	}
	payload["sign"] = sign
	if httpResp, err = exhttp.Post(url, exhttp.TypeUrlencoded, strings.NewReader(payload.ToUrlValue()), nil); err != nil {
		return nil, err
	}
	resp, sign, err = client.getRespAndSignFromHttpResp(httpResp, CancelMethod)
	if err != nil {
		return nil, err
	}
	if err = client.verifySign(resp, sign); err != nil {
		return nil, err
	}
	return resp, nil
}
