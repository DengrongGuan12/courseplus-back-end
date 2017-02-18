package main

import (
	"encoding/json"
	"fmt"
	"log"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/customer"
)

func init() {
	// LogLevel 是 Go SDK 提供的 debug 开关
	pingpp.LogLevel = 3
	//设置 API Key
	pingpp.Key = "sk_test_ibbTe5jLGCi5rzfH4OqPW9KC"
	//获取 SDK 版本
	fmt.Println("Go SDK Version:", pingpp.Version())
	//设置错误信息语言，默认是中文
	pingpp.AcceptLanguage = "zh-CN"
}

func ExampleCustomer_new() {

	// cus_id := "cus_9K4KS8jLKq50yP"

	sms_code := make(map[string]interface{})
	sms_code["id"] = "sms_BDIQG8JZnQhUN1hAbwYubhP3"
	sms_code["code"] = "123222"

	param := &pingpp.CustomerParams{
		App:         "app_CyfHGK8eXPuL9uj9",
		Source:      "tok_ALeWEHQEp1wk9Ebep6a2EhVy",
		Sms_code:    sms_code,
		Description: "create test customer",
		Email:       "newcustomer@test.com",
	}

	cus, err := customer.New(param)
	if err != nil {
		errs, _ := json.Marshal(err)
		fmt.Println(string(errs))
		log.Fatal(err)
		return
	}
	fmt.Println(cus)

}

func ExampleCustomer_get() {
	cus, err := customer.Get("cus_ALeWGZ8lsN9Czk")
	if err != nil {
		log.Fatal(err)
	}
	cusstring, _ := json.Marshal(cus)
	fmt.Println("12343556578")
	log.Printf("%v\n", string(cusstring))
}

func ExampleCustomer_list() {

	params := &pingpp.CustomerListParams{}
	params.Filters.AddFilter("limit", "", "1")
	//设置是不是只需要之前设置的 limit 这一个查询参数
	params.Single = true
	i := customer.List(params)
	for i.Next() {
		c := i.Customer()
		ch, _ := json.Marshal(c)
		fmt.Println(string(ch))
	}
	// fmt.Println(i)
}

func ExampleCustomer_update() {

	cus_id := "cus_ALeWGZ8lsN9Czk"

	metadata := make(map[string]interface{})
	metadata["red"] = "yello"

	param := &pingpp.CustomerUpdateParams{
		// Default_source: "card_yTWjr1eTWznH8Ci1CC00SWf5",
		Description: "update test customer",
		Email:       "updatecustomer@test.com",
		Metadata:    metadata,
	}

	cus, err := customer.Update(cus_id, param)
	if err != nil {
		errs, _ := json.Marshal(err)
		fmt.Println(string(errs))
		log.Fatal(err)
		return
	}
	fmt.Println(cus)

}

func ExampleCustomer_delete() {
	cus, err := customer.Delete("cus_ALeWGZ8lsN9Czk")
	if err != nil {
		log.Fatal(err)
	}
	cusstring, _ := json.Marshal(cus)
	fmt.Println("12343556578")
	log.Printf("%v\n", string(cusstring))
}

func main() {
	ExampleCustomer_new()
}
