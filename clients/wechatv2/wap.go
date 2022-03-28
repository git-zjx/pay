package wechatv2

import (
	"net/url"
	 "pay/pkg/exerror"
	"pay/pkg/param"
)

func (client *Client) wap(payload param.Params) (param.Params, error) {
	var (
		resp       = param.Params{}
		prePayResp = param.Params{}
		mwebUrl    string
		ok         bool
		err        error
	)
	payload["trade_type"] = "MWEB"
	if prePayResp, err = client.prePay(payload); err != nil {
		return nil, err
	}
	if err = client.isSuccess(prePayResp); err != nil {
		return nil, err
	}
	if _, ok = prePayResp["mweb_url"]; !ok {
		return nil, exerror.MwebUrlNotFoundErr
	}
	if mwebUrl, ok = prePayResp["mweb_url"].(string); !ok {
		return nil, exerror.MwebUrlFormatErr
	}
	if client.config.ReturnUrl != "" {
		mwebUrl += "&redirect_url=" + url.QueryEscape(client.config.ReturnUrl)
	}
	resp["wap_url"] = mwebUrl
	return resp, nil
}
