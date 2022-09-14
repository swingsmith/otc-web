package controllers

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/otc/otc-web/models"
	"github.com/otc/otc-web/tron/core"
	trxService "github.com/otc/otc-web/tron/service"
	"github.com/otc/otc-web/utils/db_util"
	"github.com/otc/otc-web/utils/id_util"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
	"time"
)

type TrxController struct {
	beego.Controller
}

//var usdtContractAddress = "0xdAC17F958D2ee523a2206206994597C13D831ec7"

func (c *TrxController) GetId() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	id := c.GetString("id", "0")
	fmt.Println("你的ID是：" + id)
	//c.SetData("你的ID是："+id)
	//c.ServeJSON(false)
	c.Ctx.WriteString("你的ID是：" + id)
}

func (c *TrxController) Post() {
	var obj = make(map[string]interface{})
	data := c.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &obj); err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	fmt.Println(obj)
	str, _ := json.Marshal(obj)
	c.Ctx.WriteString(string(str))
}

func (c *TrxController) PostWithdrawTrxUsdt() {
	var db = db_util.GetDB()
	var client = trxService.GetTrxClient()
	var config models.Config
	var usdtContractAddress, _ = config.GetTrxUsdtContractAddress()
	var rst = make(map[string]interface{})
	//var addrRst gjson.Result
	//var amountRst gjson.Result
	data := c.Ctx.Input.RequestBody
	jsonStr := string(data)

	if !gjson.Valid(jsonStr) {
		fmt.Println("error")
		rst["code"] = 3000
		rst["msg"] = "参数错误"
		//rst["status"] = 2 //确认中
		str, _ := json.Marshal(rst)
		c.Ctx.WriteString(string(str))
		return
	} else {
		fmt.Println("ok")
		addrRst := gjson.Get(jsonStr, `to_address`)
		amountRst := gjson.Get(jsonStr, "amount")
		fmt.Printf("地址为：%s", addrRst.String())

		fmt.Printf("收到的原文为：%s,toAddress为：%s  amount为：%s", jsonStr, addrRst.String(), amountRst.String())

		var prvKey, _ = config.GetTrxUsdtWithdrawPrivateKey()
		var fromAddr, _ = config.GetTrxUsdtWithdrawAddress()

		amount, _ := decimal.NewFromString(amountRst.String())
		//aa := a.Mul(decimal.NewFromInt(10).Pow(decimal.NewFromInt(6)))
		//
		trxPrice, _ := config.GetTrxPrice()
		price, _ := decimal.NewFromString(trxPrice)
		gas := decimal.NewFromInt(15)

		fee := gas.Mul(price)
		fmt.Printf("fee为：%s\n", fee.String())

		tx, err := client.TransferContract(prvKey, usdtContractAddress, addrRst.String(), (amount.Sub(fee)).Mul(decimal.New(1, 6)).IntPart(), 15*1000000)
		if err != nil {
			fmt.Printf("PostWithdrawTrxUsdt提现失败%s", err.Error())
			rst["code"] = 4000
			rst["msg"] = "提现失败"
			rst["tx"] = tx
			//rst["status"] = 2 //确认中
			str, _ := json.Marshal(rst)
			c.Ctx.WriteString(string(str))
			return
		}

		id := id_util.GenID()
		var withdraw = models.WithdrawRecord{OrderId: id, From: fromAddr, To: addrRst.String(), Gas: gas, Fee: fee, Value: amount.Sub(fee), Tx: tx, Type: "TRX_USDT", CreateTime: time.Now(), UpdateTime: time.Now(), Status: 0}
		if err := db.Debug().Model(&models.WithdrawRecord{}).Create(&withdraw).Error; err != nil {
			fmt.Printf("PostWithdrawTrxUsdt插入提现记录失败%s", err.Error())
		}

		rst["code"] = 200
		rst["msg"] = "成功"
		rst["tx"] = tx
		rst["gas"] = gas.String()
		rst["fee"] = fee.String()
		rst["value"] = amount.Sub(fee).String()
		//rst["status"] = 2 //确认中
		str, _ := json.Marshal(rst)
		c.Ctx.WriteString(string(str))
		return
	}
}

func (c *TrxController) GetTrxTransactionReceipt() {
	var client = trxService.GetTrxClient()
	var rst = make(map[string]interface{})
	tx := c.GetString("tx", "")
	if tx == "" {
		rst["code"] = 3000
		rst["msg"] = "参数不能为空"

		str, _ := json.Marshal(rst)
		c.Ctx.WriteString(string(str))
		return
	}
	fmt.Println("你的tx是：" + tx)

	trans, err := client.GetTransactionById(tx)
	if err != nil {
		rst["code"] = 4000
		rst["msg"] = "失败"
	} else {
		rst["code"] = 200
		rst["msg"] = "成功"
	}

	rets := trans.GetRet()
	if len(rets) < 1 || rets[0].ContractRet != core.Transaction_Result_SUCCESS {
		rst["status"] = 0
	} else {
		rst["status"] = 1
	}

	str, _ := json.Marshal(rst)
	c.Ctx.WriteString(string(str))
	return
}

func (c *TrxController) GetTrxAccountBalance() {
	var client = trxService.GetTrxClient()
	var rst = make(map[string]interface{})
	addr := c.GetString("address", "")
	if addr == "" {
		rst["code"] = 3000
		rst["msg"] = "参数不能为空"

		str, _ := json.Marshal(rst)
		c.Ctx.WriteString(string(str))
		return
	}
	fmt.Println("你的Address是：" + addr)

	b, err := client.GetBalanceByAddress(addr)
	if err != nil {
		rst["code"] = 4000
		rst["msg"] = "失败"
	} else {
		rst["code"] = 200
		rst["msg"] = "成功"
	}

	bb := b.Div(decimal.New(1, 6))

	rst["balance"] = bb.String()

	str, _ := json.Marshal(rst)
	c.Ctx.WriteString(string(str))
	return
}

func (c *TrxController) GetTrxUsdtBalance() {
	var client = trxService.GetTrxClient()
	var rst = make(map[string]interface{})
	addr := c.GetString("address", "")
	if addr == "" {
		rst["code"] = 3000
		rst["msg"] = "参数不能为空"

		str, _ := json.Marshal(rst)
		c.Ctx.WriteString(string(str))
		return
	}
	fmt.Println("你的Address是：" + addr)

	var config models.Config
	var usdtContractAddress, _ = config.GetTrxUsdtContractAddress()

	b, err := client.GetTrc20BalanceByAddress(usdtContractAddress, addr)
	if err != nil {
		rst["code"] = 4000
		rst["msg"] = "失败"
	} else {
		rst["code"] = 200
		rst["msg"] = "成功"
	}

	bb := b.Div(decimal.New(1, 6))

	rst["balance"] = bb.String()

	str, _ := json.Marshal(rst)
	c.Ctx.WriteString(string(str))
	return
}
