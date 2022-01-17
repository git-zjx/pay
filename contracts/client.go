package contracts

import "pay/pkg/param"

type Client interface {
	Pay(method string, request param.Params) (param.Params, error)
	Find(request param.Params, isRefund bool) (param.Params, error)
	Refund(request param.Params) (param.Params, error)
	Cancel(request param.Params) (param.Params, error)
	Close(request param.Params) (param.Params, error)
	Verify(request param.Params) (param.Params, error)
	Success()
}
