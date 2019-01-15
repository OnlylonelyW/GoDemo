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

});