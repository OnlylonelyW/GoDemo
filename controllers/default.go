package controllers

import (
	"github.com/astaxie/beego"
    _ "github.com/go-sql-driver/mysql"

)

type MainController struct {
	beego.Controller
}

//开始界面
func (this *MainController) Get() {
	this.Data["IsWel"] = true
    this.TplName = "wel.html"
}

//获取wel显示的数据
func (this *MainController) Join(){
    time := this.GetString("time")
    endtime := this.GetString("end")
    var data []Info
   
    json_list := selectData(time, endtime)
    for ele := json_list.Front(); ele != nil; ele = ele.Next()  {
        temp := ele.Value.(Info)
        data = append(data, temp)
    }
    
    var d Data
    d.Data = data
    this.Data["json"] = &d
    this.ServeJSON()

}