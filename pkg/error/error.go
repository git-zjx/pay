package error

import "errors"

var (
	PayMethodErr = errors.New("pay method err")
	PayMethodNotMatchErr = errors.New("pay method not match this client")
	PayReturnParamFormatErr = errors.New("pay return param format err")
	PayReturnParamNotHaveSignErr = errors.New("pay return param not have sign")
	SignNotMatchErr = errors.New("sign not match err")
	CertNotFoundErr = errors.New("cert not found err")
	ReqInfoNotFoundErr = errors.New("req_info not found err")
)