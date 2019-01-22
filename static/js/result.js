$(function(){ 
	var id = $("#id").text()
	$("#submit").click(function(){
		var text = $("#editarea").text();
		var form = $("<form></form>");
		form.attr('action', "/review/comment?id="+id);
		form.attr('method', 'post');
		form.attr('target', '_self');
		var input1 = $("<input type='hidden' name='"+"comment"+"' />");
		input1.attr('value', text);
		form.append(input1);
		
		form.appendTo("body");
		form.css('display', 'none');		
		form.submit();
	});

	loadData();
});

function loadData(){
	$.ajax({
		url:"/review/resultinfo?id="+getUrlParam("id"),
		type: "GET",
		dataType:"json",
		success:function(data) { 
			if (data==null) {
				return;
			}
			console.log(data);
			//table1
			filltabel1(data);
			//tabel2
			filltabel2(data);
			//tabel3
			filltabel3(data);

			//

		},
		error:function(){
			alert("error")
		}
	});


}

function filltabel1(data) {
	var td;
	var lb = ["Total", "Middle", "Little", "Other"];
	for (var t1 in lb) {
		td = $("<td></td>");
		$(".t1_1").append(td.text(data.sampleT[lb[t1]]));
	}
	
	var label1 = [".t1_2", ".t1_3"];
	for (var i=0; i < data.table1.length; i++) {
		for (item in data.table1[i]) {
			td = $("<td></td>");
			$(label1[i]).append(td.text(data.table1[i][item]));
		}
	}
}

function filltabel2(data) {
	var td;
	var lb = ["Total", "L3", "L4", "L5", "L2", "L1", "L0"];
	for (item in lb) {
		td = $("<td></td>");
		//console.log(item);
		$(".t2_1").append(td.text(data.sampleF[lb[item]]));
	}
	var label1 = [".t2_2", ".t2_3"];
	for (var i=0; i < data.table2.length; i++) {
		for (item in data.table2[i]) {
			td = $("<td></td>");
			$(label1[i]).append(td.text(data.table2[i][item]));
		}
	}
}

function filltabel3(data) {
	var td;
	var label1 = ["Total", "NoneEng", "L2", "L1", "L5", "L4", "L0", "L3"];
	var label2 = [0, 1, 4, 3, 7, 6, 2, 5];
	var label3 = ["Total", "Acc"];
	for (item in label1) {
		td = $("<td></td>");
		$(".t3_1").append(td.text(data.middle[label1[item]]));
	}

	var label1 = ["t2_2", "t2_3"];
	for (var i=0; i < data.tb3_1.length; i++) {
		for (item in data.tb3_1[i]) {
			td = $("<td></td>")
			$(".t3_"+(2+i)).append(td.text(data.tb3_1[i][item]));
		}
	}

	for (var i = 0; i<2; i++) {
		for (var j=0; j<label2.length; j++) {
			//console.log(data.m_ques[label3[i]][0]); 
			td = $("<td></td>")
			$(".t3_"+(5+i)).append(td.text(data.m_ques[label3[i]][parseInt(label2[j])]));
		}
	}

	for (item in data.tb3_acc) {
		td = $("<td></td>")
		$(".t3_7").append(td.text(data.tb3_acc[item]));
	}

	for (var i=0; i<label2.length; i++) {
		td = $("<td></td>")
		//console.log(label2[i]);
		$(".t3_8").append(td.text(data.m_ques.Suc[label2[i]]));
	}

	for (item in data.tb3_suc) {
		td = $("<td></td>")
		$(".t3_9").append(td.text(data.tb3_suc[item]));
	}
}
//获取url中的参数
    function getUrlParam(name) {
        var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
        var r = window.location.search.substr(1).match(reg);  //匹配目标参数
        if (r != null) return unescape(r[2]);
        return ""; //返回参数值
    }