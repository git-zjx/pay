package http

import (
	"fmt"
	"io"
	"net/http"
	"pay/pkg/helper"
	"pay/pkg/param"
	"runtime"
	"strings"
	"time"
)

func Request(method string, url string, authorization string, reqBody interface{}) (param.Params, error) {
	var (
		httpResp     *http.Response
		httpRespBody []byte
		request      *http.Request
		resp         param.Params
		err          error
	)
	if request, err = http.NewRequest(method, url, strings.NewReader(helper.JsonMarshal(reqBody))); err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", generateUa())
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", authorization)
	client := &http.Client{
		Timeout: time.Second * 2,
	}
	if httpResp, err = client.Do(request); err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(httpResp.Body)
	if httpRespBody, err = io.ReadAll(httpResp.Body); err != nil {
		return nil, err
	}
	if err = helper.JsonUnMarshal(httpRespBody, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func generateUa() string {
	return fmt.Sprintf("Pay-Go/%s (%s) GO/%s", "1.0.0", runtime.GOOS, runtime.Version())
}