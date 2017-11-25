package controllers

import (
	"github.com/astaxie/beego"
	 "github.com/astaxie/beego/orm"
	 _"github.com/astaxie/beego/context"
	"ihome/models"
	"fmt"
)

type AreasController struct {
	beego.Controller
}
type AreaResp struct {
	Error  string      `json:"error"`
	Errmsg string      `json : "errmsg"`
	Data   interface{} `json: "data"`
}

func (this *AreasController) RetData(AreaResp interface{}) {
	//给客户端返回json数据
	this.Data["json"] = AreaResp
	//将json写回客户端
	this.ServeJSON()
}

func (this *AreasController) GetAreas() {
	resp := AreaResp{Error: "0", Errmsg: "ok"}
	var arr_resp []models.Area

	o := orm.NewOrm()
	qs := o.QueryTable("area")
	num, err := qs.All(&arr_resp)
	if err != nil {
		resp.Error = models.RECODE_DATAERR
		resp.Errmsg = models.RecodeText(resp.Error)
		return
	}
	if num == 0 {
		//没有数据
		resp.Error = models.RECODE_NODATA
		resp.Errmsg = models.RecodeText(resp.Error)
		return
	}

	fmt.Printf("areas = %+v\n", arr_resp)
	resp.Data = arr_resp

	return

}
