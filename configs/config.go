package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	TenantID string
	UserName string
	PassWord string
	KeyName  string
}

var Config ConfigList

func init() {
	LoadConfig()
}

// 先に構造体を初期化しておく
func LoadConfig() {
	cfg, err := ini.Load("configs/config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		cfg.Section("toast").Key("tenantid").String(),
		cfg.Section("toast").Key("username").String(),
		cfg.Section("toast").Key("password").String(),
		cfg.Section("toast").Key("keyname").String(),
	}
}
