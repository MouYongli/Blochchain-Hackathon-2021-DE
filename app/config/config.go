package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

type Config struct {
	Server	*Server	`json:"Server,object"`
	Chain	*Chain	`json:"Chain,object"`
}

func (config *Config) ReadConfig() {
	bytes, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		panic(err)
		os.Exit(-1)
	}
	err = json.Unmarshal(bytes, config)
	if err != nil {
		panic(err)
		os.Exit(-1)
	}
}

var (
	onceConfig	sync.Once
	_config		*Config
)

func GetConfigInstance() *Config {
	onceConfig.Do(func() {
		_config = new(Config)
		_config.ReadConfig()
	})
	return _config
}