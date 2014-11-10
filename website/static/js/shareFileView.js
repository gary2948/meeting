$(document).ready(function(){
	$("#filePathLink").TreePath({
		ItemList: [{
			Id: $("#filePathLink").data("defaultid"),
			Title: "主目录",
			ActionType:$("#filePathLink").data("defaultactiontype")
		}],
		ItemFormat: function(value) {
			return value.Title;
		},
		ItemClick: function(sender, event, value) {
			RequestFileByID();
		}
	})

	InitFileDataGrid();
})

function RequestFileByID() {
	//var option = {};
	//option.ActionType = 'GetFileList';
	//option.FileID = $("#filePathLink").TreePath("GetLastItem").Id;
	$("#datagrid").DataGrid("EditRequestParm", $("#filePathLink").TreePath("GetLastItem"));
	$("#datagrid").DataGrid("RequestData")
}
function FileFileNameFormat(value, data, column) {
	var div = $("<div></div>")
	var alink = $("<a href='javascript:void(0)'></a>").html(value).click(function(event) {
		FileNameClick($(this), event, data);
	});
	var icon = $("<span></span>").addClass('icon')
	icon.addClass(FileIcon(data.Lc_fileType, data.Lc_fileExtension))
	div.append(icon).append(alink);
	return div;
}
//点击文件名时执行的函数
function FileNameClick(sender, event, value) {
	if (value.Lc_fileType === 1) {
		FolderExcute(sender)
	} else {
		FileExcute(sender)
	}
	event.stopPropagation()

	function FileExcute() {
		alert("file click");

	}

	function FolderExcute() {
		var opt={};
		opt.Title=value.Lc_fileName
		opt.ActionType="GetFileList";
		opt.FileID=value.Id;
		$("#filePathLink").TreePath("PushItem", opt);
		RequestFileByID();
	}
}

function FileIcon(filetype, value) { //filetype用于确定是文件夹还是文件 value为文件的类型

	if (filetype == 1) {
		return "fileExt-folder"
	} else {
		if (HasIcon(value))
			return "fileExt-" + value;
		else
			return "fileExt-" + "none";
	}
}

function HasIcon(value) {
	var extArr = ['ac3', 'ace', 'ade', 'adp', 'ai', 'aiff', 'au', 'avi', 'bat', 'bin', 'bmp', 'bup', 'cab', 'cat', 'chm', 'css', 'cue', 'dat', 'dcr', 'default', 'der', 'dic', 'divx', 'diz', 'dll', 'doc', 'docx', 'dos', 'dvd', 'dwg', 'dwt', 'emf', 'exc', 'fon', 'gif', 'hlp', 'html', 'ifo', 'inf', 'ini', 'ins', 'ip', 'iso', 'isp', 'java', 'jfif', 'jpeg', 'jpg', 'log', 'm4a', 'mid', 'mmf', 'mmm', 'mov', 'movie', 'mp2', 'mp2v', 'mp3', 'mp4', 'mpe', 'mpeg', 'mpg', 'mpv2', 'nfo', 'pdd', 'pdf', 'php', 'png', 'ppt', 'pptx', 'psd', 'rar', 'reg', 'rtf', 'scp', 'theme', 'tif', 'tiff', 'tlb', 'ttf', 'txt', 'uis', 'url', 'vbs', 'vcr', 'vob', 'wav', 'wba', 'wma', 'wmv', 'wpl', 'wri', 'wtx', 'xls', 'xlsx', 'xml', 'xsl', 'zap', 'zip'];
	for (var i = 0; i < extArr.length; i++) {
		if (extArr[i] === value)
			return true;
	}
	return false;
}

function FormartDateTime(value) {
	if (value == undefined || value == null || value == "")
		return "";
	var datetime = new Date(value);
	return datetime.getFullYear() + "/" + (datetime.getMonth() + 1) + '/' + datetime.getDate() + " " + datetime.getHours() + ":" + datetime.getMinutes();
}


function InitFileDataGrid() {
	var setting = $("#datagrid").data("DataGrid");
	if (typeof setting == "object") {
		if (typeof setting.Destroy == "function")
			$("#datagrid").DataGrid("Destroy");
	}
	
	$("#datagrid").DataGrid({
		Url: "/dataserver",
		RequestParm: $.extend({},$("#filePathLink").TreePath("GetLastItem")),
		Columns: [{
			Sortable: true,
			field: "Lc_fileName",
			title: "文件名",
			theadStyle: "col-md-6",
			Formatter: function(value, data, column) {
				return FileFileNameFormat(value, data, column);
			},
			SortASC: function(parm1, parm2, column) {
				return SortFileOrder(parm1, parm2, "Lc_fileName", "ASC")
			},
			SortDESC: function(parm1, parm2, column) {
				return SortFileOrder(parm1, parm2, "Lc_fileName", "DESC")
			},

		}, {
			Sortable: true,
			field: "Lc_fileSize",
			title: "文件大小",
			theadStyle: "col-md-2",
			SortASC: function(parm1, parm2, column) {
				return SortFileOrder(parm1, parm2, "Lc_fileSize", "ASC")
			},
			SortDESC: function(parm1, parm2, column) {
				return SortFileOrder(parm1, parm2, "Lc_fileSize", "DESC")
			},
		}, {
			Sortable: true,
			field: "Lc_createTime",
			title: "创建日期",
			theadStyle: "col-md-3",
			Formatter: function(value, data, column) {
				return FormartDateTime(value)
			},
			SortASC: function(parm1, parm2, column) {
				return SortFileOrder(parm1, parm2, "Lc_createTime", "ASC")
			},
			SortDESC: function(parm1, parm2, column) {
				return SortFileOrder(parm1, parm2, "Lc_createTime", "DESC")
			},
		}],
		RepostDataFilter: function(data) {
			if (data.RequestResult) {
				if (data.Data == null || data.Data == undefined) {
					data.total = 0;
					data.rows = [];
					return data;
				} else {
					data.total = data.Data.length;
					data.rows = data.Data;
					return data;
				}
			} else {
				if (data.Data == 0)
					location.href = "/"
			}
		},
		ShowCheckBox: true,
		ShowPagination: false,
		Height: "300",
		CellClick: function() {},
		RowClick: function(send, event, data) {
			//RowClick(send, event, data);
			//ControlButtonView()
		},
		CheckboxClick: function(sender, event, input) {
			//CheckboxClick(sender, event, input);
		},
		HeadCheckBoxClick: function(sender, event, input) {
			//ControlButtonView()
		},
		RequestAfterRender: function() {
			//ControlButtonView()
		}
	})
}