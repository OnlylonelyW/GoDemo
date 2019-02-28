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
                      //  console.log(i + ", " + $(td).text())
						if($(td).text().indexOf(s1)==-1){

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
		
		add("", "")

		$(".subtime").click(function(){
			var time = $("input[name='time']").val();
			var end = $("input[name='end']").val();
			add(time, end)
		});

		function add(data, end){
			$.ajax({
				url:"/get?time="+data+"&"+"end="+end,
				type: "GET",
				dataType:"json",
				success:function(data) { 
					console.log(data.Data);
					if(data.Data==null){
						return
					}
					$(".wrap tr:gt(0)").remove();
					for (var i=0 ; i<data.Data.length; i++){
						var li = document.createElement('tr');
						var id = document.createElement('td');
		            	var user = document.createElement('td');
		            	var action = document.createElement('td');;
		            	var date= document.createElement('td');
		            	var state = document.createElement('td');
		            	var imei = document.createElement('td');
		            	var a = document.createElement("a")

		            	a.setAttribute("href", "/info?id="+data.Data[i].Id);
		            	a.setAttribute("target", "_blank");

		            	id.innerText = i+1;
		            	user.innerText = data.Data[i].User;
		            	action.innerText = data.Data[i].Action;
		            	imei.innerText = data.Data[i].Imei;
		            	date.innerText = data.Data[i].Date;
		
		            	a.innerText = "详情";
		            	state.appendChild(a)
	            		
	            		li.appendChild(id);
	            		li.appendChild(user);
	            		li.appendChild(imei);
	            		li.appendChild(action);
	            		li.appendChild(date);
	            		li.appendChild(state);

	            		$(".wrap tr").last().after(li);
					}
				 	
				 },
				error: function(){
					alert("error")
				}
			});
		}
		
	});