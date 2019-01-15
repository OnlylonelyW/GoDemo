package controllers

import (
	"github.com/astaxie/beego"
	 "fmt"
	 "container/list"
	 "math"
	 "strconv"
	 "encoding/json"
	 "github.com/astaxie/beego/logs"
)

type ReviewController struct {
	beego.Controller
}

//评测汇总页面
func (this *ReviewController) Get() {
    this.TplName = "review.html"
}

//评测总的信息
func (this *ReviewController) GetRTotalInfo() {
	total := selectTotalR()
	t_data := selectReview(0)
	var data []ReviewEntity

	for ele := t_data.Front(); ele != nil; ele = ele.Next()  {
    	temp := ele.Value.(ReviewEntity)
    	data = append(data, temp)
    }

    total = int(math.Ceil(float64(total)/float64(20)))
	this.Data["Data"] = data 		//呈现的review数据
	this.Data["Len"] = total

	re_data := make(map[string]interface{})
	re_data["Data"] = data
	re_data["len"] = total //----总页数
	this.Data["json"] = re_data 
	this.ServeJSON()
}

//处理分页的
func (this *ReviewController) Page() {
	num := this.GetString("page")
	num1,_ := strconv.Atoi(num)
	t_data := selectReview(num1-1)
	var data []ReviewEntity
	for ele := t_data.Front(); ele != nil; ele = ele.Next()  {
    	temp := ele.Value.(ReviewEntity)
    	data = append(data, temp)
    }
    re_data := make(map[string]interface{})
	re_data["Data"] = data
    this.Data["json"] = re_data
    this.ServeJSON()
	
}

//添加评测
func (this *ReviewController) Add(){
	this.TplName="wel2.html"
}

//查询样本信息
func (this *ReviewController) Query() {
	this.TplName="wel3.html"
}

// 获取review的样本信息
func (this *ReviewController) Getall() {
	name := this.GetString("name")
	var id int
	rev := ReviewEntity{
		Name: name,
		BeginTime: this.GetString("begin"),
		EndTime: this.GetString("end"),
	}
	id = insertReview(rev)
	
	num := this.GetString("num")
	tp := this.GetString("type")
	var data []Info
	var r_list *list.List
	fmt.Println(id)
	r_list = getReview(id, num, tp)

	for ele := r_list.Front(); ele != nil; ele = ele.Next()  {
		temp := ele.Value.(Info)
		data = append(data, temp)
	}

	for _, temp := range data {
		NinsertSample(id, temp.Id)
	}
	
	Nuseful(&data, id) //判断是否评测
	//logs.Debug(data)
	d := make(map[string]interface{})
	d["Data"] = data
	d["revid"] = id
	this.Data["json"] = d
	this.ServeJSON()
	
}

// 
func (this *ReviewController) QueryInfo() {
	id := this.GetString("id")
	fmt.Println(id)
	r_list := getRDetail(id)
	var data []Info
	//var r_list *list.List
	for ele := r_list.Front(); ele != nil; ele = ele.Next()  {
		temp := ele.Value.(Info)
		data = append(data, temp)
	}
	nid, _ := strconv.Atoi(id)
	Nuseful(&data, nid) //判断是否评测

	d := make(map[string]interface{})
	d["Data"] = data
	d["revid"] = id
	this.Data["json"] = d
	this.ServeJSON()
}


var rel_list2 = make(chan []Region, 1)

//获取局部图信息
func (this *ReviewController) GetInfo() {
	name := this.GetString("id")
	rev_id := this.GetString("rev_id")
	image, json_str := selectById(name)
	logs.Debug(json_str)
	r_list := this.dealJson(json_str)
    this.Data["IsInfo2"] = true
	this.Data["url"] = image
	this.Data["rev_id"] = rev_id
	this.Data["ques_id"] = name
	rel_list2 <- r_list
	this.TplName = "info2.html"
}

// 解析log中的json
func (this *ReviewController) dealJson(json_str string) []Region{
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
    return r_list
}


func (this *ReviewController ) GetRegion() {
	this.Data["json"] = <- rel_list2
	this.ServeJSON()
}

//评测信息
type ReviewInfo struct {
	Page PageEntity
	Part []PartEntity
}

//整页信息
type PageEntity struct {
	Result string
	Rtype  string
	Grade  string
	Subject string
	All_num string
	Cut_num string
	Acc_num string
	Suc_num string
	Rev_id string
	Ques_id string
}

//分题信息
type PartEntity struct{
	Ques_id string
	Id string
	Similar string
	Cut string
	Photo string
}

//获得评测结果
func (this *ReviewController) Result() {
	//var r_list []Region
	rev := ReviewInfo{} 
    json.Unmarshal(this.Ctx.Input.RequestBody, &rev)
    logs.Debug(rev.Part==nil)
    insertRInfo(rev)
    this.Data["json"] = "hello"
    this.ServeJSON()

}

//截断为2位浮点数
func Round2(num float64) (float64) {
	num_r2 := fmt.Sprintf("%.2f", num)
	v2, _ := strconv.ParseFloat(num_r2, 64)
	return v2
}

//总结评测结果
func (this *ReviewController) ShowResult() {
	id := this.GetString("id")
	all_num := getAllSample(id)
	sampleT := SampleT(id)
	logs.Debug(sampleT)
	sampleF := SampleF(id)
	var m_num = []int{all_num, sampleT.Total}
	var n_num = []int{all_num, sampleF.Total}
	var table1 [2][4]float64
	var table2 [2][7]float64

	//将数据转化为百分比的形式
	if all_num != 0 {
		if sampleT.Total != 0 {
			for i:=0; i<2; i++ {
				table1[i][0] = Round2(float64(sampleT.Total)/float64(m_num[i]) * 100)
				table1[i][1] = Round2(float64(sampleT.Middle)/float64(m_num[i]) * 100)
				table1[i][2] = Round2(float64(sampleT.Little)/float64(m_num[i]) * 100)
				table1[i][3] = Round2(float64(sampleT.Other)/float64(m_num[i]) * 100)
			}
		}
//3 4  5 2 1 0 将数据库存储的类型和前端对应
		if(sampleF.Total!=0){
			for i:=0; i<2; i++ {
				table2[i][0] = Round2(float64(sampleF.Total)/float64(n_num[i]) * 100)
				table2[i][1] = Round2(float64(sampleF.L3)/float64(n_num[i]) * 100)
				table2[i][2] = Round2(float64(sampleF.L4)/float64(n_num[i]) * 100)
				table2[i][3] = Round2(float64(sampleF.L5)/float64(n_num[i]) * 100)
				table2[i][4] = Round2(float64(sampleF.L2)/float64(n_num[i]) * 100)
				table2[i][5] = Round2(float64(sampleF.L1)/float64(n_num[i]) * 100)
				table2[i][6] = Round2(float64(sampleF.L0)/float64(n_num[i]) * 100)
			}
		}

		
	}
	middle, m_ques := GetMParts(id)
	logs.Debug(middle)
	logs.Debug(m_ques)
 //2 1 5 4 0 3	
 	var tb3_1 [3][8]float64
 	var tb3_acc [8]float64
 	var tb3_suc [8]float64
 	var divid1 = []int{all_num, sampleT.Total, middle.Total}

 	if all_num != 0 && sampleT.Total !=0 && middle.Total != 0 {
 		for i:=0; i<3; i++ {
 			tb3_1[i][0] = Round2(float64(middle.Total)/float64(divid1[i]) * 100)
 			tb3_1[i][1] = Round2(float64(middle.NoneEng)/float64(divid1[i]) * 100)
 			tb3_1[i][2] = Round2(float64(middle.L2)/float64(divid1[i]) * 100)
 			tb3_1[i][3] = Round2(float64(middle.L1)/float64(divid1[i]) * 100)
 			tb3_1[i][4] = Round2(float64(middle.L5)/float64(divid1[i]) * 100)
 			tb3_1[i][5] = Round2(float64(middle.L4)/float64(divid1[i]) * 100)
 			tb3_1[i][6] = Round2(float64(middle.L0)/float64(divid1[i]) * 100)
 			tb3_1[i][7] = Round2(float64(middle.L3)/float64(divid1[i]) * 100)
 		}

 	}
 	var temp =[]int{2, 1, 5, 4, 0, 3} //+2
 	if m_ques.Total[0]!=0 {
 		tb3_acc[0] = Round2(float64(m_ques.Acc[0])/float64(m_ques.Total[0]) * 100)
 		tb3_acc[1] = Round2(float64(m_ques.Acc[1])/float64(m_ques.Total[1]) * 100)
 		for i:=0; i<6; i++ {
 			tb3_acc[2+i] = Round2(float64(m_ques.Acc[temp[i]+2])/float64(m_ques.Total[temp[i]+2]) * 100)
 		}

 		tb3_suc[0] = Round2(float64(m_ques.Suc[0])/float64(m_ques.Total[0]) * 100)
 		tb3_suc[1] = Round2(float64(m_ques.Suc[1])/float64(m_ques.Total[1]) * 100)
 		for i:=0; i<6; i++ {
 			tb3_suc[2+i] = Round2(float64(m_ques.Suc[temp[i]+2])/float64(m_ques.Total[temp[i]+2]) * 100)
 		}
 	}

 	this.Data["sampleT"] = sampleT
 	this.Data["sampleF"] = sampleF
 	this.Data["table2"] = table2
 	this.Data["table1"] = table1
 	this.Data["middle"] = middle
 	this.Data["m_ques"] = m_ques
 	this.Data["tb3_1"] = tb3_1
 	this.Data["tb3_suc"] = tb3_suc
 	this.Data["tb3_acc"] = tb3_acc
 	this.Data["IsResult"] = true
	this.Data["json"] = "hello"
	this.Data["id"] = id
    this.TplName = "re_result.html"
}

// 获取评语

func (this *ReviewController) GetComment() {
	id := this.GetString("id")
	comment := this.GetString("comment")
	insertComment(id, comment)
	this.Ctx.Redirect(302, "/review")
}

// 获取评测包含的题目数据
func (this *ReviewController) GetDetail() {
	id := this.GetString("id")
	r_list := getRDetail(id)
	var data []Info
    for ele := r_list.Front(); ele != nil; ele = ele.Next()  {
    	temp := ele.Value.(Info)
    	data = append(data, temp)
    }

    var d Data
    d.Data = data
    this.Data["json"] = &d
    this.ServeJSON()
}


//获取已经评测的信息
func (this *ReviewController) GetDInfo() {
	rev_id := this.GetString("rev_id")
	ques_id := this.GetString("ques_id")
	num := NRuseful(rev_id, ques_id)
	if num ==2 {
		this.Data["json"] = "false"
	}else{
		this.Data["json"] = getRInfo(rev_id, ques_id)
	}
	this.ServeJSON()
}

//删除review记录
func (this *ReviewController) Delete() {
	id := this.GetString("id")
	err := deleteReview(id)
	if err == nil {
		this.Data["json"] = "success"

	}else{
		this.Data["json"] = "fail"
	}
	this.ServeJSON()
	
}