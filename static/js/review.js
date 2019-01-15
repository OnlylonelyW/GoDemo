	var len
	$(function () {
		$("button").click(function(){
			var name = $("input[name='name']").val();
			var end = $("input[name='end']").val() ;
			var begin = $("input[name='begin']").val() ;
			var num =$("input[name='num']").val() ;
			var type = $("#s2").val()
			window.location.href="/review/add?name="+name+"&end="+end+"&begin="+begin +"&num="+num+"&type="+type;
		});
		
		loadinfo()
		$("#next").click(function(){
			var cur = parseInt($("#current").text());
			var total = parseInt($("#total").text());
			var next;
			if (total>cur){;
				loadPage(cur+1)
				next = cur + 1;
			}else {
				loadPage(1);
				next = 1;
			}
			$("#current").text(next)

		});
		//下一页
		$("#pre").click(function(){
			var cur = parseInt($("#current").text());
			var total = parseInt($("#total").text());
			var next = cur;
			if (cur>1){;
				loadPage(cur-1)
				next = cur - 1;
			}else {
				loadPage(total);
				next = total;
			}
			$("#current").text(next)

		})

		$(".delete").click(function(){
			deleteInfo(this);
			alert("deleteInfo")
		})
	});

	//翻页
	function loadPage(page){
		$.ajax({
			url: "/review/page?page="+page,
			type: "GET",
			dataType: "json",
			success: function(data){
				$(".info").remove();

				if(data==null){
					return;
				}
				addData(data);
				console.log(data)
			},
			error: function(){
				alert("loadPage error")
			}
		});
	}

	//加载评测数据
	function loadinfo(){
		$.ajax({
			url: "/review/totalinfo",
			type: "GET",
			dataType: "json",
			success: function(data){ 
				if(data==null){
					return;
				}
				console.log(data)
				addData(data)
				$("#total").text(data.len)
				$("#current").text(1)
			},
			error: function() {
				alert("error");
			}
		});
	}

	//添加评测数据到网页
	function addData(data){
		for (var i=0; i<data.Data.length; i++){
			var tr = $("<tr class='info'></tr>")
			var t1 = $("<td></td>")
			var t2 = $("<td></td>")
			var t3 = $("<td></td>")
			var t4 = $("<td></td>")
			var t5 = $("<td></td>")
			var t6 = $("<td></td>")
			var t7 = $("<td><button class='delete' value="+data.Data[i].Id+">delete</button></td>")
			t1.text(data.Data[i].Id)
			var re_info = $('<a href="/review/query?id='+data.Data[i].Id+'"></a>"')
			re_info.text(data.Data[i].Name);
			t2.append(re_info);
			t3.text(data.Data[i].BeginTime);
			t4.text(data.Data[i].EndTime);
			t5.text(data.Data[i].Summary);

			t6.text("retult");

			tr.append(t1);
			tr.append(t2);
			tr.append(t3);
			tr.append(t4);
			tr.append(t5);
			tr.append(t6);
			tr.append(t7);
			$(".tl").append(tr);
		}
		
	}


	function deleteInfo(label){
		$.ajax({
			url: "/review/delete?id="+$(label).val(),
			type: "GET",
			dataType: "json",
			success: function(data){
				$(label).parent().parent().remove();
				console.log(data)
			},
			error: function(){
				alert("error")
			}
		});
	}