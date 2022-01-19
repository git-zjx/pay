package pay

import (
	"pay/clients/alipay"
	"pay/clients/wechatv2"
	"pay/contracts"
	"pay/pkg/constant"
	"pay/pkg/param"
)

type Pay struct {
	client contracts.Client
}

func NewPay(config interface{}) *Pay {
	pay := new(Pay)
	switch config.(type) {
	case *alipay.Config:
		c, ok := config.(*alipay.Config)
		if !ok {
			return nil
		}
		pay.client = alipay.NewClient(c)
	case *wechatv2.Config:
		c, ok := config.(*wechatv2.Config)
		if !ok {
			return nil
		}
		pay.client = wechatv2.NewClient(c)
	}
	return pay
}

func (p *Pay) Web(request param.Params) (param.Params, error) {
	return p.client.Pay(constant.WEB, request)
}
func (p *Pay) Wap(request param.Params) (param.Params, error) {
	return p.client.Pay(constant.WAP, request)
}
func (p *Pay) App(request param.Params) (param.Params, error) {
	return p.client.Pay(constant.APP, request)
}
func (p *Pay) Pos(request param.Params) (param.Params, error) {
	return p.client.Pay(constant.POS, request)
}
func (p *Pay) Scan(request param.Params) (param.Params, error) {
	return p.client.Pay(constant.SCAN, request)
}
func (p *Pay) Mini(request param.Params) (param.Params, error) {
	return p.client.Pay(constant.MINI, request)
}
func (p *Pay) Mp(request param.Params) (param.Params, error) {
	return p.client.Pay(constant.MP, request)
}
func (p *Pay) Find(request param.Params, isRefund bool) (param.Params, error) {
	return p.client.Find(request, isRefund)
}
func (p *Pay) Refund(request param.Params) (param.Params, error) {
	return p.client.Refund(request)
}
func (p *Pay) Close(request param.Params) (param.Params, error) {
	return p.client.Close(request)
}
func (p *Pay) Cancel(request param.Params) (param.Params, error) {
	return p.client.Cancel(request)
}
func (p *Pay) Verify(request param.Params, isRefund bool) (param.Params, error) {
	return p.client.Verify(request, isRefund)
}
func (p *Pay) Success() {
	p.client.Success()
}
