$(document).ready(function() {
	$("#selectFile").change(function(event) {
		/* Act on the event */
		var file = this.files[0];
		if(!/image\/\w+/.test(file.type)){
			alert("请选择图像文件");
			return;
		}
		$("#filename").val(file.name)
		var reader = new FileReader();
        reader.readAsDataURL(file);
        reader.onload = function (e) {
        	$("#imageView img").attr('src', this.result);
        }
	}).width(function(){
		return $(this).parent().width()
	}).height(function(){
		return $(this).parent().height()
	}).css({
		left:function(){
			return $(this).parent().css("padding-left")
		},
		top:function(){
			return $(this).parent().css("padding-top")
		}
	});
});