package routers

import (
	"go-deepread/controllers"
	"go-deepread/services"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	drc := &controllers.DeepReadController{DrServer: services.NewDeepReadService()}
	hl, _ := drc.NewHttpHandler()

	beego.Router("/", &controllers.MainController{})
	beego.Handler("/deepread", hl)
}
