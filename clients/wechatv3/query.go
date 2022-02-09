package wechatv3

import (
	"crypto/rsa"
	"errors"
	"fmt"
	netHttp "net/http"
	"pay/clients/wechatv3/pkg/http"
	"pay/pkg/param"
)

func (client *Client) query(request param.Params) (param.Params, error) {
	var (
		resp          = param.Params{}
		privateKey    *rsa.PrivateKey
		authorization string
		queryString   string
		err           error
	)
	if _, ok := request["transaction_id"]; ok {
		transactionId, ok := request["transaction_id"].(string)
		if !ok {
			return nil, errors.New("param error, transaction_id must string")
		}
		queryString = fmt.Sprintf(QueryMethod, "out_trade_no", transactionId, client.config.MchId)
	} else if _, ok := request["out_trade_no"]; ok {
		outTradeNo, ok := request["out_trade_no"].(string)
		if !ok {
			return nil, errors.New("param error, out_trade_no must string")
		}
		queryString = fmt.Sprintf(QueryMethod, "out_trade_no", outTradeNo, client.config.MchId)
	} else {
		return nil, errors.New("param error, need transaction_id or out_trade_no")
	}
	if privateKey, err = client.generatePrivateKey(); err != nil {
		return nil, err
	}
	if authorization, err = client.generateAuthorizationHeader(nil, netHttp.MethodGet, queryString, privateKey); err != nil {
		return nil, err
	}
	if resp, err = http.Request(netHttp.MethodGet, client.getUrl(queryString), authorization, nil); err != nil {
		return nil, err
	}
	if err = client.isSuccess(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (client *Client) refundQuery(request param.Params) (param.Params, error) {
	var (
		resp          = param.Params{}
		privateKey    *rsa.PrivateKey
		authorization string
		queryString   string
		err           error
	)
	if _, ok := request["out_refund_no"]; ok {
		outRefundNo, ok := request["out_refund_no"].(string)
		if !ok {
			return nil, errors.New("param error, out_refund_no must string")
		}
		queryString = fmt.Sprintf(RefundQueryMethod, outRefundNo)
	} else {
		return nil, errors.New("param error, need out_refund_no")
	}
	if privateKey, err = client.generatePrivateKey(); err != nil {
		return nil, err
	}
	if authorization, err = client.generateAuthorizationHeader(nil, netHttp.MethodGet, queryString, privateKey); err != nil {
		return nil, err
	}
	if resp, err = http.Request(netHttp.MethodGet, client.getUrl(queryString), authorization, nil); err != nil {
		return nil, err
	}
	if err = client.isSuccess(resp); err != nil {
		return nil, err
	}
	return resp, nil
}
