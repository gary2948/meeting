/// <reference path="static/js/jquery-2.1.1.js" />
/// <reference path="CheckBox.js" />
/// <reference path="static/js/jquery.bs_pagination.js" />

(function($) {
    $.fn.DataGrid = function(option, value) {
        if (typeof option === "string") {
            if (value == undefined) {
                return $(this[0]).data("DataGrid")[option]()
            }
            return $(this[0]).data("DataGrid")[option](value)
        } else if (typeof option === "object") {
            return $(this).each(function() {
                DataGrid($(this), option)
            })
        }
    }

    function DataGrid(Dom, option) {

        var defaultSetting = DefaultSetting();
        var Setting = $.extend({}, defaultSetting, option);
        Setting.defaultSetting = defaultSetting;
        Setting.userSetting = option;


        CreateDomStruct();
        RequestData();
        BindFunction();
        $(Dom).data("DataGrid", Setting);

        function RequestData() {
            var parm
            if (Setting.ShowPagination) {
                parm = $.extend({}, Setting.RequestParm)
                parm.CurrentPage = Setting.Pagination.currentPage;
                parm.RowPerPage = Setting.Pagination.rowsPerPage;
            } else {
                parm = $.extend({}, Setting.RequestParm)
            }
            $.ajax({
                url: Setting.Url,
                type: 'POST',
                dataType: 'json',
                data: parm,
                success: function(value) {
                    value = Setting.RepostDataFilter(value);
                    Setting.TotalRows = value.total;
                    Setting.Pagination.totalPages = Math.ceil(value.total / Setting.Pagination.rowsPerPage),
                    Setting.Pagination.totalRows = value.total;
                    Setting.Rows = value.rows;
                    RefreshDataViewHead();
                    RefreshDataView();
                    CreatPagination();
                    RequestAfterRender();
                },
                error: function(value) {}
            });
        }


        function DefaultSetting() {
            var setting = {
                Url: null, //请求的服务地址
                RequestParm: null, //请求服务时附加的参数
                ShowCheckBox: false, //是否显示复选框按钮
                ShowRowNum: false, //是否显示行号
                ShowPagination: true, //是否显示行号
                Columns: [], //列集合
                Pagination: {
                    currentPage: 1,
                    rowsPerPage: 10,
                    totalPages: 0,
                    totalRows: 0,
                }, //分页的参数
                Rows: [], //行集合
                TotalRows: 0, //总行数
                //Order: "ASC",//升序还是降序
                //OrderBy:"",//按照那一列进行排序

                RepostDataFilter: function(value) {
                    return value;
                }, //对服务器端发送过来的数据进行过滤的函数
                CellClick: function() {}, //单元格点击事件
                RowClick: function() {}, //行点击事件
                CheckboxClick: function() {}, //复选框点击事件
                HeadCheckBoxClick: function() {}, //头部的复选框点击事件
                RequestAfterRender: function() {}
            }
            return setting;
        }
        //创建datagrid的dom树结构
        function CreateDomStruct() {
            var divThead = $("<div></div>");
            var divTbody = $("<div></div>");
            var divTfoot = $("<div></div>").prop("id", $(Dom).prop("id") + "divtfoot");
            $(Dom).append(divThead).append(divTbody).append(divTfoot);
            Setting.divThead = {};
            Setting.divThead.Dom = divThead;
            Setting.divTbody = {};
            Setting.divTbody.Dom = divTbody;
            Setting.divTfoot = {};
            Setting.divTfoot.Dom = divTfoot;

            var divTheadTable = $("<table></table>").addClass("table ");
            divTheadTable.append(CreateHead());
            divThead.append(divTheadTable);

            CreatPagination();
        }

        function RequestAfterRender(){
            Setting.RequestAfterRender()
        }

        function CreatPagination() {
            if (Setting.ShowPagination) {
                Setting.Pagination.onChangePage = OnPaginationChange;
                Setting.divTfoot.Dom.bs_pagination(Setting.Pagination);
                Setting.Pagination = Setting.divTfoot.Dom.bs_pagination('getAllOptions');
            }
        }

        function OnPaginationChange() {
            RequestData();
        }

        function RefreshDataViewHead() {
            Setting.divThead.Dom.find('table').empty().append(CreateHead());
        }
        //数据表格刷新
        function RefreshDataView() {
            var divTbodyTable = CreateBody().addClass("table");
            Setting.divTbody.Dom.empty().append(divTbodyTable);
            //SetWidth(Setting.divThead.Dom.find("table"), divTbodyTable);
            if (Setting.Height) {
                if ($(Dom).height() > Setting.Height) {
                    Setting.divThead.Dom.find("table").css("margin-bottom", "0px");
                    Setting.divTbody.Dom.height(Setting.Height - Setting.divThead.Dom.height() - Setting.divTfoot.Dom.height()).css("overflow-y", "auto");
                    Setting.divThead.Dom.css("padding-right", "17px");
                    Setting.divTfoot.Dom.css("padding-right", "17px");
                }
            }
        }

        function SetWidth(head, body) {
            var bodytd = $($(body).children("tbody").children("tr")[0]).children("td")
            var headtd = $($(head).children("tbody").children("tr")[0]).children("td")
            for (var i = 0; i < bodytd.length; i++) {
                $(headtd[i]).width($(bodytd[i]).width());
            }
        }

        function CreateHead() {
            var theadModel = new Object();
            var tr = $("<tr></tr>");
            if (Setting.ShowRowNum) {
                var tdnum = $("<td></td>")
                tr.append(tdnum);
            }
            if (Setting.ShowCheckBox) {
                var tdcheckbox = $("<td></td>").css('width', '20px');
                var input = $("<input type='checkbox'/>").addClass("DatagridHideCheckbox")
                tdcheckbox.append(input)
                input.CheckBox({
                    Click: HeadCheckBoxClick
                });
                tr.append(tdcheckbox);
            }
            var td;
            for (var i = 0; i < Setting.Columns.length; i++) {
                td = $("<td></td>").addClass(Setting.Columns[i].theadStyle).attr("name", Setting.Columns[i].field)
                if (Setting.Columns[i].headTextAlign == undefined) {
                    td.css('text-align', Setting.Columns[i].headTextAlign);
                }
                if (Setting.Columns[i].Sortable) { //可以进行排序
                    var span = $("<span style='display:inline-block'></span>").html(Setting.Columns[i].title).data("Column", Setting.Columns[i]).click(function(event) {
                        var column = $(this).data("Column");
                        tr.find(".SortColumns").remove();
                        if (column.Order == undefined) {
                            column.Order = "ASC"
                        }

                        function SortASC(parm1, parm2) {
                            if (typeof column.SortASC == "function") {
                                return column.SortASC(parm1, parm2, column);
                            } else if (typeof column.Sort == "function") {
                                return column.Sort(parm1, parm2, column);
                            } else {
                                if (parm1 > parm2)
                                    return 1
                                else if (parm1 == parm2)
                                    return 0
                                else
                                    return -1
                            }
                        }

                        function SortDESC(parm1, parm2) {
                            if (typeof column.SortDESC == "function") {
                                return column.SortDESC(parm1, parm2, column);
                            } else if (typeof column.Sort == "function") {
                                return -column.Sort(parm1, parm2, column);
                            } else {
                                if (parm1 > parm2)
                                    return -1
                                else if (parm1 == parm2)
                                    return 0
                                else
                                    return 1
                            }
                        }

                        var triangle = $("<span style='display:inline-block'></span>");
                        $(this).append(triangle);
                        if (column.Order == "ASC") {
                            triangle.addClass("triangle-up").addClass("SortColumns")
                            column.Order = "DESC"
                            Setting.Rows.sort(SortASC);
                            RefreshDataView();
                        } else {
                            triangle.addClass("triangle-down").addClass("SortColumns")
                            column.Order = "ASC"
                            Setting.Rows.sort(SortDESC);
                            RefreshDataView();
                        }
                    });
                    td.append(span);
                } else {
                    td.html(Setting.Columns[i].title) //不可以进行排序
                }
                tr.append(td);
            }
            return tr;
        }

        function CreateBody() {
            var table = $("<table></table>");
            for (var i = 0; i < Setting.Rows.length; i++) {
                table.append(CreateRow(Setting.Rows[i]));
            }
            return table;
        }

        function CreateRow(data, columns) {
            var tr = $("<tr></tr>").data("RowData", data);
            if (Setting.ShowRowNum) {
                tr.append($("<td></td>"))
            }
            if (Setting.ShowCheckBox) {
                var tdcheckbox = $("<td></td>").css('width', '20px').data("RowDom", tr);
                var input = $("<input type='checkbox' />").addClass("DatagridHideCheckbox").data("value", data).data("RowDom", tr);
                tdcheckbox.append(input);
                input.CheckBox({
                    Click: CellCheckBoxClick
                });
                tr.append(tdcheckbox)
            }
            for (var i = 0; i < Setting.Columns.length; i++) {
                tr.append(CreateCell(data, Setting.Columns[i]).data("RowDom", tr));
            }
            tr.click(function(event) {
                RowClick(this, event, data)
            })
            return tr;
        }

        function CreateCell(data, column) {
            var td = $("<td></td>").addClass(column.theadStyle).attr("name", column.field);
            var field = data[column.field];
            if (typeof column.Formatter == "function") {
                var tdinfo = column.Formatter(data[column.field], data, column);
                if (typeof tdinfo == "string") {
                    td.html(tdinfo);
                } else if (typeof tdinfo == "object") {
                    td.append(tdinfo);
                }
                //td.html()
            } else {
                td.html(field)
            }
            td.click(function(event) {
                CellClick(this, event, data, column);
            })
            return td;
        }

        function HeadCheckBoxClick(sender, event, input) {
            var value = $(input).CheckBox("GetValue");
            Setting.divTbody.Dom.find("td .DatagridHideCheckbox").each(function() {
                $(this).CheckBox("SetValue", value);
            })
            Setting.HeadCheckBoxClick();
            event.stopPropagation();
        }

        function CellCheckBoxClick(sender, event, input) {
            Setting.CheckboxClick();
            event.stopPropagation();
        }

        function RowClick(send, event, data) {
            Setting.RowClick(send, event, data);
        }

        function CellClick(send, event, data, column) {
            Setting.CellClick(send, event, data, column);
        }

        function GetAllCheckedCheckbox() {
            return Setting.divTbody.Dom.find("td .DatagridHideCheckbox:checked")
        }

        function GetAllCheckbox() {
            return Setting.divTbody.Dom.find("td .DatagridHideCheckbox");
        }

        function GetHeadCheckbox() {
            return Setting.divThead.Dom.find("td .DatagridHideCheckbox");
        }

        function EditRequestParm(parm) {
            $.extend(Setting.RequestParm, parm);
        }

        function BindFunction() {
            Setting.GetAllCheckedCheckbox = GetAllCheckedCheckbox;
            Setting.GetAllCheckbox = GetAllCheckbox;
            Setting.GetHeadCheckbox = GetHeadCheckbox;
            Setting.RequestData = RequestData;
            Setting.EditRequestParm = EditRequestParm;
            Setting.InsertTr = InsertTr;
            Setting.Destroy = Destroy;
        }

        function InsertTr(value) {
            var tr = CreateRow(value, Setting.Columns);
            Setting.divTbody.Dom.children('table').prepend(tr);
            return tr;
        }

        function Destroy() {
            $(Dom).empty();
            $(Dom).data("DataGrid", null);
        }
    }
})(jQuery)