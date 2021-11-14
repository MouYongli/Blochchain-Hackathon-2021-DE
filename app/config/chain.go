package config

import (
	"github.com/blockchain-jd-com/framework-go/crypto"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/ledger_model"
	"github.com/blockchain-jd-com/framework-go/utils/base58"
	"sync"
)

type Chain struct {
	GateWay	string	`json:"gateway"`
	Port 	int 	`json:"port"`
	Secure	bool	`json:"secure"`
	Priv 	string 	`json:"priv"`
	Pub 	string 	`json:"pub"`
	Pwd 	string 	`json:"pwd"`
}

type AccessKey struct {
	GateWay 	string
	Port 		int
	Secure 		bool
	PrivateKey	framework.PrivKey
	PublicKey 	framework.PubKey
	KeyPair 	ledger_model.BlockchainKeypair
}

var (
	onceChain	sync.Once
	accessKey	*AccessKey
)

func (accessKey *AccessKey) GenerateKey() {
	_chain := GetConfigInstance().Chain
	accessKey.GateWay = _chain.GateWay
	accessKey.Port = _chain.Port
	accessKey.Secure = _chain.Secure
	accessKey.PrivateKey = crypto.DecodePrivKey(_chain.Priv, base58.MustDecode(_chain.Pwd))
	accessKey.PublicKey = crypto.DecodePubKey(_chain.Pub)
	accessKey.KeyPair = ledger_model.NewBlockchainKeypair(accessKey.PublicKey, accessKey.PrivateKey)
}

func GetKeyInstance() *AccessKey {
	onceChain.Do(func() {
		accessKey = new(AccessKey)
		accessKey.GenerateKey()
	})
	return accessKey
}