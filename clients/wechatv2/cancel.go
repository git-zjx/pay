package wechatv2

import (
	"pay/pkg/helper"
	"pay/pkg/exhttp"
	"pay/pkg/param"
	"strings"
)

func (client *Client) cancel(payload param.Params) (param.Params, error) {
	var (
		httpResp = param.Params{}
		resp     = param.Params{}
		url      = client.getUrl(CancelMethod)
		sign     string
		err      error
	)
	delete(payload, "spbill_create_ip")
	if sign, err = client.generateSign(payload); err != nil {
		return nil, err
	}
	payload["sign"] = sign
	if httpResp, err = exhttp.Post(url, exhttp.TypeXML, strings.NewReader(helper.XmlMarshal(payload)), nil); err != nil {
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
