package config

import (
	"io/ioutil"

	"github.com/tal-tech/go-zero/core/logx"
	"gopkg.in/yaml.v2"
)

const configPath = "./shop.yaml"

var Servers *ServerInfo

type (
	ServerInfo struct {
		ShopServer   ShopServer   `yaml:"ShopServer"`
		LoginServer  LoginServer  `yaml:"LoginServer"`
		LogOutServer LogOutServer `yaml:"LogOutServer"`
	}

	ShopServer struct {
		Host string `yaml:"Host"`
		Port string `yaml:"Port"`
	}

	LoginServer struct {
		Host string `yaml:"Host"`
		Port string `yaml:"Port"`
	}

	LogOutServer struct {
		Host string `yaml:"Host"`
		Port string `yaml:"Port"`
	}
)

func LoadMetricsConf() error {
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		logx.Errorf("Parsing metrics config file %s failed: %s", configPath, err)
		return err
	}

	Servers = new(ServerInfo)

	// 解析配置文件
	if err := yaml.Unmarshal(content, Servers); err != nil {
		logx.Errorf("Unmarshal failed: %s", err)
		return err
	}

	return nil
}
