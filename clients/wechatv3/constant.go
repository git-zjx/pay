package wechatv3

const (
	SignatureMessageFormat string = "%s\n%s\n%d\n%s\n%s\n"
	AuthorizationFormat    string = "%s mchid=\"%s\",nonce_str=\"%s\",timestamp=\"%d\",serial_no=\"%s\",signature=\"%s\""
	AuthorizationType      string = "WECHATPAY2-SHA256-RSA2048"
	Host                   string = "https://api.mch.weixin.qq.com/"
	SandboxHost            string = "https://api.mch.weixin.qq.com/"
	AppPayMethod           string = "v3/pay/transactions/app"
	JsapiPayMethod         string = "v3/pay/transactions/jsapi"
	ScanPayMethod          string = "v3/pay/transactions/native"
	WapPayMethod           string = "v3/pay/transactions/h5"
	QueryMethod            string = "v3/pay/transactions/%s/%s?mchid=%s"
	CloseMethod            string = "v3/pay/transactions/out-trade-no/%s/close"
	RefundQueryMethod      string = "v3/refund/domestic/refunds/%s"
	RefundMethod           string = "v3/refund/domestic/refunds"
)
