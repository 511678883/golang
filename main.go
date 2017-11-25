package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/context"
	_ "ihome/routers"
	"net/http"
	"strings"
	"ihome/models"
)

func main() {
	ignoreStaticPath()
	beego.Run()
}
func ignoreStaticPath() {

	//透明static

	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url: ", orpath)
	//如果请求uri还有api字段,说明是指令应该取消静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)
}
func init() {
	orm.RegisterDataBase("default", "mysql", "root:123@tcp(127.0.0.1:3306)/hello_Mysql?charset=utf8", 30)

	orm.RegisterModel(new(models.User), new(models.House),new(models.Area),new(models.Facility),new(models.HouseImage),new(models.OrderHouse))

	orm.RunSyncdb("default", false, true)
}