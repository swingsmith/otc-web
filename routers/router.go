package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/otc/otc-web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/eth/getId", &controllers.EthController{}, "get:GetId;post:Post")

	beego.Router("/eth/postWithdrawEthUsdt", &controllers.EthController{}, "post:PostWithdrawEthUsdt")
	beego.Router("/eth/getEthTransactionReceipt", &controllers.EthController{}, "get:GetEthTransactionReceipt")
	beego.Router("/eth/getEthAccountBalance", &controllers.EthController{}, "get:GetEthAccountBalance")
	beego.Router("/eth/getEthUsdtBalance", &controllers.EthController{}, "get:GetEthUsdtBalance")

	beego.Router("/trx/postWithdrawTrxUsdt", &controllers.TrxController{}, "post:PostWithdrawTrxUsdt")
	beego.Router("/trx/getTrxTransactionReceipt", &controllers.TrxController{}, "get:GetTrxTransactionReceipt")
	beego.Router("/trx/getTrxAccountBalance", &controllers.TrxController{}, "get:GetTrxAccountBalance")
	beego.Router("/trx/getTrxUsdtBalance", &controllers.TrxController{}, "get:GetTrxUsdtBalance")
}
