package main

import (
	"encoding/json"
	"fmt"
	"log"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/card"
)

func init() {
	// LogLevel 是 Go SDK 提供的 debug 开关
	pingpp.LogLevel = 2
	//设置 API Key
	pingpp.Key = "sk_test_ibbTe5jLGCi5rzfH4OqPW9KC"
	//获取 SDK 版本
	fmt.Println("Go SDK Version:", pingpp.Version())
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

func ExampleCard_new() {

	cus_id := "cus_ALeWGZ8lsN9Czk"

	sms_code := make(map[string]interface{})
	sms_code["id"] = "sms_BDIQG8JZnQhUN1hAbwYubhP3"
	sms_code["code"] = "123222"
	param := &pingpp.CardParams{
		Source:   "tok_AMBKETCThoW7nYUBgpnvhwfu",
		Sms_code: sms_code,
	}

	card, err := card.New(cus_id, param)
	if err != nil {
		errs, _ := json.Marshal(err)
		fmt.Println(string(errs))
		log.Fatal(err)
		return
	}
	fmt.Println(card)

}

func ExampleCard_get() {
	card, err := card.Get("cus_ALeWGZ8lsN9Czk", "card_ALeWQUNinv0SHJhsCjBbJ29q")
	if err != nil {
		log.Fatal(err)
	}
	cardstring, _ := json.Marshal(card)
	fmt.Println("12343556578")
	log.Printf("%v\n", string(cardstring))
}

func ExampleCard_list() {

	params := &pingpp.CardListParams{}
	params.Filters.AddFilter("limit", "", "3")
	//设置是不是只需要之前设置的 limit 这一个查询参数
	params.Single = true
	cus_id := "cus_ALeWGZ8lsN9Czk"
	i := card.List(cus_id, params)
	for i.Next() {
		c := i.Card()
		ch, _ := json.Marshal(c)
		fmt.Println(string(ch))
	}
	// fmt.Println(i)
}

func ExampleCard_delete() {
	card, err := card.Delete("cus_ALeWGZ8lsN9Czk", "card_ALeWQUNinv0SHJhsCjBbJ29q")
	if err != nil {
		log.Fatal(err)
	}
	cardstring, _ := json.Marshal(card)
	fmt.Println("12343556578")
	log.Printf("%v\n", string(cardstring))
}

func main() {
	//ExampleCard_new()
	ExampleCard_get()
	ExampleCard_list()
	ExampleCard_delete()
}