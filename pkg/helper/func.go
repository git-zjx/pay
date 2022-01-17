package helper

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"strconv"
	"strings"
)

func Md5(str ...string) (string, error) {
	md5Init := md5.New()
	for _, v := range str {
		_, err := md5Init.Write([]byte(v))
		if err != nil {
			return "", err
		}
	}
	return hex.EncodeToString(md5Init.Sum(nil)), nil
}

func JsonMarshal(v interface{}) string {
	jsonB, _ := json.Marshal(v)
	return string(jsonB)
}

func JsonUnMarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func XmlMarshal(v interface{}) string {
	xmlB, _ := xml.Marshal(v)
	return string(xmlB)
}

func XmlUnmarshal(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}

func Sha256(str, secret string) (string, error) {
	key := []byte(secret)
	hash := hmac.New(sha256.New, key)
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode, nil
}

func PKCS7UnPadding(origData []byte) []byte {
	var bs []byte
	length := len(origData)
	unPaddingNumber := int(origData[length-1])
	if unPaddingNumber <= 16 {
		bs = origData[:(length - unPaddingNumber)]
	} else {
		bs = origData
	}
	return bs
}

func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}

func BuildResponseKeyFromMethod(method string) string {
	return strings.Replace(method, ".", "_", -1) + "_response"
}