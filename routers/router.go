package routers

import (
	"github.com/astaxie/beego"
	"ihome/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1.0/areas", &controllers.AreasController{}, "get:GetAreas")
	beego.Router("/api/v1.0/session", &controllers.SessionController{}, "get:GetSession")
}
