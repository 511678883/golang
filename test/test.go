package mian

import (
	"github.com/astaxie/beego"
)

type SessionController struct {
	beego.Controller
}

type SessionResp struct {
	Errno  string      `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

type Name struct {
	Name string `json:"name"`
}

func main() {
	f1()
}
func (this *SessionController) f1() {
	beego.Debug("get /api/v1.0/session....")

	//name := Name{Name: "zhang3"}

	resp := SessionResp{Errno: "1", Errmsg: "error"}

	this.Data["json"] = resp
	this.ServeJSON()
}
