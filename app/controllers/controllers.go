package controllers

import (
	"app/config"
	"fmt"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/sdk"
	"github.com/blockchain-jd-com/framework-go/utils/base58"
	"github.com/kataras/iris/v12"
)

type Controllers struct {}

var contractAddress = [...]string {
	"LdeNvMxTEmYbRtWeiPB8nw4TFccdqyqsAsiAE",
	"LdeNzbwk9Hn5aJUDVCAZew7o5Cg5Q9xzKhtE5",
	"LdeNtetGdHSyRgw4wDWX4Q8Mpe9kRZKznb65v",
	"LdeNfcDNKHNHKbwBfM5h2yx9n8STcF4bHG8uB",
	"LdeNxK5anSpxgAfpXbDCfM6Mx1EXgPs6sfKnp",
}

func (controllers *Controllers) GetModel(ctx iris.Context) {
	_chain := config.GetKeyInstance()
	bs := sdk.Connect(_chain.GateWay, _chain.Port, _chain.Secure, _chain.KeyPair).GetBlockchainService()

	ledgers, err := bs.GetLedgerHashs()
	if err != nil {
		fmt.Println(err)
	}

	ret := make([]map[string]string, len(contractAddress))
	for index, addr := range contractAddress {
		model := make(map[string]string)
		model["addr"] = addr

		txTemp := bs.NewTransaction(ledgers[0])
		sender := txTemp.ContractEvents()
		err = sender.Send(base58.MustDecode(addr), 1, "getModelName")
		if err != nil {
			fmt.Println(err)
		}

		ptx := txTemp.Prepare()
		ptx.Sign(framework.AsymmetricKeypair{
			PubKey:  _chain.PublicKey,
			PrivKey: _chain.PrivateKey,
		})

		resp, err := ptx.Commit()
		if err != nil {
			fmt.Println(err)
		}

		model["name"] = string(resp.OperationResults[0].Result.Bytes)

		txTemp = bs.NewTransaction(ledgers[0])
		sender = txTemp.ContractEvents()
		err = sender.Send(base58.MustDecode(addr), 1, "getModelInfo")
		if err != nil {
			fmt.Println(err)
		}

		ptx = txTemp.Prepare()
		ptx.Sign(framework.AsymmetricKeypair{
			PubKey:  _chain.PublicKey,
			PrivKey: _chain.PrivateKey,
		})

		resp, err = ptx.Commit()
		if err != nil {
			fmt.Println(err)
		}

		model["info"] = string(resp.OperationResults[0].Result.Bytes)
		ret[index] = model
	}

	ctx.JSON(ret)
}

func (controllers *Controllers) GetData(ctx iris.Context) {
	_chain := config.GetKeyInstance()
	bs := sdk.Connect(_chain.GateWay, _chain.Port, _chain.Secure, _chain.KeyPair).GetBlockchainService()

	ledgers, err := bs.GetLedgerHashs()
	if err != nil {
		fmt.Println(err)
	}

	data := make(map[string]string)
	modelName := bs.NewTransaction(ledgers[0])
	nameSender := modelName.ContractEvents()
	err = nameSender.Send(base58.MustDecode("LdeNuhTDvAAynX9WVob7rci7ratnpcfZ5VV6D"), 1, "getDataName")
	if err != nil {
		fmt.Println(err)
	}

	ptx := modelName.Prepare()
	ptx.Sign(framework.AsymmetricKeypair{
		PubKey:  _chain.PublicKey,
		PrivKey: _chain.PrivateKey,
	})
	response, err := ptx.Commit()
	if err != nil {
		fmt.Println(err)
	}

	data["name"] = string(response.OperationResults[0].Result.Bytes)

	modelInfo := bs.NewTransaction(ledgers[0])
	infoSender := modelInfo.ContractEvents()
	err = infoSender.Send(base58.MustDecode("LdeNuhTDvAAynX9WVob7rci7ratnpcfZ5VV6D"), 1, "getDataInfo")
	if err != nil {
		fmt.Println(err)
	}

	infoPtx := modelInfo.Prepare()
	infoPtx.Sign(framework.AsymmetricKeypair{
		PubKey:  _chain.PublicKey,
		PrivKey: _chain.PrivateKey,
	})
	response, err = infoPtx.Commit()
	if err != nil {
		fmt.Println(err)
	}

	data["info"] = string(response.OperationResults[0].Result.Bytes)
	data["addr"] = "LdeNuhTDvAAynX9WVob7rci7ratnpcfZ5VV6D"

	ret := make([]map[string]string, 1)
	ret[0] = data
	ctx.JSON(ret)
}

func (controllers *Controllers) Trade(ctx iris.Context) {
	contractId := ctx.Params().Get("contractId")
	_chain := config.GetKeyInstance()
	bs := sdk.Connect(_chain.GateWay, _chain.Port, _chain.Secure, _chain.KeyPair).GetBlockchainService()

	ledgers, err := bs.GetLedgerHashs()
	if err != nil {
		fmt.Println(err)
	}

	modelName := bs.NewTransaction(ledgers[0])
	nameSender := modelName.ContractEvents()
	err = nameSender.Send(base58.MustDecode(contractId), 1, "trade")
	if err != nil {
		fmt.Println(err)
	}

	ptx := modelName.Prepare()
	ptx.Sign(framework.AsymmetricKeypair{
		PubKey:  _chain.PublicKey,
		PrivKey: _chain.PrivateKey,
	})
	response, err := ptx.Commit()
	if err != nil {
		fmt.Println(err)
	}

	ret := make(map[string]string)
	ret["url"] = string(response.OperationResults[0].Result.Bytes)
	ctx.JSON(ret)
}
