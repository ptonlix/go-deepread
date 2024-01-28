package server

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type ServerConfig struct {
	ServerUrl   string
	AccessToken string
}

var ServerConf ServerConfig

func init() {
	// 获取ChatGpt配置
	var err error
	ServerConf.ServerUrl, err = beego.AppConfig.String("serverurl")
	if err != nil {
		logs.Error("Server配置加载失败: serverurl")
	}

	ServerConf.AccessToken, err = beego.AppConfig.String("accesstoken")
	if err != nil {
		logs.Error("Server: accesstoken")
	}

}
