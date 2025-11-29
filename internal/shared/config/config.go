package config

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

type DefaultConfig struct {
	Apps struct {
		Name     string `json:"name"`
		HttpPort string `json:"httpPort"`
	}
}

func New(path string) *DefaultConfig {
	fmt.Printf("new config...")

	viper.SetConfigFile(path)
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	defaultConfig := DefaultConfig{}
	err := viper.Unmarshal(&defaultConfig)
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(defaultConfig)
	fmt.Println(b)

	return &defaultConfig
}

func (c *DefaultConfig) AppAddress() string {
	return fmt.Sprintf(":%v", c.Apps.HttpPort)
}
