package eth_util

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

var _client *ethclient.Client

func init()  {
	var err error
	//_client, err = ethclient.Dial("https://mainnet.infura.io/v3/ce08f108306a4cd2ab2b6b19a34060d6")
	_client, err = ethclient.Dial("http://192.168.1.230:8888")
	if err != nil {
		fmt.Printf("以太坊連接失敗！%v",err)
	}
}

func GetEthClient() *ethclient.Client {
	return _client
}
