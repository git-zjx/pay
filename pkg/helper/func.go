package helper

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"math/rand"
	"net"
	"sort"
	"strconv"
	"strings"
	"time"
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

func GenerateQueryStringExceptSign(param map[string]interface{}) string {
	var data []string
	for k, v := range param {
		if v != "" && k != "sign" {
			data = append(data, fmt.Sprintf(`%s=%s`, k, v))
		}
	}
	sort.Strings(data)
	return strings.Join(data, "&")
}

func GenerateRandomString(l int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	bytes := []byte(str)
	var result []byte = make([]byte, 0, l)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func LocalIp() string {
	var (
		err error
	)
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		ipNet, isIpNet := addr.(*net.IPNet)
		if isIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ""
}
