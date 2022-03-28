package alipay

import (
	"fmt"
	"pay/pkg/helper"
	"pay/pkg/exhttp"
	"pay/pkg/param"
	"strings"
)

func (client *Client) mini(payload param.Params) (param.Params, error) {
	var (
		httpResp = param.Params{}
		resp     = param.Params{}
		url      = client.getUrl()
		sign     string
		err      error
	)
	payload["method"] = MiniPayMethod
	payload["biz_content"] = helper.JsonMarshal(payload["biz_content"])
	if sign, err = client.generateSign(payload); err != nil {
		return nil, err
	}
	payload["sign"] = sign
	fmt.Println(payload.ToUrlValue())
	if httpResp, err = exhttp.Post(url, exhttp.TypeUrlencoded, strings.NewReader(payload.ToUrlValue()), nil); err != nil {
		return nil, err
	}
	resp, sign, err = client.getRespAndSignFromHttpResp(httpResp, MiniPayMethod)
	if err != nil {
		return nil, err
	}
	if err = client.verifySign(resp, sign); err != nil {
		return nil, err
	}
	return resp, nil
}
