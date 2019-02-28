package controllers

import (
	"database/sql"
)
//--------------------------------default-----------------------------------
//log信息
type Info struct{
	Id int
	User string
	Action string
	Imei string
	Date string
	Result string    // 样本是否有效
    Rtype  string    // 无效类型
    Subject string   // 学科
    Grade  string    // 学段
    All_num string   // 题目总数
    Cut_num string   // 切出数量
    Acc_num string   // 切对数量
    Suc_num string   // 搜对数量

}
//wel返回数据
type Data struct{
	Data []Info
}

//图片区域信息
type Region struct{
	Name string
	Area string
}

//-----------------------------review-----------------------------------
//评测的总体信息
type ReviewEntity struct{
	Id int 
	Name string 
	BeginTime string 
	EndTime string
	Summary sql.NullString 
	Num sql.NullString
	Type sql.NullString 
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

//有效样本数量
type SampleTrue struct {
	Total int
	Middle int
	Little int
	Other int
}
