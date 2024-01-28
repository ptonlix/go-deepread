package main

import (
	_ "go-deepread/routers"

	_ "go-deepread/init/wework"

	_ "go-deepread/services"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

var (
	AppName      string // 应用名称
	AppVersion   string // 应用版本
	BuildVersion string // 编译版本
	BuildTime    string // 编译时间
	GitRevision  string // Git版本
	GitBranch    string // Git分支
	GoVersion    string // Golang信息
)

func init() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"./log/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
}

func main() {

	Version()
	beego.Run()
}

// Version 版本信息
func Version() {
	logs.Info("App Name:\t\t%s", AppName)
	logs.Info("App Version:\t%s", AppVersion)
	logs.Info("Build version:\t%s", BuildVersion)
	logs.Info("Build time:\t%s", BuildTime)
	logs.Info("Git revision:\t%s", GitRevision)
	logs.Info("Git branch:\t%s", GitBranch)
	logs.Info("Golang Version:\t%s", GoVersion)
}
