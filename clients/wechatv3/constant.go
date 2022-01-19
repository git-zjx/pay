package wechatv3

const (
	SignatureMessageFormat string = "%s\n%s\n%d\n%s\n%s\n"
	AuthorizationFormat    string = "%s mchid=\"%s\",nonce_str=\"%s\",timestamp=\"%d\",serial_no=\"%s\",signature=\"%s\""
	AuthorizationType      string = "WECHATPAY2-SHA256-RSA2048"
	Host                   string = "https://api.mch.weixin.qq.com/v3/"
	SandboxHost            string = "https://api.mch.weixin.qq.com/v3/"
	AppPayMethod           string = "pay/transactions/app"
	JsapiPayMethod         string = "pay/transactions/jsapi"
	ScanPayMethod          string = "pay/transactions/native"
	WapPayMethod           string = "pay/transactions/h5"
	QueryMethod            string = "pay/transactions/%s/%s"
	CloseMethod            string = "pay/transactions/out-trade-no/%s/close"
	CancelMethod           string = "secapi/pay/reverse"
	RefundQueryMethod      string = "pay/refundquery"
	RefundMethod           string = "secapi/pay/refund"
)
