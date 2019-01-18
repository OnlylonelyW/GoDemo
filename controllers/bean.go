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