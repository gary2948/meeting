/// <reference path="jquery-1.10.2.js" />

(function($) {
    $.fn.UploadFile = function(option, parm) {
        var result = this;

        if (typeof option === "string") {
            if (parm == undefined) {
                return $(this[0]).data("FileUpload")[option]()
            }
            return $(this[0]).data("FileUpload")[option](parm)
        } else if (typeof option === "object") {
            return $(this).each(function() {
                UploadFile($(this), option)
            })
        }


        /*
        $(this).each(function(index, el) {
            if (typeof option === "object") {
                UploadFile($(this), option);
            } else {
                if (FunctionObject[option] !== undefined) {
                    var val = FunctionObject[option].call(this, parm)
                    result = (val == null) ? result : val;
                }

            }
        });
        return result;*/
    }
    var UnUpload = 0,
        StartUpload = 1,
        Stop = 2,
        Cancel = 3,
        Finish = 4,
        UploadError = 5;

    function UploadFile(Dom, option) {

        var FileUploadModel = CreateDefaultObject();
        $.extend(FileUploadModel, option); //文件上传模型

        FileUploadModel.Dom = Dom;
        CreateViewFrame();
        AddEvent();
        BindFunction();
        Dom.data("FileUpload", FileUploadModel);


        //设置默认参数
        function CreateDefaultObject() {
            var DefaultOption = {
                Url: "ws://localhost:8080/upload?",
                SelectFileBtn: "选择文件",
                Stop: "暂停",
                Start: "上传",
                Cancel: "取消",
                Finish: "完成",
                multiple: true,
                DirectUpload: true,
                FileList: [],
                view: {},
                UploadFinished: null
            }
            return DefaultOption;
        }

        //创建框架
        function CreateViewFrame() {
            var head = $("<div></div>").css("position", 'relative');
            var btn = $("<botton></botton>").addClass('btn btn-primary').html(FileUploadModel.SelectFileBtn)
            var input = $("<input type='file'/>").css({
                position: "absolute",
                opacity: 0,
                left: 0,
                top: 0
            });
            if (FileUploadModel.multiple) {
                input.attr("multiple", "multiple")
            }
            var viewPanle = $("<div></div>");
            head.append(btn).append(input);

            Dom.append(head).append(viewPanle);
            input.width(btn.outerWidth()).height(btn.outerHeight());
            var fileListView = CreateViewPanel(FileUploadModel.FileList);
            viewPanle.append(fileListView);

            FileUploadModel.view.head = {};
            FileUploadModel.view.head.HeadDom = head;
            FileUploadModel.view.head.input = input;
            FileUploadModel.view.viewPanle = {};
            FileUploadModel.view.viewPanle.viewPanle = viewPanle;
            FileUploadModel.view.viewPanle.FileListView = fileListView
        }

        function CreateViewPanel(data) {
            var ul = $("<ul></ul>").addClass("list-group");
            for (var i = 0; i < data; i++) {
                ul.append(CreateViewItem(data[i]));
            }
            return ul;
        }

        function CreateViewItem(file) {
            var li = $("<li></li>").addClass("list-group-item");
            var divInfo = $("<div></div>").addClass("row");

            var divInfoFileName = $("<div></div>").addClass("col-md-8").append($("<span></span>").html(file.name));

            var divInfoFunc = $("<div></div>").addClass("col-md-4");
            var Upload = $("<button></button>").addClass("btn btn-primary navbar-right").html(FileUploadModel.Start);
            var Stop = $("<button></button>").addClass("btn btn-primary navbar-right").html(FileUploadModel.Stop).hide();
            var Cancel = $("<button></button>").addClass("btn btn-primary navbar-right").html(FileUploadModel.Cancel);
            var Finish = $("<button></button>").addClass("btn btn-primary navbar-right").html(FileUploadModel.Finish).hide();
            divInfoFunc.append(Cancel).append(Finish).append(Stop).append(Upload);
            divInfo.append(divInfoFileName).append(divInfoFunc);

            var progress = $("<div></div>").addClass("progress").css({
                "margin-top": "5px",
                "margin-bottom": "0px",
                height: "3px"
            });
            var progressValue = $("<div></div>").addClass("progress-bar").css("width", "0%");
            progress.append(progressValue);

            li.append(divInfo).append(progress);


            Upload.click(function() {
                file.UploadState = StartUpload;
                Cancel.hide();
                Upload.hide();
                var callback = {};
                callback.Progress = function() {
                    progressValue.css("width", file.Percent * 100 + "%");
                };
                UploadFile(file, callback);
            });
            if (FileUploadModel.DirectUpload) {
                Upload.click();
            }
            Stop.click(function() {});
            Cancel.click(function() {
                if (file.socket != null && file.socket != undefined) {
                    file.socket.close();
                }
                li.remove();
                var index = FileUploadModel.FileList.indexOf(file, 0);
                FileUploadModel.FileList.splice(index, 1);
            });


            var liModel = {};
            liModel.li = li;
            liModel.Upload = Upload;
            liModel.Stop = Stop;
            liModel.Cancel = Cancel;
            liModel.Finish = Finish;
            liModel.Progress = progress;
            liModel.ProgressValue = progressValue;
            liModel.FileModel = file;
            li.data("liModel", liModel);
            return li;
        }

        function FileSelectChange() {
            var array = new Array();
            for (var i = 0; i < this.files.length; i++) {
                array[i] = CreateFileListItem(this.files[i]);
                FileUploadModel.FileList.push(array[i]);
            }
            AddListViewItem.call(FileUploadModel.view.viewPanle.FileListView, array);
        }

        function AddListViewItem(files) {
            for (var i = 0; i < files.length; i++) {
                FileUploadModel.view.viewPanle.FileListView.append(CreateViewItem(files[i]))
            }
        }

        function CreateFileListItem(file) {
            file.Percent = 0;
            file.UploadState = UnUpload;
            return file;
        }

        function AddEvent() {
            FileUploadModel.view.head.input.change(function(event) {
                FileSelectChange.call(this, event)
            });
        }

        function UploadFile(file, Option) {
            var socket = new WebSocket(FileUploadModel.Url + "&FileName=" + encodeURIComponent(file.name));
            file.socket = socket;
            socket.onmessage = function(event) {
                var data = JSON.parse(event.data);
                if (data.Code == 0) {
                    file.Percent = data.Data / file.size;
                    Option.Progress();
                } else if (data.Code == 1) {
                    if (file.Percent == 1) {
                        //上传完成
                        if (typeof FileUploadModel.UploadFinished == "function") {
                            FileUploadModel.UploadFinished(file, data.Data);
                        }
                        file.FinishedData=data.Data;
                        file.UploadState = Finish;
                    } else {
                        //上传出错
                        file.UploadState = UploadError;
                    }
                    socket.close();
                }
            }
            socket.onerror = function() {}
            socket.onopen = function() {
                socket.send(file);
            }
        }

        function HasUpload() {
            var list = FileUploadModel.FileList;
            for (var i = 0; i < list.length; i++) {
                if (list[i].UploadState == StartUpload)
                    return true
            }
            return false;
        }

        function AlterUrl(url) {
            FileUploadModel.Url = url
        }

        function ClearItems() {
            var model = FileUploadModel; //$(this).data("FileUpload");
            var arr = model.FileList;
            for (var i = 0; i < arr.length; i++) {
                arr[i].socket.close();
            }
            model.FileList = [];
            $(model.view.viewPanle.viewPanle).empty();
            model.view.viewPanle.FileListView = CreateViewPanel(model.FileList);
            $(model.view.viewPanle.viewPanle).append(model.view.viewPanle.FileListView);
        }

        function Distroy() {
            ClearItems();
            $(Dom).removeData('FileUpload');
            $(Dom).empty();
        }

        function GetItems() {
            return FileUploadModel.FileList;
        }

        function BindFunction() {
            FileUploadModel.HasUpload = HasUpload;
            FileUploadModel.AlterUrl = AlterUrl;
            FileUploadModel.ClearItems = ClearItems;
            FileUploadModel.GetItems = GetItems;
            FileUploadModel.HasUpload = HasUpload;
        }

        /*var FunctionObject = {
            HasUpload: HasUpload,
            AlterUrl: AlterUrl,
            ClearItems: ClearItems,
            GetItems: GetItems
        }*/
    }



})(jQuery);