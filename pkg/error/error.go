package error

import "errors"

var (
	PayMethodErr                 = errors.New("pay method err")
	PayMethodNotMatchErr         = errors.New("pay method not match this client")
	PayReturnParamFormatErr      = errors.New("pay return param format err")
	PayReturnParamNotHaveSignErr = errors.New("pay return param not have sign")
	SignNotMatchErr              = errors.New("sign not match err")
	SignNotFoundErr              = errors.New("sign not found err")
	SignFormatErr                = errors.New("sign format err")
	CodeNotFoundErr              = errors.New("code not found err")
	CodeFormatErr                = errors.New("code format err")
	CertNotFoundErr              = errors.New("cert not found err")
	ReqInfoNotFoundErr           = errors.New("req_info not found err")
	ReqInfoFormatErr             = errors.New("req_info format err")
	PrepayIdNotFoundErr          = errors.New("prepay_id not found err")
	PrepayIdFormatErr            = errors.New("prepay_id format err")
	MwebUrlNotFoundErr             = errors.New("mweb_url not found err")
	MwebUrlFormatErr             = errors.New("mweb_url format err")
)
