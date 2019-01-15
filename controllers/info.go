package controllers

import (
	"github.com/astaxie/beego"
	 "encoding/json"
)

type InfoController struct {
	beego.Controller
}

//var rel_list = make(chan []Region, 1)

func (this *InfoController) Info() {
	
	// this.Data["url"] = image
	// rel_list <- r_list
	this.TplName = "info.html"
}


func (this *InfoController ) GetRegion() {
    name := this.GetString("id")
    image, json_str := selectById(name)

    var dat map[string]interface{}
    var r_list []Region
    json.Unmarshal([]byte(json_str), &dat)

    if v, ok := dat["questions"]; ok {
        question := v.([]interface{})
        for _, item := range question{
            w_item := item.(map[string]interface{})
            region := Region{
                Name:w_item["similarId"].(string),
                Area:w_item["region"].(string),
            }
            r_list = append(r_list, region)
        }
    }else if v, ok := dat["similarIds"]; ok {
        areas := v.([]interface{})
        for _, i := range areas {
            region := Region{
                Name:i.(string),
                Area:"single",
            }
            r_list = append(r_list, region)
        }
        
    } else if c, ok :=dat["templateId"]; ok{
        if c != "" {
            region := Region{
                Name:c.(string),
                Area:"template",
            }
            r_list = append(r_list, region)
        }
        
    }

    result := make(map[string]interface{})
    result["image"] = image
    result["list"] = r_list

	this.Data["json"] = result
	this.ServeJSON()
}