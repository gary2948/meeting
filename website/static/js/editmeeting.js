$(document).ready(function() {
	$("#meetinglink").addClass('active');
	//日历实现
	$("#calendar").datepicker({
		showOtherMonths: true,
		minDate: 0,
		maxDate: 0
	});
	Localization($)
	$("#calendar").datepicker("option", $.datepicker.regional["zh-TW"]);
	$("#calendar").find('.ui-datepicker-header a').hide()

	//实现tab标签页
	$("#container>div").hide();
	$("#meetingInfoTab").click(function(event) {
		$(this).parents("ul").children('.active').removeClass('active')
		$(this).parent().addClass('active')
		$("#container>div").hide();
		$("#meetingInfo").show();
	});
	$("#meetingLogTab").click(function(event) {
		$(this).parents("ul").children('.active').removeClass('active')
		$(this).parent().addClass('active')
		$("#container>div").hide();
		$("#meetingLog").show();
	});
	$("#meetingInfoTab").click()

	$("#MeetingFileList").BlockItemList({
		InfoFormat: function(value) {
			return value;
		},
		BtnFormat: function(value) {
			return "x"
		}
	});
	$("#MeetingInviterList").BlockItemList({
		InfoFormat: function(value) {
			return value;
		},
		BtnFormat: function(value) {
			return "x"
		}
	});

	$("#uploadDlg").UploadFile({
		UploadFinished: function(file, data) {}
	}).hide()
	$("#uploadFileClick").click(function(event) {
		$("#uploadDlg").UploadFile("AlterUrl", "ws://" + location.host + "/upload?Action=Meeting")
		$("#uploadDlg").dialog({
			autoOpen: true,
			height: 300,
			width: 600,
			modal: true,
			close: function() {
				var items = $("#uploadDlg").UploadFile("GetItems");

				for (var i = 0; i < items.length; i++) {
					if (items[i].UploadState == 4) {
						$("#MeetingFileList").BlockItemList("AddItem", items[i].FinishedData);
					}
				}
				$("#uploadDlg").UploadFile("ClearItems");
			}
		}).focus();
	});

	$("#inputInvitee").keypress(function(event) {
		if (event.key == "Enter") {
			var value = $(this).val().replace(/[ \f\n\r\v]{1,}|\t{2,}/, " ").split(" ");
			for (var i = 0; i < value.length; i++) {
				if (!value[i].match(/^([a-zA-Z0-9_\.-])+@([a-zA-Z0-9_-])+((\.[a-zA-Z0-9_-]{2,3}){1,2})$/)) {
					$(this).parent().addClass('has-error');
					return;
				}
			}
			$("#MeetingInviterList").BlockItemList("AddItem", value)
			$(this).val("");
		} else {
			$(this).parent().removeClass('has-error')
		}
	});

	$("#saveChange").click(function(event) {
		var option={}
		option.Lc_topic=$("#Lc_topic").val();
		option.Lc_schema=$("#Lc_schema").val();
		option.Files=$("#MeetingFileList").BlockItemList("GetItems");
		option.MeetingInviterList=$("#MeetingInviterList").BlockItemList("GetItems");
		option.Password=$("#password").val();
		option.ActionType=""
		$.ajax({
			url: '/path/to/file',
			type: 'default GET (Other values: POST)',
			dataType: 'default: Intelligent Guess (Other values: xml, json, script, or html)',
			data: {param1: 'value1'},
		});
		
	});

});


function Localization($) {
	$.datepicker.regional['zh-TW'] = {
		closeText: '关闭',
		prevText: '&#x3C;上月',
		nextText: '下月&#x3E;',
		currentText: '今天',
		monthNames: ['一月', '二月', '三月', '四月', '五月', '六月',
			'七月', '八月', '九月', '十月', '十一月', '十二月'
		],
		monthNamesShort: ['一月', '二月', '三月', '四月', '五月', '六月',
			'七月', '八月', '九月', '十月', '十一月', '十二月'
		],
		dayNames: ['星期日', '星期一', '星期二', '星期三', '星期四', '星期五', '星期六'],
		dayNamesShort: ['周日', '周一', '周二', '周三', '周四', '周五', '周六'],
		dayNamesMin: ['日', '一', '二', '三', '四', '五', '六'],
		weekHeader: '周',
		dateFormat: 'yy/mm/dd',
		firstDay: 1,
		isRTL: false,
		showMonthAfterYear: true,
		yearSuffix: '年'
	};
	$.datepicker.setDefaults($.datepicker.regional['zh-TW']);
}