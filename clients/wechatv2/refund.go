package wechatv2

import (
	"crypto/tls"
	"pay/pkg/helper"
	"pay/pkg/exhttp"
	"pay/pkg/param"
	"strings"
)

func (client *Client) refund(payload param.Params) (param.Params, error) {
	var (
		httpResp    = param.Params{}
		resp        = param.Params{}
		url         = client.getUrl(RefundMethod)
		certificate *tls.Certificate
		sign        string
		err         error
	)
	if certificate, err = client.generateCertificate(); err != nil {
		return nil, err
	}
	if sign, err = client.generateSign(payload); err != nil {
		return nil, err
	}
	payload["sign"] = sign
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{*certificate},
		InsecureSkipVerify: true,
	}
	if httpResp, err = exhttp.Post(url, exhttp.TypeXML, strings.NewReader(helper.XmlMarshal(payload)), tlsConfig); err != nil {
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
