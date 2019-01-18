	var realwidth, realheight;
	var list;
	var color=["red","green","yellow","orange","black","blue", "pink", "MAROON", "TEAL", "PURPLE"];
	var width, height;
	var n_w, n_h;

	$(function(){ 

		var id = getUrlParam("id")
		urlLoad(id)

	});

	function urlLoad(id){
		$.ajax({
				url:"/getRegion?id="+id,
				type: "GET",
				dataType:"json",
				success:function(data) { 
					if(data==null){
						return;
					}

					imageUrl = data["image"];
					var dat = data["list"]
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
					
					
					
				 	
				},
				error: function(){
					alert("error")
				}
		});
	}

	//获取url中的参数
    function getUrlParam(name) {
        var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
        var r = window.location.search.substr(1).match(reg);  //匹配目标参数
        if (r != null) return unescape(r[2]);
        return ""; //返回参数值
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
		    	var div = $("<div class='area'></div>"); 
		    	var t_color = color[Math.floor(Math.random() * 7 )]
		    	div.attr("style", "width:"+w+"px; height:"+h+"px; top:"+y+"px; left:"+x+"px; opacity:0.4; background-color:"+t_color);
		    	div.attr('name', data[i].Name)
		    	$('.inner').append(div);
		    }

		    $('.area').click(function(){
				var name= $(this).attr("name");
	    		deleteItem();
	    		add(name);
	    		$('body,html').animate({scrollTop:0},100);
	    	});
	    	
	    }
	    img.src = url;

	}