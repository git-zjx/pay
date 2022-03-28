package wechatv2

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/tls"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
	"pay/clients/wechatv2/pkg/sign"
	"pay/pkg/constant"
	 "pay/pkg/exerror"
	"pay/pkg/helper"
	"pay/pkg/param"
	"strings"
)

type Config struct {
	AppId      string
	MchId      string
	ApiKey     string
	CertPath   string
	KeyPath    string
	Pkcs12Path string
	ReturnUrl  string
	NotifyUrl  string
	SignType   string
	Sandbox    bool
}

type Client struct {
	config *Config
}

func NewClient(config *Config) *Client {
	client := new(Client)
	client.config = config
	return client
}

func (client *Client) Pay(method string, request param.Params) (param.Params, error) {
	payload := client.generatePayload(request)
	switch method {
	case constant.MP:
		return client.mp(payload)
	case constant.WAP:
		return client.wap(payload)
	case constant.APP:
		return client.app(payload)
	case constant.MINI:
		return client.mini(payload)
	case constant.POS:
		return client.pos(payload)
	case constant.SCAN:
		return client.scan(payload)
	default:
		return nil, exerror.PayMethodErr
	}
}

func (client *Client) Find(request param.Params, isRefund bool) (param.Params, error) {
	payload := client.generatePayload(request)
	if isRefund {
		return client.refundQuery(payload)
	}
	return client.query(payload)
}

func (client *Client) Refund(request param.Params) (param.Params, error) {
	payload := client.generatePayload(request)
	return client.refund(payload)
}

func (client *Client) Close(request param.Params) (param.Params, error) {
	payload := client.generatePayload(request)
	return client.close(payload)
}

func (client *Client) Cancel(request param.Params) (param.Params, error) {
	payload := client.generatePayload(request)
	return client.cancel(payload)
}

func (client *Client) Verify(request param.Params, isRefund bool) (param.Params, error) {
	var (
		encryptionB, bs []byte
		block           cipher.Block
		blockSize       int
		reqSign         string
		ok              bool
		err             error
	)
	if _, ok = request["sign"]; !ok {
		return nil, exerror.SignNotFoundErr
	}
	if isRefund {
		var reqInfoStr string
		if _, ok = request["req_info"]; !ok {
			return nil, exerror.ReqInfoNotFoundErr
		}
		if reqInfoStr, ok = request["req_info"].(string); !ok {
			return nil, exerror.ReqInfoFormatErr
		}
		if encryptionB, err = base64.StdEncoding.DecodeString(reqInfoStr); err != nil {
			return nil, err
		}
		md5Key, err := helper.Md5(client.config.ApiKey)
		if err != nil {
			return nil, err
		}
		key := strings.ToLower(md5Key)
		if len(encryptionB)%aes.BlockSize != 0 {
			return nil, errors.New("encryptedData is error")
		}
		if block, err = aes.NewCipher([]byte(key)); err != nil {
			return nil, err
		}
		blockSize = block.BlockSize()
		err = func(dst, src []byte) error {
			if len(src)%blockSize != 0 {
				return errors.New("crypto/cipher: input not full blocks")
			}
			if len(dst) < len(src) {
				return errors.New("crypto/cipher: output smaller than input")
			}
			for len(src) > 0 {
				block.Decrypt(dst, src[:blockSize])
				src = src[blockSize:]
				dst = dst[blockSize:]
			}
			return nil
		}(encryptionB, encryptionB)
		if err != nil {
			return nil, err
		}
		bs = helper.PKCS7UnPadding(encryptionB)
		var reqInfo param.Params
		if err = helper.XmlUnmarshal(bs, &reqInfo); err != nil {
			return nil, err
		}
		request["req_info"] = reqInfo
	}
	if reqSign, ok = request["sign"].(string); !ok {
		return nil, exerror.SignFormatErr
	}
	if err = client.verifySign(request, reqSign); err != nil {
		return nil, err
	}
	return request, nil
}

func (client *Client) Success() {
	fmt.Println("<xml><return_code><![CDATA[SUCCESS]]></return_code><return_msg><![CDATA[OK]]></return_msg></xml>")
}

func (client *Client) generateSign(params param.Params) (string, error) {
	var (
		signRes string
		err     error
	)
	if signRes, err = sign.Generate(params, client.config.SignType, client.config.ApiKey); err != nil {
		return signRes, err
	}
	return signRes, nil
}

func (client *Client) verifySign(params param.Params, retSign string) error {
	var (
		err error
	)
	if err = sign.Verify(params, client.config.SignType, client.config.ApiKey, retSign); err != nil {
		return err
	}
	return nil
}

func (client *Client) getRespAndSignFromHttpResp(httpResp param.Params) (param.Params, string, error) {
	var (
		retSign string
		ok      bool
		err     error
	)
	if err = client.isSuccess(httpResp); err != nil {
		return nil, "", err
	}
	if _, ok = httpResp["sign"]; !ok {
		return nil, "", exerror.PayReturnParamNotHaveSignErr
	}
	if retSign, ok = httpResp["sign"].(string); !ok {
		return nil, "", exerror.SignFormatErr
	}
	return httpResp, retSign, nil
}

func (client *Client) isSuccess(data param.Params) error {
	var (
		returnCode, resultCode string
		ok                     bool
	)
	if _, ok = data["return_code"]; !ok {
		return exerror.CodeNotFoundErr
	}
	if _, ok = data["result_code"]; !ok {
		return exerror.CodeNotFoundErr
	}
	if returnCode, ok = data["return_code"].(string); !ok {
		return exerror.CodeFormatErr
	}
	if resultCode, ok = data["result_code"].(string); !ok {
		return exerror.CodeFormatErr
	}
	if returnCode == "SUCCESS" && resultCode == "SUCCESS" {
		return nil
	}
	return errors.New(fmt.Sprintf("%s", data))
}

func (client *Client) getUrl(endPoint string) string {
	if client.config.Sandbox {
		return SandboxHost + endPoint
	}
	return Host + endPoint
}

func (client *Client) generatePayload(request param.Params) param.Params {
	request["appid"] = client.config.AppId
	request["mch_id"] = client.config.MchId
	request["sign_type"] = client.config.SignType
	request["notify_url"] = client.config.NotifyUrl
	request["spbill_create_ip"] = helper.LocalIp()
	request["nonce_str"] = helper.GenerateRandomString(32)
	return request
}

func (client *Client) generateCertificate() (*tls.Certificate, error) {
	var (
		certPem, keyPem []byte
		certificate     tls.Certificate
		err             error
	)
	if client.config.CertPath != "" && client.config.KeyPath != "" {
		if certPem, err = ioutil.ReadFile(client.config.CertPath); err != nil {
			return nil, err
		}
		if keyPem, err = ioutil.ReadFile(client.config.KeyPath); err != nil {
			return nil, err
		}
	} else if client.config.Pkcs12Path != "" {
		var pfxData []byte
		if pfxData, err = ioutil.ReadFile(client.config.Pkcs12Path); err != nil {
			return nil, err
		}
		blocks, err := pkcs12.ToPEM(pfxData, client.config.MchId)
		if err != nil {
			return nil, err
		}
		for _, b := range blocks {
			keyPem = append(keyPem, pem.EncodeToMemory(b)...)
		}
		certPem = keyPem
	} else {
		return nil, exerror.CertNotFoundErr
	}
	if certificate, err = tls.X509KeyPair(certPem, keyPem); err != nil {
		return nil, err
	}
	return &certificate, nil
}
