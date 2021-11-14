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

func (controllers *Controllers) GetModel(ctx iris.Context) {
	_chain := config.GetKeyInstance()
	bs := sdk.Connect(_chain.GateWay, _chain.Port, _chain.Secure, _chain.KeyPair).GetBlockchainService()

	ledgers, err := bs.GetLedgerHashs()
	if err != nil {
		fmt.Println(err)
	}

	model := make(map[string]string)
	modelName := bs.NewTransaction(ledgers[0])
	nameSender := modelName.ContractEvents()
	err = nameSender.Send(base58.MustDecode("LdeNyV1dRhp8UMgCDYxMdmWM8rXddsyRiHfsf"), 1, "getModelName")
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

	model["name"] = string(response.OperationResults[0].Result.Bytes)

	modelInfo := bs.NewTransaction(ledgers[0])
	infoSender := modelInfo.ContractEvents()
	err = infoSender.Send(base58.MustDecode("LdeNyV1dRhp8UMgCDYxMdmWM8rXddsyRiHfsf"), 1, "getModelInfo")
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

	model["info"] = string(response.OperationResults[0].Result.Bytes)
	model["addr"] = "LdeNyV1dRhp8UMgCDYxMdmWM8rXddsyRiHfsf"

	ret := make([]map[string]string, 1)
	ret[0] = model
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
	err = nameSender.Send(base58.MustDecode("LdeNgwNFVTNGgtKyPz6e2ZKBvGz7Kdri6RfeV"), 1, "getDataName")
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
	err = infoSender.Send(base58.MustDecode("LdeNgwNFVTNGgtKyPz6e2ZKBvGz7Kdri6RfeV"), 1, "getDataInfo")
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
	data["addr"] = "LdeNgwNFVTNGgtKyPz6e2ZKBvGz7Kdri6RfeV"

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

	link := string(response.OperationResults[0].Result.Bytes)
	ctx.JSON(link)
}