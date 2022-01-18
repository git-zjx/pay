package http

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	netHttp "net/http"
	"pay/pkg/helper"
	"pay/pkg/param"
	"time"
)

const TypeXML = "application/xml"
const TypeUrlencoded = "application/x-www-form-urlencoded"

func Post(url, contentType string, reqBody io.Reader, tlsConfig *tls.Config) (param.Params, error) {
	var (
		httpResp     *netHttp.Response
		httpRespBody []byte
		resp         param.Params
		err          error
	)
	client := &netHttp.Client{
		Timeout: time.Second * 2,
	}
	if tlsConfig != nil {
		client.Transport = &netHttp.Transport{TLSClientConfig: tlsConfig, DisableKeepAlives: true, Proxy: netHttp.ProxyFromEnvironment}
	}
	if httpResp, err = client.Post(url, contentType, reqBody); err != nil {
		return nil, err
	}
	if httpResp.StatusCode != netHttp.StatusOK {
		return nil, errors.New(fmt.Sprintf("http code: %d", httpResp.StatusCode))
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(httpResp.Body)
	if httpRespBody, err = io.ReadAll(httpResp.Body); err != nil {
		return nil, err
	}
	if contentType == TypeUrlencoded {
		if err = helper.JsonUnMarshal(httpRespBody, &resp); err != nil {
			return nil, err
		}
	} else if contentType == TypeXML {
		if err = helper.XmlUnmarshal(httpRespBody, &resp); err != nil {
			return nil, err
		}
	}
	return resp, nil
}
