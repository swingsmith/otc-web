package service

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/sha3"
	"log"
	"math"
	"math/big"
	"github.com/otc/otc-web/utils/eth_util"
	"strconv"

	//"strconv"
	token "github.com/otc/otc-web/contracts_erc20"
	//	token "ChainBackend.com/" // for demo
)

var client *ethclient.Client

func init() {
	//var err error
	//client, err = ethclient.Dial("http://192.168.1.207:8888")
	//if err != nil {
	//	//log.Fatal(err)
	//}
	client = eth_util.GetEthClient()
}

func GetNowBlockHeight() (int64, error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		fmt.Printf("GetNowBlockHeight出错，error: %v", err)
		return 0, err
	}
	fmt.Println(header.Number.String()) // 5671744
	h,err := strconv.ParseInt(header.Number.String(),10,64)
	return h, nil
}

func GetNewAccount() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		//log.Fatal(err)
		return "", "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("新的私钥为：", hexutil.Encode(privateKeyBytes)[2:]) // 0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		//log.Fatal("error casting public key to ECDSA")
		return "", "", err
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("新的地址为：", address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	return address, hexutil.Encode(privateKeyBytes)[2:], nil
}

func PrivateKeyToAddress(prvKey string) (string, error) {
	privateKey, err := crypto.HexToECDSA(prvKey)
	if err != nil {
		//log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		//log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return fromAddress.Hex(), err
}

func GetAccountBalance(address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	fmt.Println(balance) // 25893860171173005034

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041
	return balance, nil
}

func SendETH(prvKey string, toAddr string, amount *big.Int) (string, error) {
	privateKey, err := crypto.HexToECDSA(prvKey)
	if err != nil {
		//log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		//log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	toAddress := common.HexToAddress(toAddr)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		//log.Fatal(err)
	}

	//value := big.NewInt(1000000000000000000) // in wei (1 eth)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		//log.Fatal(err)
	}
	fmt.Printf("GasPrice: %s\n", gasPrice.String())

	//toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	//gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	//	From: 		fromAddress,
	//	To: 		&toAddress,
	//	GasPrice: 	gasPrice,
	//	Data: 		data,
	//})
	//fmt.Printf("GasLimit: %d", gasLimit)

	tx := types.NewTransaction(nonce, toAddress, amount, 21000, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		//log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		//log.Fatal(err)
		return "SendEthFail!", err
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

	return signedTx.Hash().Hex(), nil
}

func GetSendETHGasFee() (decimal.Decimal, error) {
	//fromAddr := "0xaDc4806C2e31EB540Dd793980473A935fa960274"
	//toAddr := "0xCF52B0A8316eC79A8659945A4978898213AEC4fF"
	////amount := new(big.Int).SetInt64(1)
	//
	//fromAddress := common.HexToAddress(fromAddr)
	//toAddress := common.HexToAddress(toAddr)
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	//if err != nil {
	//	//log.Fatal(err)
	//}

	//value := big.NewInt(1000000000000000000) // in wei (1 eth)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		//log.Fatal(err)
	}

	//toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	//var data []byte
	//gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	//	From: 		fromAddress,
	//	To: 		&toAddress,
	//	GasPrice: 	gasPrice,
	//	Data: 		data,
	//})


	fee := decimal.NewFromBigInt(gasPrice,0).Mul(decimal.NewFromInt(21000))

	return fee, nil
}

////token
//func  DeployToken(prvKey string, name string, symbol string, decimals uint8,totalSupply *big.Int ,equityMultiplier *big.Int) (string, string, error) {
//	privateKey, err := crypto.HexToECDSA(prvKey)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	publicKey := privateKey.Public()
//	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//	if !ok {
//		log.Fatal("error casting public key to ECDSA")
//	}
//
//	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
//	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	gasPrice, err := client.SuggestGasPrice(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	chainID, err := client.NetworkID(context.Background())
//	if err != nil {
//		//log.Fatal(err)
//	}
//
//	auth, err := bind.NewKeyedTransactorWithChainID(privateKey,chainID)
//	auth.Nonce = big.NewInt(int64(nonce))
//	auth.Value = big.NewInt(0)     // in wei
//	auth.GasLimit = uint64(3000000) // in units
//	auth.GasPrice = gasPrice
//
//	address, tx, _, err := token.DeployToken(auth, client, name,symbol,decimals,totalSupply,equityMultiplier)
//	if err != nil {
//		//log.Fatal(err)
//		return "", "", err
//	}
//
//	fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
//	fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0
//
//	return tx.Hash().Hex(), address.Hex(), nil
//}

func GetName(tokenAddr string) (string, error) {
	// Golem (GNT) Address
	tokenAddress := common.HexToAddress(tokenAddr)
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		//log.Fatal(err)
	}

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		//log.Fatal(err)
		return "GetName Error", err
	}

	fmt.Printf("name: %s\n", name)

	return name, nil
}

func GetDecimals(tokenAddr string) (uint8, error) {
	// Golem (GNT) Address
	tokenAddress := common.HexToAddress(tokenAddr)
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		//log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		//log.Fatal(err)
		return 0, err
	}

	fmt.Printf("decimals: %s\n", decimals)

	return decimals, nil
}

func GetTokenBalance(tokenAddr string, addr string) (*big.Int, error) {
	// Golem (GNT) Address
	tokenAddress := common.HexToAddress(tokenAddr)
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		//log.Fatal(err)
	}

	address := common.HexToAddress(addr)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}

	fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"

	decimals, err := GetDecimals(tokenAddr)
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"

	return bal, nil
}

func TransferToken(tokenAddr string, prvKey string, toAddr string, amount *big.Int) (string, error) {

	privateKey, err := crypto.HexToECDSA(prvKey)
	if err != nil {
		//log.Fatal(err)
		fmt.Printf("私钥转换失败：%s", prvKey)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		//log.Fatal("error casting public key to ECDSA")
	}

	fmt.Printf("crypto.PubkeyToAddress(*publicKeyECDSA) %s", publicKeyECDSA)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("FromAddress: %s",fromAddress.Hex())
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		//log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		//log.Fatal(err)
	}
	fmt.Printf("GasPrice: %s\n", gasPrice.String())

	toAddress := common.HexToAddress(toAddr)
	tokenAddress := common.HexToAddress(tokenAddr)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	//amount := new(big.Int)
	//amount.SetString("1000000000000000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	//gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	//	To:   &toAddress,
	//	Data: data,
	//})
	//if err != nil {
	//	//log.Fatal(err)
	//}
	//fmt.Println(gasLimit) // 23256
	//gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	//	From: 		fromAddress,
	//	To: 		&tokenAddress,
	//	GasPrice: 	gasPrice,
	//	Data: 		data,
	//})

	//fmt.Printf("GasLimit: %d", gasLimit)

	tx := types.NewTransaction(nonce, tokenAddress, value, 150000, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		//log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		//log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		//log.Fatal(err)
		return "Transfer Token Fail!", err
	}

	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
	return signedTx.Hash().Hex(), nil
}

func GetTransferGasFee() (decimal.Decimal, error) {
	//var config = models.Config{}
	//tokenAddr, _ := config.GetEthUsdtContractAddress()
	//fromAddr := "0xaDc4806C2e31EB540Dd793980473A935fa960274"
	//toAddr := "0xCF52B0A8316eC79A8659945A4978898213AEC4fF"
	//amount := new(big.Int).SetInt64(1)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		//log.Fatal(err)
	}

	//fromAddress := common.HexToAddress(fromAddr)
	//toAddress := common.HexToAddress(toAddr)
	//tokenAddress := common.HexToAddress(tokenAddr)
	//
	//transferFnSignature := []byte("transfer(address,uint256)")
	//hash := sha3.NewLegacyKeccak256()
	//hash.Write(transferFnSignature)
	//methodID := hash.Sum(nil)[:4]
	//fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb
	//
	//paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	//fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d
	//
	////amount := new(big.Int)
	////amount.SetString("1000000000000000000000", 10) // 1000 tokens
	//paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	//fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000
	//
	//var data []byte
	//data = append(data, methodID...)
	//data = append(data, paddedAddress...)
	//data = append(data, paddedAmount...)

	//gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	//	To:   &toAddress,
	//	Data: data,
	//})
	//if err != nil {
	//	//log.Fatal(err)
	//}
	//fmt.Println(gasLimit) // 23256
	//gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	//	From: 		fromAddress,
	//	To: 		&tokenAddress,
	//	GasPrice: 	gasPrice,
	//	Data: 		data,
	//})

	fee := decimal.NewFromBigInt(gasPrice,0).Mul(decimal.NewFromInt(150000))

	return fee, nil
}

func BurnToken(tokenAddr string, prvKey string, amount *big.Int) (string, error) {

	privateKey, err := crypto.HexToECDSA(prvKey)
	if err != nil {
		//log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		//log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		//log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		//log.Fatal(err)
	}

	//toAddress := common.HexToAddress(toAddr)
	tokenAddress := common.HexToAddress(tokenAddr)

	transferFnSignature := []byte("burn(uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	//paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	//fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	//amount := new(big.Int)
	//amount.SetString("1000000000000000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	var data []byte
	data = append(data, methodID...)
	//data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	//gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	//	To:   &toAddress,
	//	Data: data,
	//})
	//if err != nil {
	//	//log.Fatal(err)
	//}
	//fmt.Println(gasLimit) // 23256
	gasLimit := uint64(60000) // in units

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		//log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		//log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		//log.Fatal(err)
		return "Burn Token Fail!", err
	}

	fmt.Println("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
	return signedTx.Hash().Hex(), nil
}

func GetTransactionReceipt(hash string) (bool, error) {
	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(hash))
	if err != nil {
		//log.Fatal(err)
		return false, err
	}

	fmt.Println(receipt.Status) // 1
	b, _ := json.Marshal(receipt)
	fmt.Println(string(b))   // ...
	if receipt.Status == 1 {
		return true, nil
	} else {
		return false, nil
	}
}
