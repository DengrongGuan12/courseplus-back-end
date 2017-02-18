package common

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/pingplusplus/pingpp-go/pingpp"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/base"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/err"
	"io/ioutil"
	"net/http"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/models"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/util"
	"github.com/goinggo/mapstructure"
	"github.com/pingplusplus/pingpp-go/pingpp/charge"
)

type PayController struct {
	base.BaseController
}

// @description 购买单份课时
// @router /period [get]
func (c *PayController) PayPeriod(){
	periodId, e := c.GetInt("period_id")
	if e != nil || periodId <=0 {
		c.RetError(err.ErrInputData, "period_id")
	}
	channel := c.GetString("channel")
	if channel == "" {
		c.RetError(err.ErrInputData, "channel")
	}

	//获取period价格信息
	order := models.GetOrderByObjectId("PERIOD", periodId)
	if order == nil{
		c.RetError(err.ErrNil, "order")
	}
	orderId := order.Id
	orderno := order.Orderno
	amount := uint64(models.GetPeriodById(periodId).Price)
	resultCharge := util.Pay(channel, orderId, "PEROID", orderno, amount)

	result := make(map[string]interface{})
	result["data"] = resultCharge
	c.Data["json"] = result
	c.ServeJSON()
}

// @description 购买单份资料
// @router /document [get]
func (c *PayController) PayDocument(){
	documentId, e := c.GetInt("document_id")
	if e != nil || documentId <=0 {
		c.RetError(err.ErrInputData, "document_id")
	}
	channel := c.GetString("channel")
	if channel == "" {
		c.RetError(err.ErrInputData, "channel")
	}

	orderId := models.GetOrderByObjectId("DOC", documentId).Id
	//获取period价格信息
	amount := uint64(models.GetDocumentById(documentId).Price)
	resultCharge := util.Pay(channel, orderId, "DOC", "",amount)

	////添加订单order信息
	//var order models.Order

	result := make(map[string]interface{})
	result["data"] = resultCharge
	c.Data["json"] = result
	c.ServeJSON()
}

// @description 回答问题支付
// @router /document [get]
func (c *PayController) PayAnswer(){
	questionId, e := c.GetInt("question_id")
	if e != nil || questionId <=0 {
		c.RetError(err.ErrInputData, "question_id")
	}
	channel := c.GetString("channel")
	if channel == "" {
		c.RetError(err.ErrInputData, "channel")
	}

	orderId := models.GetOrderByObjectId("PEEK", questionId).Id
	//获取period价格信息
	resultCharge := util.Pay(channel, orderId, "PEEK","", 100)

	result := make(map[string]interface{})
	result["data"] = resultCharge
	c.Data["json"] = result
	c.ServeJSON()
}

// @router /getCharge [get]
func (c *PayController) GetCharge() {
	orderId ,e := c.GetInt("order_id")
	if e != nil && orderId <= 0 {
		c.RetError(err.ErrInputData, "order_id")
	}
	//根据order_id获取charge_id
	order := models.GetOrderByOrderId(orderId)
	if order == nil{
		c.RetError(err.ErrNil, "order")
	}
	//根据charge_id获取charge对象
	ch, e := charge.Get(order.Charge_id)
	if e != nil {
		c.RetError(err.ErrGetCharge, "charge_id")
	}

	result := make(map[string]interface{})
	result["data"] = ch
	c.Data["json"] = result
	c.ServeJSON()
}

// @router /payOrder
func (c *PayController) PayOrder(){
	orderId, e := c.GetInt("order_id")
	if e != nil || orderId <=0 {
		c.RetError(err.ErrInputData, "order_id")
	}
	channel := c.GetString("channel")
	if channel == "" {
		c.RetError(err.ErrInputData, "channel")
	}

	order := models.GetOrderByOrderId(orderId)
	//获取period价格信息
	amount := uint64(order.Price)
	resultCharge := util.Pay(channel, orderId, "DOC", "",amount)

	result := make(map[string]interface{})
	result["data"] = resultCharge
	c.Data["json"] = result
	c.ServeJSON()
}


// @router /webhook [post]
func (c *PayController) Webhook() {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Ctx.Request.Body)

	//示例 - 签名在头部信息的 x-pingplusplus-signature 字段
	//signed := `BX5sToHUzPSJvAfXqhtJicsuPjt3yvq804PguzLnMruCSvZ4C7xYS4trdg1blJPh26eeK/P2QfCCHpWKedsRS3bPKkjAvugnMKs+3Zs1k+PshAiZsET4sWPGNnf1E89Kh7/2XMa1mgbXtHt7zPNC4kamTqUL/QmEVI8LJNq7C9P3LR03kK2szJDhPzkWPgRyY2YpD2eq1aCJm0bkX9mBWTZdSYFhKt3vuM1Qjp5PWXk0tN5h9dNFqpisihK7XboB81poER2SmnZ8PIslzWu2iULM7VWxmEDA70JKBJFweqLCFBHRszA8Nt3AXF0z5qe61oH1oSUmtPwNhdQQ2G5X3g==`
	signed := c.Ctx.Request.Header.Get("X-Pingplusplus-Signature")

	//示例 - 待验签的数据
	//data := `{"id":"evt_eYa58Wd44Glerl8AgfYfd1sL","created":1434368075,"livemode":true,"type":"charge.succeeded","data":{"object":{"id":"ch_bq9IHKnn6GnLzsS0swOujr4x","object":"charge","created":1434368069,"livemode":true,"paid":true,"refunded":false,"app":"app_vcPcqDeS88ixrPlu","channel":"wx","order_no":"2015d019f7cf6c0d","client_ip":"140.227.22.72","amount":100,"amount_settle":0,"currency":"cny","subject":"An Apple","body":"A Big Red Apple","extra":{},"time_paid":1434368074,"time_expire":1434455469,"time_settle":null,"transaction_no":"1014400031201506150354653857","refunds":{"object":"list","url":"/v1/charges/ch_bq9IHKnn6GnLzsS0swOujr4x/refunds","has_more":false,"data":[]},"amount_refunded":0,"failure_code":null,"failure_msg":null,"metadata":{},"credential":{},"description":null}},"object":"event","pending_webhooks":0,"request":"iar_Xc2SGjrbdmT0eeKWeCsvLhbL"}`
	data := buf.String()
	// 请从 https://dashboard.pingxx.com 获取「Ping++ 公钥」
	publicKey, err := ioutil.ReadFile("../ca/pingpp_rsa_public_key.pem")
	if err != nil {
		fmt.Errorf("read failure: %v", err)
	}

	//base64解码再验证
	decodeStr, _ := base64.StdEncoding.DecodeString(signed)
	errs := pingpp.Verify([]byte(data), publicKey, decodeStr)
	if errs != nil {
		fmt.Println(errs)
		return
	} else {
		fmt.Println("success")
	}

	webhook, err := pingpp.ParseWebhooks(buf.Bytes())
	//fmt.Println(webhook.Type)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(c.Ctx.ResponseWriter, "fail")
		return
	}

	if webhook.Type == "charge.succeeded" {
		// TODO your code for charge
		var returnCharge pingpp.Charge
		mapstructure.Decode(webhook.Data.Object, &returnCharge)
		//returnCharge := pingpp.Charge{}(webhook.Data)
		order_id := returnCharge.Metadata["order_id"].(int)
		order := models.GetOrderByOrderId(order_id)
		payTime := util.GetCurrentTime()
		order.Pay_time = &payTime
		models.UpdateOrder(*order)

		order_type := returnCharge.Metadata["order_type"].(string)
		switch order_type {
		case "COURSE":
			courseId := order.Object_id
			
		case "ASK":
			questionId := order.Object_id
			//添加至user_buy_record
			record := models.User_buy_record{
				User: order.User,
				Type: order.Type,
				Object_id: questionId,
				Pay_time: *order.Pay_time,
			}
			models.AddUserBuyRecord(record)
			//修改ask状态为PAIED
			models.PayQuestionById(questionId)
		case "DOC_PACK":

		case "DOC", "PEROID", "PEEK":
			record := models.User_buy_record{
				User: order.User,
				Type: order.Type,
				Object_id: order.Object_id,
				Pay_time: *order.Pay_time,
			}
			models.AddUserBuyRecord(record)
		default:

		}
		c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	} else if webhook.Type == "refund.succeeded" {
		// TODO your code for refund
		c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	} else {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	}
}
