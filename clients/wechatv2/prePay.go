package wechatv2

import (
	"pay/pkg/helper"
	"pay/pkg/http"
	"pay/pkg/param"
	"strings"
)

func (client *Client) prePay(payload param.Params) (param.Params, error) {
	var (
		httpResp = param.Params{}
		resp     = param.Params{}
		url      = client.getUrl(PayMethod)
		sign     string
		err      error
	)
	if sign, err = client.generateSign(payload); err != nil {
		return nil, err
	}
	payload["sign"] = sign
	if httpResp, err = http.Post(url, http.TypeXML, strings.NewReader(helper.XmlMarshal(payload)), nil); err != nil {
		return nil, err
	}
	resp, sign, err = client.getRespAndSignFromHttpResp(httpResp)
	if err != nil {
		return nil, err
	}
	if err = client.verifySign(resp, sign); err != nil {
		return nil, err
	}
	return resp, nil
}
