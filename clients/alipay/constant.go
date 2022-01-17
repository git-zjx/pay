package alipay

const (
	Host              string = "https://openapi.alipay.com/gateway.do"
	SandboxHost       string = "https://openapi.alipaydev.com/gateway.do"
	PayMethod         string = "alipay.trade.pay"
	PreCreateMethod   string = "alipay.trade.precreate"
	QueryMethod       string = "alipay.trade.query"
	CancelMethod      string = "alipay.trade.cancel"
	CloseMethod       string = "alipay.trade.close"
	AppPayMethod      string = "alipay.trade.app.pay"
	MiniPayMethod     string = "alipay.trade.create"
	WapPayMethod      string = "alipay.trade.wap.pay"
	PagePayMethod     string = "alipay.trade.page.pay"
	RefundQueryMethod string = "alipay.trade.fastpay.refund.query"
	RefundMethod      string = "alipay.trade.refund"
)
