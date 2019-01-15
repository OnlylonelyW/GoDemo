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
}
