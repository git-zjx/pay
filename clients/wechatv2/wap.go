package wechatv2

import (
	"net/url"
	"pay/pkg/param"
)

func (client *Client) wap(payload param.Params) (param.Params, error) {
	var (
		resp       = param.Params{}
		prePayResp = param.Params{}
		err        error
	)
	payload["trade_type"] = "MWEB"
	if prePayResp, err = client.prePay(payload); err != nil {
		return nil, err
	}
	if err = client.isSuccess(prePayResp); err != nil {
		return nil, err
	}
	mwebUrl := prePayResp["mweb_url"].(string)
	if client.config.ReturnUrl != "" {
		mwebUrl += "&redirect_url=" + url.QueryEscape(client.config.ReturnUrl)
	}
	resp["wap_url"] = mwebUrl
	return resp, nil
}
