package controllers

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/otc/otc-web/models"
	ethService "github.com/otc/otc-web/service/eth_service"
	"github.com/otc/otc-web/utils/db_util"
	"github.com/otc/otc-web/utils/id_util"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
	"time"
)

type EthController struct {
	beego.Controller
}

//var usdtContractAddress = "0xdAC17F958D2ee523a2206206994597C13D831ec7"

func (c *EthController) GetId() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	id := c.GetString("id", "0")
	fmt.Println("你的ID是：" + id)
	//c.SetData("你的ID是："+id)
	//c.ServeJSON(false)
	c.Ctx.WriteString("你的ID是：" + id)
}

func (c *EthController) Post() {
	var obj = make(map[string]interface{})
	data := c.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &obj); err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	fmt.Println(obj)
	str, _ := json.Marshal(obj)
	c.Ctx.WriteString(string(str))
}

func (c *EthController) PostWithdrawEthUsdt() {
	var db = db_util.GetDB()
	var config models.Config
	var usdtContractAddress, _ = config.GetEthUsdtContractAddress()
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

		var prvKey, _ = config.GetEthUsdtWithdrawPrivateKey()
		var fromAddr, _ = config.GetEthUsdtWithdrawAddress()

		amount, _ := decimal.NewFromString(amountRst.String())
		//aa := a.Mul(decimal.NewFromInt(10).Pow(decimal.NewFromInt(6)))
		//
		ethPrice, _ := config.GetEthPrice()
		price, _ := decimal.NewFromString(ethPrice)
		gas, _ := ethService.GetTransferGasFee()
		gas = gas.Mul(decimal.NewFromFloat(1.3)).Div(decimal.New(1,18))

		//fee = fee.Mul(decimal.NewFromFloat(1.3)).Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(18))).Mul(price)
		fee := gas.Mul(price)
		fmt.Printf("fee为：%s\n", fee.String())

		tx, err := ethService.TransferToken(usdtContractAddress, prvKey, addrRst.String(), (amount.Sub(fee)).Mul(decimal.New(1,6)).BigInt())
		if err != nil {
			fmt.Printf("PostWithdrawEthUsdt提现失败%s", err.Error())
			rst["code"] = 4000
			rst["msg"] = "提现失败"
			rst["tx"] = tx
			//rst["status"] = 2 //确认中
			str, _ := json.Marshal(rst)
			c.Ctx.WriteString(string(str))
			return
		}

		id := id_util.GenID()

		var withdraw = models.WithdrawRecord{OrderId: id, From: fromAddr, To: addrRst.String(), Gas: gas, Fee: fee, Value: amount.Sub(fee), Tx: tx, Type: "ETH_USDT", CreateTime: time.Now(), UpdateTime: time.Now(), Status: 0}
		if err := db.Debug().Model(&models.WithdrawRecord{}).Create(&withdraw).Error; err != nil {
			fmt.Printf("PostWithdrawEthUsdt插入提现记录失败%s", err.Error())
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

func (c *EthController) GetEthTransactionReceipt() {
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

	b, err := ethService.GetTransactionReceipt(tx)
	if err != nil {
		rst["code"] = 4000
		rst["msg"] = "失败"
	} else {
		rst["code"] = 200
		rst["msg"] = "成功"
	}

	if b {
		rst["status"] = 1
	} else {
		rst["status"] = 0
	}

	str, _ := json.Marshal(rst)
	c.Ctx.WriteString(string(str))
	return
}

func (c *EthController) GetEthAccountBalance() {
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

	b, err := ethService.GetAccountBalance(addr)
	if err != nil {
		rst["code"] = 4000
		rst["msg"] = "失败"
	} else {
		rst["code"] = 200
		rst["msg"] = "成功"
	}

	bb := decimal.NewFromBigInt(b, 0).Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(18)))

	rst["balance"] = bb.String()

	str, _ := json.Marshal(rst)
	c.Ctx.WriteString(string(str))
	return
}

func (c *EthController) GetEthUsdtBalance() {
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
	var usdtContractAddress, _ = config.GetEthUsdtContractAddress()

	b, err := ethService.GetTokenBalance(usdtContractAddress, addr)
	if err != nil {
		rst["code"] = 4000
		rst["msg"] = "失败"
	} else {
		rst["code"] = 200
		rst["msg"] = "成功"
	}

	bb := decimal.NewFromBigInt(b, 0).Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(6)))

	rst["balance"] = bb.String()

	str, _ := json.Marshal(rst)
	c.Ctx.WriteString(string(str))
	return
}
