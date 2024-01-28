package wework

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type WeworkConfig struct {
	AgentId    int
	CorpId     string
	CorpSecret string
	PKey       string
	PToken     string
}

var WeworkConf WeworkConfig

func init() {
	// 获取ChatGpt配置
	var err error
	WeworkConf.AgentId, err = beego.AppConfig.Int("agentid")
	if err != nil {
		logs.Error("Wework配置加载失败: corpid")
	}
	WeworkConf.CorpId, err = beego.AppConfig.String("corpid")
	if err != nil {
		logs.Error("Wework配置加载失败: corpid")
	}
	WeworkConf.CorpSecret, err = beego.AppConfig.String("corpsecret")
	if err != nil {
		logs.Error("Wework配置加载失败: corpsecret")
	}
	WeworkConf.PKey, err = beego.AppConfig.String("pkey")
	if err != nil {
		logs.Error("Wework配置加载失败: pkey")
	}
	WeworkConf.PToken, err = beego.AppConfig.String("ptoken")
	if err != nil {
		logs.Error("Wework配置加载失败: ptoken")
	}

}
