package error

import "errors"

var (
	PayMethodErr = errors.New("pay method err")
	PayMethodNotMatchErr = errors.New("pay method not match this client")
	PayReturnParamFormatErr = errors.New("pay return param format err")
	PayReturnParamNotHaveSignErr = errors.New("pay return param not have sign")
)