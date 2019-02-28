	$(document).ready(function () {
		$("button").click(function(){
			var s1 = $("#s1").val();
			var s2 = $("#s2").val();
			var temp = [];
			$("table tr").css({"display":""})
			if(s1=="android"){
				$("table tr").each(function(i, item){
					if(i>0){
						var td = $(item).find("td").eq(1);
						if($(td).text().indexOf(s1)!=-1){
							$(item).css({"display":"none"})
						}
					}
				});
			}else if(s1=="ios"){
				$("table tr").each(function(i, item){
					if(i>0){
						var td = $(item).find("td").eq(1);
						if($(td).text().indexOf(s1)!=-1){
							$(item).css({"display":"none"})
						}
					}
				});
			}
			switch(s2){
				case"multiple":
					$("table tr").each(function(i, item){
						if(i>0){
							var td = $(item).find("td").eq(3);
							if($(td).text().indexOf(s2)==-1){
								$(item).css({"display":"none"})
							}
						}
					});
					break;

				case"template":
					$("table tr").each(function(i, item){
						if(i>0){
							var td = $(item).find("td").eq(3);
							if($(td).text().indexOf(s2)==-1){
								$(item).css({"display":"none"})
							}
						}
					});
					break;
				case"single":
					$("table tr").each(function(i, item){
						if(i>0){
							var td = $(item).find("td").eq(3);
							if($(td).text().indexOf(s2)==-1){
								$(item).css({"display":"none"})
							}
						}
					});
					break;
				default:
			}




		});
		var rev_id;
		var finish = new Array(0, 0, 0);
		var unfinish = new Array(0, 0, 0);
		var total = new Array(0, 0, 0);
		var label = [".p_line1", ".p_line2", ".p_line3"]
		//后取样本数据信息name=test2&end=20190101&begin=20190101&num=20&type=
		$.ajax({
			url:"/review/queryinfo?id="+getUrlParam("id"),
			type: "GET",
			dataType:"json",
			success:function(data) { 
				if (data==null) {
					return
				}
				console.log(data.Data);
				rev_id = data.revid;
				for (var i=0 ; i<data.Data.length; i++){
					var li = document.createElement('tr');
					//序号
					var id = document.createElement('td');
					id.innerText = i+1;
					li.appendChild(id);
					//客户端
					var imei = document.createElement('td');
					imei.innerText = data.Data[i].Imei;
					li.appendChild(imei);
					//用户id
					var user = document.createElement('td');
					user.innerText = data.Data[i].User;
					li.appendChild(user);
					//日期
					var date= document.createElement('td');
					str = data.Data[i].Date.split("_");
					date.innerText = str[0];
					li.appendChild(date);
					// 命中模板与否???????????????????????
					var action = document.createElement('td');
					action.innerText = data.Data[i].Action;
					li.appendChild(action);
					//评测状态
					var state = document.createElement('td');
					var isEvaluated = str[1]
					if (isEvaluated == "1") {
	            		state.innerText = "已评"
	            		state.setAttribute("class", "reved")
	            	} else {
	            		state.innerText = "未评"
	            	}
	            	li.appendChild(state);
	            	//// 样本有效性
	            	var isSampleValid = document.createElement('td');
	            	isSampleValid.innerText = (data.Data[i].Result == "1"? "有效": "无效");
	            	li.appendChild(isSampleValid);
	            	// 学段
	            	var grade = document.createElement('td');
	            	var gradeMap = new Array("无法判断", "小学", "中学");
	            	grade.innerText = gradeMap[parseInt(data.Data[i].Grade)];
	            	li.appendChild(grade);
	            	// 学科
	            	var subject = document.createElement('td');
	            	var subjectMap = new Array("其他", "理科", "数学", "英语", "文科", "语文");
	            	subject.innerText = subjectMap[parseInt(data.Data[i].Subject)];
	            	li.appendChild(subject);
	            	// 题目总量
	            	var quesNum = document.createElement('td');
	            	quesNum.innerText = data.Data[i].All_num;
	            	li.appendChild(quesNum);
	            	// 切出数量
	            	var cutNum = document.createElement('td');
	            	cutNum.innerText = data.Data[i].Cut_num;
	            	li.appendChild(cutNum);
	            	// 搜对数量
					var succNum = document.createElement('td');
	            	succNum.innerText = data.Data[i].Suc_num;
	            	li.appendChild(succNum);
	            	// 切对率
					var cutRate = document.createElement('td');
					cutRate.innerText = (parseInt(data.Data[i].Acc_num) / parseInt(data.Data[i].All_num)).toFixed(2);
	            	li.appendChild(cutRate);
	            	// 搜对率
					var succRate = document.createElement('td');
					succRate.innerText = (parseInt(data.Data[i].Suc_num) / parseInt(data.Data[i].All_num)).toFixed(2);
	            	li.appendChild(succRate);
	            	// 详情
	            	var nowstate = document.createElement('td');
	            	var a = document.createElement("a");
	            	a.setAttribute("href", "/review/infoimpl?id="+data.Data[i].Id+"&"+"rev_id="+rev_id);
	            	a.setAttribute("target", "_blank")
	            	a.innerText = "详情";
	            	nowstate.appendChild(a)
	            	li.appendChild(nowstate);

	            	$(".wrap tr").last().after(li);
	            	
	            	//统计已评未评个数
	            	str = data.Data[i].Date.split("_")
	            	if (str[1]=="1"){
	            		//nowstate.innerText = "已评"
	            		finish[2]++;
	            		if(action.innerText.indexOf("ios")!=-1){
	            			finish[1]++;
	            		}else{
	            			finish[0]++;
	            		}
	            		nowstate.setAttribute("class","reved")
	            	}else{
	            		//nowstate.innerText = "未评"
	            		unfinish[2]++;
	            		if(action.innerText.indexOf("ios")!=-1){
	            			unfinish[1]++;
	            		}else{
	            			 unfinish[0]++;
	            		}
	            	}

            		
				}
				for (var j=0; j<3; j++){
	            	total[j] = finish[j] + unfinish[j];
	            }
	            	
	            	for (var m=0; m<3; m++){
	            		var td = $("<td>"+finish[m]+"</td>");
	            		$(".p_line1").append(td);
	            	}
	            	for (var m=0; m<3; m++){
	            		var td = $("<td>"+unfinish[m]+"</td>");
	            		$(".p_line2").append(td);
	            	}
	            	for (var m=0; m<3; m++){
	            		var td = $("<td>"+total[m]+"</td>");
	            		$(".p_line3").append(td);
	            	}
			 },
			error: function(){
				alert("error")
			}
		});
	});

	//获取url中的参数
    function getUrlParam(name) {
        var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
        var r = window.location.search.substr(1).match(reg);  //匹配目标参数
        if (r != null) return unescape(r[2]);
        return ""; //返回参数值
    }