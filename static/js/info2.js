	var realwidth, realheight;
	var list;
	var color=["red","green","yellow","orange","blue", "pink", "MAROON", "TEAL", "PURPLE"];
	var width, height;
	var n_w, n_h;
	$(function(){ 
		loadRegion();



		$('#s1').change(function(){
			if($(this).val()==0) {
				$(".d_area").each(function(i, item){
					$(item).css({"display": "None"})
				});
			}else {
				$(".d_area").each(function(i, item){
					$(item).css({"display": ""})
				});
			}
		});
		
		$(".show").click(function(){
			$(".table").slideToggle();
		});

		$(".submit").click(function(){
			var rev_id = getUrlParam("rev_id");
			var ques_id = getUrlParam("id");

			var result = $("#s1").val();
			var rtype = $("#s2").val();
			var grade = $("#s3").val();

			var subject = $("#s4").val();
			var all_num = $("#s5").val();
			var cut_num = $("#s6").val();
			var acc_num = $("#s7").val();
			var suc_num = $("#s8").val();
			var ob = new Object();
			ob.result = result;
			ob.rtype = rtype;
			ob.grade = grade;
			ob.subject = subject;
			ob.all_num = all_num;
			ob.cut_num = cut_num;
			ob.acc_num = acc_num;
			ob.suc_num = suc_num;
			ob.rev_id = rev_id;
			ob.ques_id = ques_id;

			var part = new Array();
			if ($(".d_area") != null) {
				$(".d_area").each(function(i, item){
					var o = new Object();

					var id = $(item).find("span[name='id']").text();
					var similar = $(item).find("span[name='similar']").text();
					var cut = $(item).find("select[name='cut']").val();
					var photo = $(item).find("select[name='photo']").val();
					o.id = id;
					o.similar = similar;
					o.cut = cut;
					o.photo = photo;
					o.ques_id = ques_id;
					part.push(o)
				});
			}
			var result_ = new Object();
			result_.page = ob;
			if((part.length>0) && ($("#s1").val()!=0)){
				result_.part = part;
			}
			console.log(JSON.stringify(result_));
			$.ajax({
				url:"/review/result",
				type: "post",
				contentType:"application/json; charset=utf-8",
				data : JSON.stringify(result_),
				dataType:"json",
				success:function(data) { 
					$("#suc").text("success")
				},
				error : function() {
		            alert("异常"); 
		        }
			});
			



		})
		
	});

	function loadRegion(){
		$.ajax({
			url:"/review/getRegion?id="+getUrlParam("id")+"&rev_id=" + getUrlParam("rev_id"),
			type: "GET",
			dataType:"json",
			success:function(data) { 
				if(data==null){
					return;
				}
				console.log(data)
				var imageUrl = data["image"];
				var dat = data["list"];
				var img = $('<img src="'+imageUrl+'" width="100%" height="100%">')
				$(".inner").append(img)
				if(dat==null){
					return;
				}

				$("img").load(function() {
					width = $("img").width();
					height = $('img').height();
					if(dat[0].Area=="template"){
						deleteItem();
						add(dat[0].Name);
						return;
					}
					if(dat[0].Area=="single"){
						deleteItem();
						for (var i=0; i<dat.length; i++){
							add(dat[i].Name)
						}
						return;
					}
					getImageWidth(imageUrl, dat); 
				});
				console.log(data)
				
				filling()
			},
			error: function(){
				alert("error")
			}
		});
	}

	//填写已经评测的数据
	function filling(){
		var rev_id = getUrlParam("rev_id");
		var ques_id = getUrlParam("id");
		$.ajax({
			url:"/review/detailinfo?rev_id="+rev_id+"&ques_id="+ques_id,
			type:"GET",
			dataType:"json",
			success:function(data){
				if (data=="false") {
					//alert("false")
					return
				}else{
					console.log(data)
					$("#s1").val(data.Page.Result)
					$("#s2").val(data.Page.Rtype)
					$("#s3").val(data.Page.Grade)
					$("#s4").val(data.Page.Subject)
					$("#s5").val(data.Page.All_num)
					$("#s6").val(data.Page.Cut_num)
					$("#s7").val(data.Page.Acc_num)
					$("#s8").val(data.Page.Suc_num)
					$(".d_area").each(function(i, item){
						$(item).find("select[name='cut']").val(data.Part[i].Cut);
						$(item).find("select[name='photo']").val(data.Part[i].Photo);
					});
				}
			},
			error: function() {
				alert("fail")
			}
		});
	}

	//添加iframe
	function add(url){
		var iframe = $('<iframe src="http://testsearchq.youdao.com/question/'+url+'"><p>您的浏览器不支持  iframe 标签。</p></iframe>')
		$(".right").append(iframe)
	}

	//删除所有子元素
	function deleteItem(){
		$(".right").empty()
	}

	function getImageWidth(url, data){
		var img = new Image();
		
	    img.onload = function(){
	    	realheight = img.height;
	    	realwidth = img.width;
		    n_w = realwidth/width;
		    n_h = realheight/height;

		    for (var i=0; i<data.length; i++){
		    	var temp = data[i].Area.split(",");
		    	var x = Math.round(temp[0]/n_w);
		    	var y = Math.round(temp[1]/n_h);
		    	var w = Math.round(temp[2]/n_w);
		    	var h = Math.round(temp[3]/n_h);
		    	var div = $("<div class='area'>"+(i+1)+"</div>"); 
		    	var t_color = color[Math.floor(Math.random() * 7 )]
		    	div.attr("style", "width:"+w+"px; height:"+h+"px; top:"+y+"px; left:"+x+"px; opacity:0.4; background-color:"+t_color);
		    	div.attr('name', data[i].Name)
		    	$('.inner').append(div);
		    	var t_div = $('<div class="d_area">id:<span class="t1" name="id">'+(i+1)+'</span>similar:<span class="t1" name="similar">'+data[i].Name+'</span><span class="t1">切准：<select name="cut"><option value=""></option><option value="1">是</option><option value="0">否</option></select></span><span class="t1">拍搜正确:<select name="photo"><option value=""></option><option value="1">是</option><option value="0">否</option></select></span></div>');
		    	$(".table").append(t_div)
		    }

		    $('.area').click(function(){
				var name= $(this).attr("name");
	    		deleteItem();
	    		add(name);
	    		//$('body,html').animate({scrollTop:0},200);
	    	});
	    	
	    }
	    img.src = url;

	}

	//获取url中的参数
    function getUrlParam(name) {
        var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
        var r = window.location.search.substr(1).match(reg);  //匹配目标参数
        if (r != null) return unescape(r[2]);
        return ""; //返回参数值
    }