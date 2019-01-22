package routers

import (
	"searchq-operationsys/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/get", &controllers.MainController{}, "get:Join")
    beego.Router("/info", &controllers.InfoController{}, "get:Info")
    beego.Router("/getRegion", &controllers.InfoController{}, "get:GetRegion") //获取图片区域信息
    
    //review
    beego.Router("/review", &controllers.ReviewController{}, "get:Get")
    beego.Router("/review/totalinfo", &controllers.ReviewController{}, "get:GetRTotalInfo")
    beego.Router("/review/add", &controllers.ReviewController{}, "get:Add") // 添加评测
    //查询评测信息
    beego.Router("/review/query", &controllers.ReviewController{}, "get:Query") // 获取查询评测样本信息的页面
    beego.Router("/review/queryinfo", &controllers.ReviewController{}, "get:QueryInfo") // 查询评测样本信息
    beego.Router("/review/get", &controllers.ReviewController{}, "get:Getall") //获取评测样本总信息
    //评价页面
    beego.Router("/review/infoimpl", &controllers.ReviewController{}, "get:InfoImpl")
    beego.Router("/review/getRegion", &controllers.ReviewController{}, "get:GetRegion") //获取图片区域信息
    beego.Router("/review/detailinfo", &controllers.ReviewController{}, "get:GetDInfo") //获取评测填写的信息
    beego.Router("/review/result", &controllers.ReviewController{}, "*:Result")
    //翻页
    beego.Router("/review/page", &controllers.ReviewController{}, "get:Page")
    //评测总结
    beego.Router("/review/showresult", &controllers.ReviewController{}, "get:ShowResult")
     beego.Router("/review/resultinfo", &controllers.ReviewController{}, "get:ResultInfo")

    beego.Router("/review/comment", &controllers.ReviewController{}, "*:GetComment") //获取对评测结果的评论
    beego.Router("/review/getdetail", &controllers.ReviewController{}, "*:GetDetail")

    beego.Router("/review/delete", &controllers.ReviewController{}, "get:Delete") //删除评测记录


}

