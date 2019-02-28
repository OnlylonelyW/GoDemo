package controllers

import (
	"database/sql"
    "container/list"
    _ "github.com/go-sql-driver/mysql"
    "time"
    "strconv"
    //"strings"
    "fmt"
)



var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:lunatizi123@tcp(localhost:3306)/log?charset=utf8")
}

//根据id获取json ,image
func selectById(id string) (string, string){
	
	rows, err := db.Query("select image, result from log where id=?", id)
	defer rows.Close()
	checkErr(err)
	var json string
	var image string
	if rows.Next() {
		rows.Scan(&image, &json)

		
	}
	return image, json


}

//查询一段时间内的整体数据
func selectData(t string, endday string) *list.List{ 

	var logDay string
	nTime := time.Now()
	var rows *sql.Rows
	var err error
	if t =="" {
		yesTime := nTime.AddDate(0,0,-1)
		logDay = yesTime.Format("20060102")
		rows, err = db.Query("select id, user, imei, action, date from log where date="+logDay+" order by rand() limit 100")
		fmt.Println(" "+t +endday)
	}else{
		rows, err = db.Query("select id, user, imei, action, date from log where date>=? and date<=?", t, endday)
		fmt.Println(" "+t +endday)
	}

	defer rows.Close()
	checkErr(err)
	info_list := list.New()
	for rows.Next() {
		var id int
		var user string
		var imei string
		var action string
		var date string
		err := rows.Scan(&id, &user, &imei, &action, &date)
		checkErr(err)
		info := Info{
			Id:id,
			User:user,
			Imei:imei,
			Action:action,
    		Date : date,
		}
		info_list.PushBack(info)
	}
	return info_list
}

//---------------------------------------review-------------------------------------------------------------------------------------
//分页获取review评测结果
func selectReview(num int) *list.List{
	n1 := num*20
	rows, err := db.Query("select id, name, endTime, beginTime, summary, num, type from review order by id desc limit ?, 20", n1)
	checkErr(err)
	defer rows.Close()
	re_list := list.New()

	for rows.Next() {
		review := ReviewEntity{}
		err := rows.Scan(&review.Id, &review.Name, &review.EndTime, &review.BeginTime, &review.Summary, &review.Num, &review.Type)
		checkErr(err)
		re_list.PushBack(review)
		
	}
	return re_list
}

//获取review的总条数
func selectTotalR() int {
	rows, err := db.Query("select count(*) from review")
	defer rows.Close()
	checkErr(err)
	var i int

	if rows.Next() {
		rows.Scan(&i)
	}
	return i
}

// insert review
func insertReview(r ReviewEntity) (int, string) {
	var num int
	row1, _:= db.Query("select id from review where name=? and beginTime=? and endTime=? and type=?", r.Name, r.BeginTime, r.EndTime, r.Type.String)
	defer row1.Close()
	if row1.Next() {
		row1.Scan(&num)
		fmt.Println("existed")
		row1.Close()
		fmt.Println(num)
		return num, "yes"
	}
	db.Exec("insert into review(name, beginTime, endTime, type, num) values(?, ?, ?, ?, ?)", r.Name, r.BeginTime, r.EndTime, r.Type.String, r.Num.String)
	rows, _ := db.Query("select id from review where name=? and beginTime=? and endTime=? and type=?", r.Name, r.BeginTime, r.EndTime, r.Type.String)
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&num)
	}
	return num, "no"
}

// 获取review样本
func getReview(id int, num string, tp string) *list.List{
	rows1, _ := db.Query("select beginTime, endTime from review where id=?", id)
	defer rows1.Close()
	var end string
	var begin string
	if rows1.Next() {
		rows1.Scan(&begin, &end)
	}
	//times := strings.Split(time, "-")
	var num2 string
	db.QueryRow("select count(*) from reviewQuestion where idReview=?", id).Scan(&num2)
	fmt.Println("已经存在样本数量：" + num2)
	var rows *sql.Rows
	var err error
	if(num2==num){
		rows, err = db.Query("select id, user, imei, action, date from log where id in (select idQuestion from reviewQuestion where idReview=?)", id)
	}else{
		switch tp {
			case "0":
				fmt.Println(num)
				//select id, user, imei, action, date from log where date>=20181212 and date<=20181212 order by rand() limit 30;
				rows, err = db.Query("select id, user, imei, action, date from log where date>=? and  date<=? order by rand() limit ?", begin, end, num)
			case "1":
				rows, err = db.Query("select id, user, imei, action, date from log where date>="+begin+" and date<="+end+" and action like ? order by rand() limit ?","%%multiple%%", num )

			case "2":
				rows, err = db.Query("select id, user, imei, action, date from log where date>="+begin+" and date<="+end+" and action like ? order by rand() limit ?","%%" + "single" + "%%", num)
		}
	
		
	}
	//rows, err := db.Query("select id, user, imei, action, date from log where date>="+times[0]+" and date<="+times[1]+" order by rand() limit ?", num)

	checkErr(err)
	defer rows.Close()
	info_list := list.New()
	for rows.Next() {
		var id int
		var user string
		var imei string
		var action string
		var date string
		err := rows.Scan(&id, &user, &imei, &action, &date)
		checkErr(err)
		info := Info{
			Id:id,
			User:user,
			Imei:imei,
			Action:action,
    		Date : date,
		}
		info_list.PushBack(info)
	}
	return info_list
}

//insert 样本信息
func NinsertSample(id_rev int, id_ques int) {
	db.Exec("insert into reviewQuestion(idReview, idQuestion) values(?, ?)", id_rev, id_ques)
}

//判断样本是否评测
func Nuseful(data *[]Info, id_rev int){
	stmt, err := db.Prepare("select idReview from reviewQuestion where resultType=2 and idQuestion=? and idReview=?")
	for i := 0; i < len(*data); i++ {
		var name string
		err = stmt.QueryRow((*data)[i].Id, id_rev).Scan(&name)
		if err != nil {
			(*data)[i].Date = (*data)[i].Date +"_"+"1"
		}else {
			(*data)[i].Date = (*data)[i].Date +"_"+"0"
		}
	}
	stmt.Close()
}


func existPartReview(id int, squenceid int) int {
	var num int
	db.QueryRow("select count(*) from reviewPartQuestion where idQuestion=? and sequenceId=?", id, squenceid).Scan(&num)
	return num
}





//插入评测信息
func insertRInfo(rev ReviewInfo) {
	page := rev.Page
	x1, _ := strconv.Atoi(page.Rev_id)
	x2, _ :=  strconv.Atoi(page.Ques_id)
	x3, _ :=  strconv.Atoi(page.Result)
	x4, _ :=  strconv.Atoi(page.Rtype)
	x5, _ :=  strconv.Atoi(page.Grade)
	x6, _ :=  strconv.Atoi(page.Subject)
	x7, _ :=  strconv.Atoi(page.All_num)
	x8, _ :=  strconv.Atoi(page.Cut_num)
	x9, _ :=  strconv.Atoi(page.Acc_num)
	x10, _ :=  strconv.Atoi(page.Suc_num)
	db.Exec("update reviewQuestion set resultType=?, errorType=?, grade=?, subject=?, allNumber=?, cutNumber=?, cutAccurateNumber=?, searchTrueNumber=? where idReview=? and idQuestion=?", x3, x4, x5, x6, x7, x8 ,x9, x10, x1, x2)
	
	if rev.Part!=nil {
		for _, num := range rev.Part {
			y1, _ := strconv.Atoi(num.Ques_id)
			y2, _ := strconv.Atoi(num.Id)
			y3 := num.Similar
			y4, _ := strconv.Atoi(num.Cut)
			y5, _ :=  strconv.Atoi(num.Photo)
			if existPartReview(y1, y2)==0 {
				db.Exec("insert into reviewPartQuestion(idQuestion, sequenceId, similarId, cutAccurateNumber, searchTrueNumber) values(?, ?, ?, ?, ?)", y1, y2, y3, y4, y5)
			}else{
				db.Exec("update reviewPartQuestion set cutAccurateNumber=?, searchTrueNumber=? where idQuestion=? and sequenceId=?", y4, y5, y1, y2)
			}
		}
	}
}

//插入评论信息
func insertComment(id string, comment string){
	db.Exec("update review set  summary = ? where id = ?", comment, id)
	fmt.Println(comment)
}

//有效样本相关值获取
func SampleT(id string) SampleTrue {
	var sampleT SampleTrue
	db.QueryRow("select count(*) from reviewQuestion where idReview = ? and resultType=1", id).Scan(&sampleT.Total)
	db.QueryRow("select count(*) from (select * from reviewQuestion where idReview= ? and resultType=1) q where q.grade =2", id).Scan(&sampleT.Middle)
	db.QueryRow("select count(*) from (select * from reviewQuestion where idReview= ? and resultType=1 ) q where q.grade =1", id).Scan(&sampleT.Little)
	db.QueryRow("select count(*) from (select * from reviewQuestion where idReview= ? and resultType=1 ) q where q.grade =0", id).Scan(&sampleT.Other)
	return sampleT
}	

//无效样本
// 					<option value="0">其他</option>
// 				  	<option value="1">模糊</option>
// 				  	<option value="2">非K12</option>
// 				  	<option value="3">横屏拍摄</option>
// 				  	<option value="4">纯口算、计算</option>
// 				  	<option value="5">纯手写作业</option>
type SampleFalse struct{
	Total int
	L0 int
	L1 int
	L2 int
	L3 int
	L4 int
	L5 int
}
//获取无效样本信息
func SampleF(id string) SampleFalse {
	var samplef SampleFalse
	db.QueryRow("select count(*) from reviewQuestion where idReview = ? and resultType=0", id).Scan(&samplef.Total)
	db.QueryRow("select count(*) from (select * from reviewQuestion where idReview= ? and resultType=0) q where q.errorType =0", id).Scan(&samplef.L0)
	db.QueryRow("select count(*) from (select * from reviewQuestion where idReview= ? and resultType=0) q where q.errorType =1", id).Scan(&samplef.L1)
	db.QueryRow("select count(*) from (select * from reviewQuestion where idReview= ? and resultType=0) q where q.errorType =2", id).Scan(&samplef.L2)
	db.QueryRow("select count(*) from (select * from reviewQuestion where idReview= ? and resultType=0) q where q.errorType =3", id).Scan(&samplef.L3)
	db.QueryRow("select count(*) from (select * from reviewQuestion where idReview= ? and resultType=0) q where q.errorType =4", id).Scan(&samplef.L4)
	db.QueryRow("select count(*) from (select * from reviewQuestion where idReview= ? and resultType=0) q where q.errorType =5", id).Scan(&samplef.L5)
	return samplef
}

//获取样本总数
func getAllSample(id string) int {
	var num int 
	db.QueryRow("select count(*) from reviewQuestion where idReview = ? and resultType != 2", id).Scan(&num)
	return num
}

//中学样本数量
// 					<option value="0">其他</option>
// 				  	<option value="1">理科</option>
// 				  	<option value="2">数学</option>
// 				  	<option value="3">英语</option>
// 				  	<option value="4">文科</option>
// 				  	<option value="5">语文</option>
type Middle struct {
	Total int
	NoneEng int
	L0 int      			
	L1 int
	L2 int
	L3 int
	L4 int
	L5 int
}

//中学样本切题
// 1 是
// 0 否
type M_ques struct {
	Total [8]int
	Acc [8]int
	Suc [8]int
}

//获取中学样本信息和切题信息
//					<option value="0">无法判断</option>
// 				  	<option value="1">小学</option>
// 				  	<option value="2">中学</option>
func GetMParts(id string) (Middle, M_ques) {
	var middle Middle
	db.QueryRow("select count(*) from reviewQuestion where idReview = ? and grade=2 and resultType=1", id).Scan(&middle.Total)
	db.QueryRow("select count(*) from reviewQuestion where idReview = ? and grade=2 and subject=0 and resultType=1", id).Scan(&middle.L0)
	db.QueryRow("select count(*) from reviewQuestion where idReview = ? and grade=2 and subject=1 and resultType=1", id).Scan(&middle.L1)
	db.QueryRow("select count(*) from reviewQuestion where idReview = ? and grade=2 and subject=2 and resultType=1", id).Scan(&middle.L2)
	db.QueryRow("select count(*) from reviewQuestion where idReview = ? and grade=2 and subject=3 and resultType=1", id).Scan(&middle.L3)
	db.QueryRow("select count(*) from reviewQuestion where idReview = ? and grade=2 and subject=4 and resultType=1", id).Scan(&middle.L4)
	db.QueryRow("select count(*) from reviewQuestion where idReview = ? and grade=2 and subject=5 and resultType=1", id).Scan(&middle.L5)
	middle.NoneEng = middle.Total- middle.L3

	var m_ques M_ques

	var acc [8]int
	db.QueryRow("select sum(cutAccurateNumber) from reviewQuestion where idReview=? and grade=2 and resultType=1", id).Scan(&acc[0])
	for i:=1; i<7; i++ {
		db.QueryRow("select sum(cutAccurateNumber) from reviewQuestion where idReview=? and grade=2 and subject=? and resultType=1", id, i-1).Scan(&acc[i+1])
	}
	acc[1] = acc[0]- acc[5]

	var total [8]int
	db.QueryRow("select sum(allNumber) from reviewQuestion where idReview=? and grade=2 and resultType=1", id).Scan(&total[0])
	for i:=1; i<7; i++ {
		db.QueryRow("select sum(allNumber) from reviewQuestion where idReview=? and grade=2 and subject=? and resultType=1", id, i-1).Scan(&total[i+1])
	}
	total[1] = total[0]-total[5]

	var suc [8]int
	db.QueryRow("select sum(searchTrueNumber) from reviewQuestion where idReview=? and grade=2 and resultType=1", id).Scan(&suc[0])
	for i:=1; i<7; i++ {
		db.QueryRow("select sum(searchTrueNumber) from reviewQuestion where idReview=? and grade=2 and subject=? and resultType=1", id, i-1).Scan(&suc[i+1])
	}
	suc[1] = suc[0]-suc[5]
	
	m_ques = M_ques{
		Total:total,
		Acc:acc,
		Suc:suc,
	}
	return middle, m_ques

}

// 得到一次评测所包含的样本信息
func getRDetail(id_rev  string) *list.List{
	rows, err := db.Query("select id, user, imei, action, date from log where id in (select idQuestion from reviewQuestion where idReview=?)", id_rev)
	defer rows.Close()
	checkErr(err)
	info_list := list.New()
	for rows.Next() {
		var id int
		var user string
		var imei string
		var action string
		var date string
		err := rows.Scan(&id, &user, &imei, &action, &date)
		checkErr(err)

		// info := Info{
		// 	Id:id,
		// 	User:user,
		// 	Imei:imei,
		// 	Action:action,
  //   		Date : date,
		// }
		// info_list.PushBack(info)

		// get detail info
		var info Info
		var reviewInfo ReviewInfo
		num := NRuseful(id_rev, strconv.Itoa(id))
		if num == 2 {
			info = Info{
				Id:id,
				User:user,
				Imei:imei,
				Action:action,
	    		Date : date,
	    	}
		} else {
			reviewInfo = getRInfo(id_rev, strconv.Itoa(id))
			info = Info {
				Id: id,
				User:user,
				Imei:imei,
				Action:action,
	    		Date : date,
	    		Result: reviewInfo.Page.Result,
	    		Rtype: reviewInfo.Page.Rtype,
	    		Subject: reviewInfo.Page.Subject,
	    		Grade: reviewInfo.Page.Grade,
			    All_num: reviewInfo.Page.All_num,
			    Cut_num: reviewInfo.Page.Cut_num,
			    Acc_num: reviewInfo.Page.Acc_num,
			    Suc_num:reviewInfo.Page.Suc_num,
			}
		}

		info_list.PushBack(info)

	}
	return info_list
}

// 得到一次样本评测中的详细数据
func getRInfo(rev string, id_ques string) ReviewInfo {
	rows, err := db.Query("select * from reviewQuestion where idQuestion=? and idReview=?", id_ques, rev)
	defer rows.Close()
	checkErr(err)
	var page PageEntity
	var temp string
	if rows.Next() {
		err := rows.Scan(&temp, &page.Rev_id, &page.Ques_id, &page.Result, &page.Rtype, &page.Grade,&page.Subject, &page.All_num, &page.Cut_num, &page.Acc_num, &page.Suc_num)
		checkErr(err)
	}

	var part []PartEntity
	rows1, err1 := db.Query("select * from reviewPartQuestion where idQuestion=?", id_ques)
	defer rows1.Close()
	checkErr(err1)
	for rows1.Next() {
		var p PartEntity
		err2 := rows1.Scan(&temp, &p.Ques_id, &p.Id, &p.Similar, &p.Cut, &p.Photo)
		checkErr(err2)
		part = append(part, p)
	}

	review := ReviewInfo{
		Page: page,
		Part: part,
	}

	return review
}

//判断样本是否评测
func NRuseful(rev string, ques string) int {
	var num int
	db.QueryRow("select resultType from reviewQuestion where idReview=? and idQuestion=?", rev, ques).Scan(&num)
	return num
}

//删除评测记录
func deleteReview(id string) error{
	tx, err := db.Begin()
	checkErr(err)
	defer clearTransaction(tx)
	tx.Exec("delete from review where id=?", id)
	tx.Exec("delete from rev_question where id_rev=?", id)
	err1 := tx.Commit()
	return err1
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }

}

func clearTransaction(tx *sql.Tx){
    err := tx.Rollback()
    if err != sql.ErrTxDone && err != nil{
        checkErr(err)
    }
}

