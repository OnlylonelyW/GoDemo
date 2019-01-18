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

		//后取样本数据信息name=test2&end=20190101&begin=20190101&num=20&type=
		$.ajax({
			url:"/review/get?name="+getUrlParam("name")+"&num="+ getUrlParam("num") +"&type=" + getUrlParam("type") + "&end=" + getUrlParam("end") + "&begin=" + getUrlParam("begin"),
			type: "GET",
			dataType:"json",
			success:function(data) { 
				console.log(data.Data);
				if (data==null) {
					return
				}
				rev_id = data.revid;
				for (var i=0 ; i<data.Data.length; i++){
					var li = document.createElement('tr');
					var id = document.createElement('td');
	            	var user = document.createElement('td');
	            	var action = document.createElement('td');;
	            	var date= document.createElement('td');
	            	var state = document.createElement('td');
	            	var imei = document.createElement('td');
	            	var a = document.createElement("a")
	            	var nowstate = document.createElement('td');

	            	a.setAttribute("href", "/review/infoimpl?id="+data.Data[i].Id+"&"+"rev_id="+rev_id);
	            	a.setAttribute("target", "_blank")
	            	str = data.Data[i].Date.split("_")
	            	//console.log(str)

	            	id.innerText = i+1;
	            	user.innerText = data.Data[i].User;
	            	action.innerText = data.Data[i].Action;
	            	if (str[1]=="1"){
	            		nowstate.innerText = "已评"
	            		nowstate.setAttribute("class","reved")
	            	}else{
	            		nowstate.innerText = "未评"
	            	}
	            	imei.innerText = data.Data[i].Imei;
	            	date.innerText = str[0];
	
	            	a.innerText = "详情";
	            	state.appendChild(a)
            		
            		li.appendChild(id);
            		li.appendChild(user);
            		li.appendChild(imei);
            		li.appendChild(action);
            		li.appendChild(date);
            		li.appendChild(nowstate);
            		li.appendChild(state);

            		$(".wrap tr").last().after(li);
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