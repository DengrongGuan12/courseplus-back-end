package main

import (
	"encoding/json"
	"fmt"
	"log"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/customs"
)

func init() {
	// LogLevel 是 Go SDK 提供的 debug 开关
	pingpp.LogLevel = 2
	//设置 API Key
	pingpp.Key = "sk_test_ibbTe5jLGCi5rzfH4OqPW9KC"
	//设置错误信息语言，默认是中文
	pingpp.AcceptLanguage = "zh-CN"

	//设置商户的私钥 记得在Ping++上配置公钥
	pingpp.AccountPrivateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAx2MktxcKBEqdYRi2IgYcupPQIN5cxgiBL5udCCBJBNBbXPaq
uOE1qspfhB1KUzHXATnCONiSzubLcBTnwi2tz0ErRCeJZSERRCpbKx4eu6b1neUT
Wkga7xpZxWONEvkmZo5Nlhf4fXRPUYnO/bdGCNGpQ/HSJfWLtzmhCqO1aJwVhcDm
DMYz4bTkZavhFBdVyXf/8n7UKylk03eymlKJ1swQpeFcxaKfzsk1mJU7mc93mCWj
aR+VWkNbw4AQHDyHgbzH+zYARzCluiy5hXdixGEP+iO4ZBk48rEs1hKTvGz1k+jh
LCdkdpBRjq0pK/htjA3Ce8pF2AJs+fgN6ZUumQIDAQABAoIBAFa4MEfRpXGoYjrQ
3KZ/sg8UKvmgvQkEuetS60GViSym0pXkUuyGRyk5S8HSW3lDvBe0X10KFRAYIXNm
JEa4R1hVJ9REveVWNIRJR83BE+zZ+QnrkDc8FTrZYyIO4lTWOHVyfxxA4Lrv02/L
WFPRWoyLY+tBSf1ohpPyZLCT81rDglT1Z4svX020y8tXvnQqQiOjl4q7Zu4b26HU
TQ463ntMEhM5u7y9MFcxGRaOpF/gARlMGqDu6T8h/oYMiOSLoXOuTR7B80yaX/Mj
RZfUBoZMb5thX9qBLQ7dYnTkwaxwerYPrYvQrW9vtsswZ5NeIbEmCZyorUe8DOmQ
hT1+HmECgYEA/iQERHhZKHXnP0gvhl/uEOGOvLjD5H1D6zClzOHMmOcIF5OuEQb0
VcSMV+8emN7SCp/b/LVgKa27Mla9eXm+EXABRFcI7qGYsYXfbCD7EYX3TaJSp/30
jyLBy+MsHCTEiLeylSh7kHqgTR8tKND8UIzXo9aM7JqwFqleeXGyh7MCgYEAyNiU
EUzyBAv9sui3ZgVYRiVvTilk2HVTY6u61/mMOLsTrX3eYQaqb4GRJJShJO9mmsxX
RHBEZQJvUqqF9PapOsyv8HKuF5+UP6svHnJo7sn9gCvV/h1HTHqzFcYSvUaXnrym
D/0Tthf8CDeuGp5UFWMoFZF14HTr1oQROGAASoMCgYA0bZmzxmAeSLR8CZhEUGX8
dYvMwxEmgfERA+gwbCSZJpA0zPKL8LNXPkT1nw7g2pbaOkBX0dMUxhJoQBy2grcD
QegBATOGhy/I76U32VXyN4DdMy96GJnrLXBtb2AaLjudOMhOnRtgouuO/W+DjBmB
RIz377sC1KafBjHHO/1ooQKBgDQqfJrZv2ppquVTKH9pF/pwMq68daL7JkOXERqT
iGYbwQqozJ+q2Y3Iu2gi6o/rVl0SggAWoM0TitKP0+dCQcYx7+imAK3GFv1KexyP
Xs3WzO8Dc7ti42fr3qPjJG7g7PSfzwoME5iSNjX0MFZdlT1Q2dJwS4uXEsJO3yIj
XS/9AoGBALRApgtUA7Odw4tjCLGvxXuLFnyRkg6hFqoXAP2j8H9bJDOlSSVwQTFd
ahbcIDtQJS57vXUGK2uspbFKLm1WCFzPVyuxDIW6oue/kO+YxxU3NA58zk8oaORq
eA3YvHc7ZmRjVnVkxnXjKofrL6jF5A+lXSXnXchrv2ZYI+1pOsIV
-----END RSA PRIVATE KEY-----
`
}

func ExampleCustoms_new() {
	params := &pingpp.CustomsParams{
		App:          "app_1Gqj58ynP0mHeX1q",
		Charge:       "ch_avbPyT4aDCe1r9W1eHfv1SuP",
		Trade_no:     "123456",
		Customs_code: "GUANGZHOU",
		Amount:       1,
		Is_split:     false,
		Extra: map[string]interface{}{
			"pay_account":   "123456",
			"certif_type":   "02",
			"customer_name": "customer name",
			"certif_id":     "62148502",
			"tax_amount":    10,
		},
	}

	customs, err := customs.New(params)
	if err != nil {
		log.Fatal(err)
		return
	}
	content, _ := json.Marshal(customs)
	fmt.Println(string(content))
}

func ExampleCustoms_get() {
	customs, err := customs.Get("cus_id")
	if err != nil {
		log.Fatal(err)
		return
	}
	content, _ := json.Marshal(customs)
	fmt.Println(string(content))
}

func main() {
	ExampleCustoms_new()
}
