$(document).ready(function($) {
	$("#homelink").addClass('active')


	$.ajax({
		url: '/dataserver',
		type: "post",
		dataType: "json",
		data: {
			ActionType: "DynamicInfo"
		},
		success: function(value) {
			if (value.RequestResult) {
				if(value.Data!=undefined && value.Data!=null)
				for (var i = 0; i < value.Data.length; i++) {
					$("#dynamicInfolist").append(CreateListItem(value.Data[i]))
				}
			} else {

			}

		},
		error: function(value) {}
	});
})

function CreateListItem(value) {
	var node = $("<div></div>").addClass('list-group-item')
	var row = $("<div></div>").addClass('row')
	var left = $("<div></div>").addClass('col-md-2')
	var right = $("<div></div>").addClass('col-md-10')
	left.append($("<img/>").attr('src', value.IconAddr).css('width', '100%'))
	var top = $("<div></div>")
	var mid = $("<div></div>")
	//var bottom = $("<div></div>")
	top.append($("<label></label>").html(value.NickName))//.append($("<label></label>").addClass('navbar-right').html("会议回放"))
	var p1=$("<p></p>");
	if(value.MeetingAction==0){
		p1.html("创建了会议:"+value.MeetingTopic)
	}else if(value.MeetingAction==1){
		p1.html("开启了会议:"+value.MeetingTopic)
	}else if(value.MeetingAction==4){
		p1.html("邀请您参加会议:"+value.MeetingTopic)
	}
	var p2 =$("<p></p>")
	p2.html("会议码:<a href=''>"+value.MeetingCode+"</a>")
	var p3=$("<p></p>")
	p3.html("时间:"+FormartDateTime(value.MeetingActionDate))
	mid.append(p1).append(p2).append(p3)
	//mid.html(value.info)
	//bottom.append($("<span></span>").addClass("navbar-right").html("回复"))
	right.append(top).append(mid)//.append(bottom)
	row.append(left).append(right)
	node.append(row)
	return node;

}


function FormartDateTime(value) {
	if (value == undefined || value == null || value == "")
		return "";
	var datetime = new Date(value);
	return datetime.getFullYear() + "/" + (datetime.getMonth() + 1) + '/' + datetime.getDate() + " " + datetime.getHours() + ":" + datetime.getMinutes();
}