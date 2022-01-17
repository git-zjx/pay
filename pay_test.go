package pay

import (
	"fmt"
	"pay/clients/alipay"
	"pay/pkg/param"
	"testing"
)

func TestPay(t *testing.T) {
	config := new(alipay.Config)
	config.AppPrivateKey = "MIIEpQIBAAKCAQEAr3ocBwt/W2YaV/VvFxkZSJA6fghyjhH//AFrAfPP2rZ3NQnSpZFERMYUlLGc3AjQhOs8hJXmDJrW+esidOFebCpqZq3BYECanOLCFnudTI3WeEPKs3r8wy9wCXIQeMHbOtnGwDhw7iYA/UZPPkcAB8f0wyBsVK0Xos2GgEA92O+82YXb/tq3S6DfnSAkdGIIlQ53o3O8eJm7QrgjwB1DTLYEx0LOf8aM4nMyphQR4syORPQSV++6rXuVrwe0jxcbrd4T6XHNcEzQB/FQMTBPn6Be7AHurptzHsY/B/wni7UNP7GSdOdO7xv75WSYrfy/iY3Wmdxf/XhqSORY3D711QIDAQABAoIBAQCM1R3lcY7XVgzSh0KPcS9fk5G+UR1PdJbUNHcjbABn8oWd5bJP+1SlNayS4jGYTuK7qug5KO0nNKZQkixnTfEwMqKOoelPTMpKG5vV24QVSsjUYOQwRAbUyB/NFOSvZjaC9wGSiDnqiEnG/EThIK6fkBWa/Uy0cO9FVFocWHLKxdlMqb9AbQ6eStrMpCqjv1J1FcRtqeiy8uLeatDSpoDneQStlzWDrSscdz2ZQD/2ppmsPqM2vpC0g7Kbrdcc3IqBe4ZrfRCB0ZDp/IO9zaT51KpYiA0hy7hLKZMZPCtobSk//ppXumD++lcpWzKD0X9e+Dfr13gnpTCPgZQE9suBAoGBAN8ZUW0ucnCN9mJU+9iGQT/dnpMJmJxRpFCW1tYgRJ6VC8s30a2mDtym/Cw9VcGs1fuGGArlQEuUYGN0VXmfajQ+Bg6HY9S/gTBaiEuNAcvK8S0PNpMJdlUdUA8dCnKRc+I+S6u7NgtmLXey/jfTbzT4xBFkB8/YOvgGFI1m5EhRAoGBAMla6k6pQTywSLQ0ZWcKVoaPlYmbviUkQ6lBi4xsIbXofftshB/Gf7Z+P+jdCew58RTgypuEoGM3koW9+rURZXo9B/ZlqyFbGFXib2ywUsBrSLjNUKZxii/klLZfWXN2orJVpfrdJBNGmyRna6At9xzb10bwE/PHiC9BzHVO2fhFAoGBALI8QedeQiNV509r4cB8kch6P+PsuLW6K/IOcBilsuyW2tNCBwwaLKlv5utZHRgcAuBtouuhd5pqMg+Cs371Mx4Fp0UYOVOQo5+D1Hu3bYXo3oFHNCyIVLdvMbTBWMVrGw/XARF0AZtdyFlm8N6c1q2VSN8z8WHFuGbKRMUrPJnhAoGBAIA+nI9fM8LAkH3eBVu8dOGdX+PWQyQK1eFectAMKuheXKcfNYO7fKox/OiGqARB3y+qAMFOloy632K5Xo2mt9hEOOcRWA6Vo4lIACncn8gYTKgPdLeeByJ71s/VKPbmb7df36hI4uo4BSYJjL3nqMVDq/htfne89RsMMHnir+d1AoGAIhJ8YfeJLkUZWkrUVeEYuUYTQlNRxFJUmSkTgjXPe/4m9ey++6VLRL+g+saRHs31rYTTP6bARnmwlWySka1DWemX9H9B0GFRwGYBoFlMDm9k0HyBfyY7dOAsz27guzyoIsqTzerGPb/0U4/imVoOmA6FjFzbV8GL0OThp111mAw="
	config.AlipayPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1zH5wGT/QA4R95wzS5JH1eUHEAqjBa+YKkiEH+6p7oY1/2NKBzpYRXhIPf2NsYQNKcZobmvD4tArOjZ2SSngtYD2EOvkGDlK2aFNHwE+Nyk2wA3FOaa2+A+8I2appxr1YlEqPtoU4+lJ7LbRx6AiQRSkxQ8DIM+JWjh5hqAs/72dhCGDBbwb/LkGNMsf49j5YviTPmgLzvrWJZpj4LCrAHvNu3HTCF+pr+AYk9Nhqo/Aiwqq7GPFcw1lRUi7KzNLLP7wH9ZFU5EBz+pD3QhbmhRc7VNjSVCEcsDWioQI+lazwRlrqS5djxkRyq9kVyyVyCNSQzrwgZ3cSsFk+1t4zwIDAQAB"
	config.AppId = "2016091700533850"
	config.SignType = "RSA2"
	config.NotifyUrl = "https://baidu.com"
	config.ReturnUrl = "https://baidu.com"
	config.Sandbox = true
	req := make(param.Params)
	req["out_trade_no"] = "123"
	req["total_amount"] = "123"
	req["subject"] = "123"
	pay := NewPay(config)
	res, err := pay.Wap(req)
	fmt.Println(err)
	fmt.Println(res)
}
