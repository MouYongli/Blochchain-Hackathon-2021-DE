package controllers

import (
	"app/config"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/sdk"
	"github.com/blockchain-jd-com/framework-go/utils/base58"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"net/http"
)

type Controllers struct {}

var contractAddress = [...]string {
	"LdeNvYsTHZuoLYadGcrKqn96fy8zfo1YQKEJf",
	"LdeNwqCNuThz5wCgFiyHR8xt5azfvJzc6rxgA",
	"LdeNnmcCrBDy8djM7j7CPX57sNpbeiTQJdCqZ",
	"LdeNsHBsCXSt2fyqnbZaS4oNr2mwqYn4iPRUo",
	"LdeNmtM1gLwkvfTNK55XcYKkT7hUp38wMW2Tv",
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
	err = nameSender.Send(base58.MustDecode("LdeNuQ8etNT1gjAynzp7cbG4w6QSgpvbXDb3u"), 1, "getDataName")
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
	err = infoSender.Send(base58.MustDecode("LdeNuQ8etNT1gjAynzp7cbG4w6QSgpvbXDb3u"), 1, "getDataInfo")
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
	data["addr"] = "LdeNuQ8etNT1gjAynzp7cbG4w6QSgpvbXDb3u"

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
	ret["token"] = string(response.OperationResults[0].Result.Bytes)
	ctx.JSON(ret)
}

func (controllers *Controllers) UploadImg(ctx iris.Context) {
	file, info, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
		return
	}
	defer file.Close()
	// write image file to folder
	path := config.GetConfigInstance().Server.Uploads + info.Filename
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
		return
	}
	err = ioutil.WriteFile(path, data, 0666)
	if err != nil {
		fmt.Println(err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
		return
	}

	val := make(map[string]string)
	val["img_url"] = info.Filename
	bytesData, _ := json.Marshal(val)
	resp, _ := http.Post(config.GetConfigInstance().Server.Model,"application/json", bytes.NewReader(bytesData))
	body, _ := ioutil.ReadAll(resp.Body)
	ret := make(map[string]interface{})
	err = json.Unmarshal(body, &ret)
	result := make(map[string]interface{})
	result["status"] = ret["status"]
	result["url"] = config.GetConfigInstance().Server.Downloads + ret["url"].(string)
	ctx.JSON(result)
}
