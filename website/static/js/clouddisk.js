$(document).ready(function() {
	$("#clouddisklink").addClass('active');
	$("#functionBlock a").click(function(event) {
		$($("#functionBlock .active").removeClass('active').data('linkblock')).hide();
		var contain = $($(this).addClass('active').data('linkblock')).show();
		myfileFuncClick(this);
	});

	switch (window.location.hash) {
		case "#myfileFunc":
			$("#myfileFunc").click();
			break;
		case "#myshareFunc":
			$("#myshareFunc").click();
			break;
		case "#myrecycledFunc":
			$("#myrecycledFunc").click();
			break;
		default:
			$("#myfileFunc").click();
	}


	InitButtonFunction();
	InitShareButtonFunc();
	InitRecycledButton();
})

function myfileFuncClick(event) {
	if ($(event).data('linkblock') == '#myfileToolbarContainer') {
		InitFilePath();
		InitFileDataGrid();
	} else if ($(event).data('linkblock') == '#myshareToolbarContainer') {
		InitShareDataGrid();
	} else if ($(event).data('linkblock') == '#myrecycledToolbarContainer') {
		InitRecycledDataGrid();
	}
}

function InitFilePath() {
	var setting = $("#filePathLink").data("TreePath");
	if (typeof setting == "object") {
		if (typeof setting.Destroy == "function")
			$("#filePathLink").TreePath("Destroy");
	}
	$("#filePathLink").TreePath({
		ItemList: [{
			Id: 0,
			Lc_fileName: "主目录"
		}],
		ItemFormat: function(value) {
			return value.Lc_fileName;
		},
		ItemClick: function(sender, event, value) {
			RequestFileByID();
		}
	})
}

function InitFileDataGrid() {
	var setting = $("#fileListMain").data("DataGrid");
	if (typeof setting == "object") {
		if (typeof setting.Destroy == "function")
			$("#fileListMain").DataGrid("Destroy");
	}

	var he = $(window).height()- $("#fileListMain")[0].getBoundingClientRect().top-5

	$("#fileListMain").DataGrid({
		Url: "/dataserver",
		RequestParm: {
			ActionType: 'GetFileList',
			FileID: 0
		},
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
		Height: he,
		CellClick: function() {},
		RowClick: function(send, event, data) {
			RowClick(send, event, data);
			ControlButtonView()
		},
		CheckboxClick: function(sender, event, input) {
			CheckboxClick(sender, event, input);
		},
		HeadCheckBoxClick: function(sender, event, input) {
			ControlButtonView()
		},
		RequestAfterRender: function() {
			ControlButtonView()
		}
	})
}

function RowClick(send, event, data) {
	var headcheckbox = $("#fileListMain").DataGrid("GetHeadCheckbox");
	if (headcheckbox.CheckBox("GetValue")) {
		headcheckbox.CheckBox("SetValue", false);
		$("#fileListMain").DataGrid("GetAllCheckedCheckbox").each(function(index, el) {
			$(this).CheckBox("SetValue", false);
		});
		$(send).find('.DatagridHideCheckbox').CheckBox("SetValue", true);
	} else {
		var input = $(send).find('.DatagridHideCheckbox');
		var value = input.CheckBox("GetValue");
		var checked = $("#fileListMain").DataGrid("GetAllCheckedCheckbox");
		checked.each(function(index, el) {
			$(this).CheckBox("SetValue", false);
		});
		if (checked.length > 1) {
			input.CheckBox("SetValue", true)
		} else {
			if (value) {
				input.CheckBox("SetValue", false)
			} else {
				input.CheckBox("SetValue", true)
			}
		}
	}
}

function InitButtonFunction() {
	//文件上传
	$("#uploadDlg").UploadFile({}).hide()
	$("#fileButtonUpload").click(function(event) {
		$("#uploadDlg").UploadFile("AlterUrl", "ws://" + location.host + "/upload?Action=FileCloud&Folderid=" + $("#filePathLink").TreePath("GetLastItem").Id)
		$("#uploadDlg").dialog({
			autoOpen: true,
			height: 300,
			width: 600,
			modal: true,
			close: function() {
				RequestFileByID();
			}
		}).focus();
	});

	//新建文件夹
	$("#fileButtonNewFolder").click(function(event) {
		var value = {
			Lc_fileName: "新建文件夹",
			Lc_fileSize: 0,
			Lc_createTime: ""
		};
		var row = $("#fileListMain").DataGrid("InsertTr", value)
		var td = row.find('[name="Lc_fileName"]');
		var commit = false;
		td.children().hide();
		var input = $("<input type='text' />").data("oldValue", value.Lc_fileName).change(function(event) {
			//在此处添加文件夹名称改变时的处理逻辑
			if (this.value.length == 0) {
				row.remove();
			} else {
				if (!commit) {
					CreateNewFolder.call(this);
					commit = !commit;
				}
			}
		}).focusout(function(event) {
			//在此处添加输入框失去焦点时的处理逻辑
			if (this.value.length == 0)
				row.remove();
			else {
				if (!commit) {
					CreateNewFolder.call(this);
					commit = !commit;
				}
			}
		}).appendTo(td).focus().val(value.Lc_fileName);

		function CreateNewFolder() {
			var option = new Object()
			option.url = "/dataserver";
			option.data = {
				ActionType: "AddFolder",
				FolderName: this.value,
				ParentID: $("#filePathLink").TreePath("GetLastItem").Id
			}
			option.success = function(value) {
				RequestFileByID();
			}
			option.error = function(value) {}
			AjaxCall(option);
		}
	});

	//重命名
	$("#fileButtonRename").click(function(event) {
		var row = $("#fileListMain").DataGrid("GetAllCheckedCheckbox").data("RowDom")
		var value = row.data("RowData");
		var td = row.find('[name="Lc_fileName"]');
		td.children().hide();
		//var li = $("#fileListMain").find("input:checked").parents('li').data("liStruct");
		var input = $("<input type='text' />").data("oldValue", value.Lc_fileName).change(function(event) {
			//名称改变时执行的语句
			if (this.value.length == 0) {
				input.remove();
				td.children().show();
			} else {
				FileRename.call(this)
			}
		}).focusout(function(event) {
			//失去焦点时执行语句
			input.remove();
			td.children().show();
		}).appendTo(td).focus().val(value.Lc_fileName);

		function FileRename() {
			var option = new Object()
			option.url = "/dataserver"
			option.data = {
				ActionType: "RenameFile",
				RecordID: value.Id,
				Rename: input.val(),
				ParentID: $("#filePathLink").TreePath("GetLastItem").Id
			}
			option.success = function(value) {
				RequestFileByID();
			}
			option.error = function(value) {}
			AjaxCall(option);
		}
	});
	//删除
	$("#fileButtonDelete").click(function(event) {
		var liList = $("#fileListMain").DataGrid("GetAllCheckedCheckbox")
		var ARR = [];
		liList.each(function(index, el) {
			ARR.push($(this).data("value").Id);
		});
		var option = new Object()
		option.url = "/dataserver";
		option.data = {
			ActionType: "DeleteFile",
			FileIdList: ARR,
			ParentID: $("#filePathLink").TreePath("GetLastItem").Id
		}
		option.success = function(data) {
			RequestFileByID();
		}
		option.error = function() {}
		AjaxCall(option);
	});

	//分享
	$("#fileButtonShare").click(function(event) {
		var liList = $("#fileListMain").DataGrid("GetAllCheckedCheckbox");
		var ARR = [];
		var ARRID = [];
		var ShareName, fileExt, FileSize, filetype;
		liList.each(function(index, el) {
			ARR.push($(this).data("value"));
			ARRID.push(ARR[ARR.length - 1].Id);
		}); //获取所有的分享数据集合
		if (ARR.length == 1) {
			if (ARR[0].Lc_fileType == 0) {
				ShareName = ARR[0].Lc_fileName;
				fileExt = ARR[0].Lc_fileExtension;
				filetype = 0
				filesize = ARR[0].Lc_fileSize;
			} else if (ARR[0].Lc_fileType == 1) {
				ShareName = ARR[0].Lc_fileName;
				fileExt = "";
				filetype = 1;
				filesize = 0;
			}
		} else {
			ShareName = ARR[0].Lc_fileName + "等" + ARR.length + "个文件";
			fileExt = "";
			filetype = 2;
			filesize = 0;

		}
		var option = new Object()
		option.url = "/dataserver";
		option.data = {
			ActionType: "ShareFile",
			FileIdList: ARRID,
			ParentID: $("#filePathLink").TreePath("GetLastItem").Id,
			shareName: ShareName,
			fileExt: fileExt,
			fileSize: filesize,
			fileType: filetype
		}
		option.success = function(data) {
			//需要在此处添加分享成功的代码
		}
		option.error = function() {}
		AjaxCall(option);
	});

	//搜索
	$("#searchFileButton").click(function(event) {
		//需要进行修改的
		var option = {};
		option.ActionType = 'GetFileList';
		option.SearchKey = $("#searchFile").val();
		$("#fileListMain").DataGrid("EditRequestParm", option);
		$("#fileListMain").DataGrid("RequestData")
	});
}

function CheckboxClick(sender, event, input) {
	var input = $("#fileListMain").DataGrid("GetHeadCheckbox");
	if (input.CheckBox("GetValue"))
		input.CheckBox("SetValue", false);
	ControlButtonView();
}

function ControlButtonView() {
	var checked = $("#fileListMain").DataGrid("GetAllCheckedCheckbox");
	if (checked.length != 0) {
		$("#fileButtonHide").show();
		if (checked.length == 1) {
			$("#fileButtonRename").show()
		} else {
			$("#fileButtonRename").hide()
		}
	} else {
		$("#fileButtonHide").hide();
	}
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
		$("#filePathLink").TreePath("PushItem", value);
		RequestFileByID();
	}
}

function RequestFileByID() {
	var option = {};
	option.ActionType = 'GetFileList';
	option.FileID = $("#filePathLink").TreePath("GetLastItem").Id;
	$("#fileListMain").DataGrid("EditRequestParm", option);
	$("#fileListMain").DataGrid("RequestData")
}

function SortFileOrder(parm1, parm2, key, order) {
	if (parm1.Lc_fileType == parm2.Lc_fileType) {
		if (parm1[key] > parm2[key]) {
			return (order == "ASC") ? 1 : -1;
		} else if (parm1[key] == parm2[key]) {
			return 0;
		} else {
			return (order == "ASC") ? -1 : 1;
		}
	} else {
		if (parm1.Lc_fileType == 0)
			return 1;
		else
			return -1;
	}
}

function RealizeShare() {
	InitShareDataGrid();
}

function InitShareDataGrid() {
	var setting = $("#fileListMain").data("DataGrid");
	if (typeof setting == "object") {
		if (typeof setting.Destroy == "function")
			$("#fileListMain").DataGrid("Destroy");
	}

	var he = $(window).height()- $("#fileListMain")[0].getBoundingClientRect().top-5

	$("#fileListMain").DataGrid({
		Url: "/dataserver",
		RequestParm: {
			ActionType: 'GetShareList'
		},
		Columns: [{
			Sortable: true,
			field: "ShareName",
			title: "分享名",
			theadStyle: "ShareColPrecent",
			Formatter: function(value, data, column) {
				var div = $("<div></div>")
				var alink = $("<a href='/sharefileview/"+data.ShareId+"'></a>").html( (value.length>20)?value.substr(0,20)+"...":value).click(function(event) {
					/* Act on the event */
					event.stopPropagation()
				});;
				var icon = $("<span></span>").addClass('icon')
				icon.addClass(FileIcon(data.Lc_fileType, data.Lc_fileExtension))
				div.append(icon).append(alink);
				return div;
			},
			Sort: SortShareOrder
		}, {
			Sortable: true,
			field: "ShareCode",
			title: "分享码",
			theadStyle: "ShareCodePrecent",
			Sort: SortShareOrder
		}, {
			Sortable: true,
			field: "DownloadCount",
			title: "下载次数",
			theadStyle: "ShareDownloadPrecent",
			Formatter: function(value, data, column) {
				return FormartDateTime(value)
			},
			Sort: SortShareOrder
		}, {
			Sortable: true,
			field: "ShareSize",
			title: "大小",
			theadStyle: "ShareSizePrecent",
			Formatter: function(value, data, column) {
				return value
			},
			Sort: SortShareOrder
		}, {
			Sortable: true,
			field: "ShareTime",
			title: "日期",
			theadStyle: "ShareDatePrecent",
			Formatter: function(value, data, column) {
				return FormartDateTime(value)
			},
			Sort: SortShareOrder
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
		Height: he,
		CellClick: function() {},
		RowClick: function(send, event, data) {
			RowClick(send, event, data);
			ShareButtonView()
		},
		CheckboxClick: function(sender, event, input) {
			ShareCheckboxClick(sender, event, input);
		},
		HeadCheckBoxClick: function(sender, event, input) {
			HeadCheckboxClick(sender, event, input)
		}
	})
}

function SortShareOrder(parm1, parm2, colum) {
	if (parm1[colum.field] > parm2[colum.field])
		return 1
	else if (parm1[colum.field] == parm2[colum.field])
		return 0
	else
		return -1
}

function FormartDateTime(value) {
	if (value == undefined || value == null || value == "")
		return "";
	var datetime = new Date(value);
	return datetime.getFullYear() + "/" + (datetime.getMonth() + 1) + '/' + datetime.getDate() + " " + datetime.getHours() + ":" + datetime.getMinutes();
}

function ShareCheckboxClick(sender, event, input) {
	var input = $("#fileListMain").DataGrid("GetHeadCheckbox");
	if (input.CheckBox("GetValue"))
		input.CheckBox("SetValue", false);
	ShareButtonView()
}

function HeadCheckboxClick(sender, event, input) {
	ShareButtonView()
}

function ShareButtonView() {
	var checked = $("#fileListMain").DataGrid("GetAllCheckedCheckbox");
	if (checked.length != 0) {
		if (checked.length == 1) {
			$("#gotoSharePage").prop("disabled", false);
			$("#cancelShare").prop("disabled", false);
		} else {
			$("#gotoSharePage").prop("disabled", true);
			$("#cancelShare").prop("disabled", false);
		}
	} else {
		$("#gotoSharePage").prop("disabled", true);
		$("#cancelShare").prop("disabled", true);
	}
}

function InitShareButtonFunc() {
	$("#gotoSharePage").click(function(event) {
		//转到文件分享页面
	});
	$("#cancelShare").click(function(event) {
		var liList = $("#fileListMain").DataGrid("GetAllCheckedCheckbox")
		var ARR = [];
		liList.each(function(index, el) {
			ARR.push($(this).data("value").ShareId);
		});

		var option = new Object()
		option.url = "/dataserver";
		option.data = {
			ActionType: "CancelShareRecord",
			FileIdList: ARR
		}
		option.success = function(data) {
			$("#fileListMain").DataGrid("RequestData");
		}
		option.error = function() {}
		AjaxCall(option);
	});
}

function InitRecycledDataGrid() {
	var setting = $("#fileListMain").data("DataGrid");
	if (typeof setting == "object") {
		if (typeof setting.Destroy == "function")
			$("#fileListMain").DataGrid("Destroy");
	}

	var he = $(window).height()- $("#fileListMain")[0].getBoundingClientRect().top-5

	$("#fileListMain").DataGrid({
		Url: "/dataserver",
		RequestParm: {
			ActionType: 'GetRecycledList'
		},
		Columns: [{
			Sortable: true,
			field: "Lc_fileName",
			title: "文件名",
			theadStyle: "col-md-6",
			Formatter: function(value, data, column) {
				var div = $("<div></div>")
				//var alink = $("<a href='/sharefileview/"+data.ShareId+"'></a>").html( (value.length>20)?value.substr(0,20)+"...":value).click(function(event) {
				//	event.stopPropagation()
				//});;
				var icon = $("<span></span>").addClass('icon')
				icon.addClass(FileIcon(data.Lc_fileType, data.Lc_fileExtension))
				div.append(icon).append(value);
				return div;
				//return value;
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
				if (data.Data == 0) {
					location.href = "/"
				}
			}
		},
		ShowCheckBox: true,
		ShowPagination: false,
		Height: he,
		CellClick: function() {},
		RowClick: function(send, event, data) {
			RowClick(send, event, data);
			RecycledButtonView();
		},
		CheckboxClick: function(sender, event, input) {
			RecycledCheckboxClick(sender, event, input)
		},
		HeadCheckBoxClick: function(sender, event, input) {
			RecycledButtonView()
		}
	})
}

function RecycledCheckboxClick(sender, event, input) {
	var input = $("#fileListMain").DataGrid("GetHeadCheckbox");
	if (input.CheckBox("GetValue"))
		input.CheckBox("SetValue", false);
	RecycledButtonView()
}

function RecycledButtonView() {
	var checked = $("#fileListMain").DataGrid("GetAllCheckedCheckbox");
	if (checked.length != 0) {
		$("#ShiftDeleteFile").prop("disabled", false);
		$("#ReBackFile").prop("disabled", false);
	} else {
		$("#ShiftDeleteFile").prop("disabled", true);
		$("#ReBackFile").prop("disabled", true);
	}
}

function InitRecycledButton() {
	$("#ShiftDeleteFile").click(function(event) {
		var liList = $("#fileListMain").DataGrid("GetAllCheckedCheckbox")
		var ARR = [];
		liList.each(function(index, el) {
			ARR.push($(this).data("value").ShareId);
		});

		var option = new Object()
		option.url = "/dataserver";
		option.data = {
			ActionType: "CancelShareRecord",
			FileIdList: ARR
		}
		option.success = function(data) {
			$("#fileListMain").DataGrid("RequestData");
		}
		option.error = function() {}
		AjaxCall(option);
	});
	$("#ReBackFile").click(function(event) {
		var liList = $("#fileListMain").DataGrid("GetAllCheckedCheckbox")
		var ARR = [];
		liList.each(function(index, el) {
			ARR.push($(this).data("value").ShareId);
		});

		var option = new Object()
		option.url = "/dataserver";
		option.data = {
			ActionType: "ReBackFile",
			FileIdList: ARR
		}
		option.success = function(data) {
			$("#fileListMain").DataGrid("RequestData");
		}
		option.error = function() {}
		AjaxCall(option);
	});
	$("#ClearRecycleBin").click(function(event) {
		var option = new Object()
		option.url = "/dataserver";
		option.data = {
			ActionType: "ClearRecycleBin",
		}
		option.success = function(data) {
			$("#fileListMain").DataGrid("RequestData");
		}
		option.error = function() {}
		AjaxCall(option);
	});
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

//传入参数对象
//url:请求的地址
//data为请求传入的参数数据
//success请求成功后执行的函数
//error请求失败后执行的函数
function AjaxCall(option) {
	$.ajax({
		url: option.url,
		type: 'POST',
		dataType: 'json',
		data: option.data,
		success: function(value) {
			if (value.RequestResult) {
				option.success.call(this, value, option);
			} else {
				if (data.Data == 0) {
					location.href = "/"
					return
				}
			}
		},
		error: function(value) {
			option.error.call(this, value, option)
		}
	});
}