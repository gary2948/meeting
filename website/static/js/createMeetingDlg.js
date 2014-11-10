/// <reference path="f:\study\webapplication1\webapplication2\js\jquery-1.10.2.js" />
var CloseDlgCallBack;
var CommitInfoCallBack;

function CreateMeetingDlg(send, event) {
	$("#CreateMeetingDlg").dialog({
		autoOpen: true,
		height: 400,
		width: 800,
		modal: true,
		close: function() {
			if (typeof CloseDlgCallBack === "function")
				CloseDlgCallBack.call()
		}
	})
}

function CreateMeetingTimeList(Dom) {
	for (var i = 8 * 60; i < 8 * 60 + 48 * 30; i += 30) {
		var li = $("<li></li>");
		var a = $("<a></a>");
		a.attr("href", "javascript:void(0)");
		var i1 = i % (48 * 30);
		a.html(((parseInt(i1 / 60) > 9) ? parseInt(i1 / 60) : "0" + parseInt(i1 / 60)) + ":" + ((i1 % 60 == 0) ? "00" : i1 % 60))
		a.click(function() {
			$("#meetingTimeValue").html($(this).html())
		})
		li.append(a);
		Dom.append(li);
	}
}

function CreateMeetingLengthHourList(Dom) {
	for (var i = 0; i < 24; i++) {
		var li = $("<li></li>");
		var a = $("<a></a>");
		a.attr("href", "javascript:void(0)");

		a.html((i > 9) ? i : "0" + i);
		a.click(function() {
			$("#meetingLengthHour").html($(this).html())
		})
		li.append(a);
		Dom.append(li);
	}
}

function CreateMeetingLengthMinList(dom) {
	for (var i = 0; i < 60; i++) {
		var li = $("<li></li>");
		var a = $("<a></a>");
		a.attr("href", "javascript:void(0)");

		a.html((i > 9) ? i : "0" + i);
		a.click(function() {
			$("#meetingLengthMin").html($(this).html())
		})
		li.append(a);
		dom.append(li);
	}
}


function SetDate(date) {
	$("#meetingdate").val(date.getFullYear() + "-" + ((date.getMonth() > 8) ? (date.getMonth() + 1).toString() : "0" + (date.getMonth() + 1)).toString() + "-" + (date.getDate() > 9 ? date.getDate().toString() : "0" + date.getDate().toString()));
}

function SetTime(date) {
	var mins = date.getMinutes();
	if (mins >= 30) {
		if (mins - 30 > 15) {
			date.setHours(date.getHours() + 1)
			date.setMinutes(0)
		} else {
			date.setMinutes(30)
		}
	} else {
		if (30 - mins > 15) {
			date.setMinutes(0)
		} else {
			date.setMinutes(30)
		}
	}
}


(function($) {
	$(document).ready(function() {

		$("#meetingdate").datepicker({
			dateFormat: "yy-mm-dd",
			minDate: 0
		});
		var date = new Date();
		SetDate(date);

		var meetingTimeList = $("#meetingTimeList");
		CreateMeetingTimeList(meetingTimeList);

		var meetingLengthHourList = $("#meetingLengthHourList");
		CreateMeetingLengthHourList(meetingLengthHourList)

		var meetingLengthMinList = $("#meetingLengthMinList");
		CreateMeetingLengthMinList(meetingLengthMinList)

		$("#createMeetingClick").click(function(event) {
			var date = $("#meetingdate").datepicker("getDate");
			var time = $("#meetingTimeValue").html();
			var arr = time.split(":");
			date.setHours(date.getHours() + parseInt(arr[0]), date.getMinutes() + parseInt(arr[1]));


			$.ajax({
				url: '/dataserver',
				type: 'POST',
				dataType: 'json',
				data: {
					ActionType: 'CreateMeeting',
					Topic: $("#meetingtopic").val(),
					Schema: $("#meetingschema").val(),
					BeginTime: date.getTime(),
					EndTime: function() {
						date.setHours(date.getHours() + parseInt($("#meetingLengthHour").html()), date.getMinutes() + parseInt($("#meetingLengthMin").html()))
						return date.getTime()
					},
					PassWord: $("#meetingPassword").val()
				},
				success: function(value) {
					if (!value.ResultModel && value.Data == 0)
						location.href = "/"

					$("#CreateMeetingDlg").dialog("close");
					if (typeof CommitInfoCallBack === "function")
						CommitInfoCallBack()
				},
				error: function(value) {

				}
			})
		});

	});
})(jQuery)