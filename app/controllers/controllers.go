package controllers

import (
	"app/config"
	"github.com/blockchain-jd-com/framework-go/sdk"
	"github.com/kataras/iris/v12"
)

type Controllers struct {}

//func (controllers *Controllers) Index(ctx iris.Context) {
//	ctx.View("index.html")
//}

func (controllers *Controllers) Query() {
	_chain := config.GetKeyInstance()
	_ = sdk.Connect(_chain.GateWay, _chain.Port, _chain.Secure, _chain.KeyPair).GetBlockchainService()
}

func (controllers *Controllers) GetContract(ctx iris.Context) {

}

func (controllers *Controllers) PostContract(ctx iris.Context) {

}