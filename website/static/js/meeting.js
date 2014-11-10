/// <reference path="jquery-1.7.1.js" />
jQuery(document).ready(function($) {
    $("#meetinglink").addClass('active');
    $("#container>div").hide();

    $("#startMeeting").click(function(event) {
        $(this).parents("ul").children('.active').removeClass('active')
        $(this).parent().addClass('active')
        $("#container>div").hide();
        $("#startContainer").show();
    });

    $("#unStartMeeting").click(function(event) {
        $(this).parents("ul").children('.active').removeClass('active')
        $(this).parent().addClass('active')
        $("#container>div").hide();
        $("#unStartContainer").show();
    });

    $("#endMeeting").click(function(event) {
        $(this).parents("ul").children('.active').removeClass('active')
        $(this).parent().addClass('active')
        $("#container>div").hide();
        $("#endContainer").show();
    });

    $("#allMeeting").click(function(event) {
        $(this).parents("ul").children('.active').removeClass('active')
        $(this).parent().addClass('active')
        $("#container>div").hide();
        $("#allContainer").show();
    });
    CreateGrid();

    switch (window.location.hash) {
        case "#startMeeting":
            $("#startMeeting").click();
            break;
        case "#unStartMeeting":
            $("#unStartMeeting").click();
            break;
        case "#endMeeting":
            $("#endMeeting").click();
            break;
        case "#allMeeting":
            $("#allMeeting").click();
            break;
        default:
            $("#startMeeting").click();
    }
});


function CreateGrid() {
    $("#startContainer").DataGrid({
        Url: '/dataserver',
        RequestParm: {
            ActionType: "StartMeetingList"
        },
        RepostDataFilter: RepostDataFilter,
        Columns: [{
            field: "id",
            title: "会议主题",
            theadStyle: "col-md-4",
            Formatter: MeetingTopicFunc
        }, {
            field: "none",
            title: "状态",
            theadStyle: "col-md-2",
            Formatter: function() {
                return $("<div></div>").html("正在进行")
            }
        }, {
            field: "Lc_userName",
            title: "创建人",
            theadStyle: "col-md-2"
        }, {
            field: "Lc_beginTime",
            title: "开始时间",
            theadStyle: "col-md-2",
            Formatter: TimeFunc
        }, {
            field: "none",
            title: "操作",
            theadStyle: "col-md-2",
            Formatter: EnterMeetingFunc
        }],
    })

    $("#unStartContainer").DataGrid({
        Url: '/dataserver',
        RequestParm: {
            ActionType: "UnStartMeetingList"
        },
        RepostDataFilter: RepostDataFilter,
        Columns: [{
            field: "id",
            title: "会议主题",
            theadStyle: "col-md-4",
            Formatter: MeetingTopicFunc
        }, {
            field: "none",
            title: "状态",
            theadStyle: "col-md-2",
            Formatter: function() {
                return $("<div></div>").html("还未开始");
            }
        }, {
            field: "Lc_userName",
            title: "创建人",
            theadStyle: "col-md-2"
        }, {
            field: "Lc_planTime",
            title: "开始时间",
            theadStyle: "col-md-2",
            Formatter: TimeFunc
        }, {
            field: "none",
            title: "操作",
            theadStyle: "col-md-2",
            Formatter: StartMeetingFunc
        }],
    })


    $("#endContainer").DataGrid({
        Url: '/dataserver',
        RequestParm: {
            ActionType: "EndMeetingList"
        },
        RepostDataFilter: RepostDataFilter,
        Columns: [{
            field: "id",
            title: "会议主题",
            theadStyle: "col-md-4",
            headTextAlign: "center",
            Formatter: MeetingTopicFunc
        }, {
            field: "none",
            title: "状态",
            theadStyle: "col-md-2",
            Formatter: function() {
                return $("<div></div>").html("还未开始");
            }
        }, {
            field: "Lc_userName",
            title: "创建人",
            theadStyle: "col-md-2"
        }, {
            field: "Lc_endTime",
            title: "结束时间",
            theadStyle: "col-md-2",
            Formatter: TimeFunc
        }, {
            field: "none",
            title: "操作",
            theadStyle: "col-md-2",
            Formatter: ViewMeetingFunc
        }],
    })

    $("#allContainer").DataGrid({
        Url: '/dataserver',
        RequestParm: {
            ActionType: "AllMeetingList"
        },
        RepostDataFilter: RepostDataFilter,
        Columns: [{
            field: "id",
            title: "会议主题",
            theadStyle: "col-md-4",
            Formatter: MeetingTopicFunc
        }, {
            field: "none",
            title: "状态",
            theadStyle: "col-md-2",
            Formatter: function(value, row, index, column) {
                var div = $("<div></div>");
                if (row.Lc_status == 0) {
                    div.html("还未开始");
                } else if (row.Lc_status == 1) {
                    div.html("正在进行");
                } else if (row.Lc_status == 2) {
                    div.html("已经结束");
                }
                return div;
            }
        }, {
            field: "Lc_userName",
            title: "创建人",
            theadStyle: "col-md-2"
        }, {
            field: "none",
            title: "时间",
            theadStyle: "col-md-2",
            Formatter: AllTimeFunc
        }, {
            field: "none",
            title: "操作",
            theadStyle: "col-md-2",
            Formatter: AllMeetingFunc
        }],
    })

    function RepostDataFilter(data) {
        if (data.RequestResult) {
            if (data.Data == null || data.Data == undefined) {
                data.total = 0;
                data.rows = [];
                return data;
            } else {
                data.total = data.Data.TotalRow;
                data.rows = data.Data.Rows == null ? [] : data.Data.Rows;
                return data;
            }
        } else {
            if(data.Data==0){
                location.href="/"
            }
        }
    }
}


function MeetingTopicFunc(value, row, index, column) {
    var div = $("<div></div>").addClass('row')
    var md3 = $("<div></div>").addClass('col-md-3')
    md3.append($("<img/>").attr('src', row.img).css('width', '100%'))
    var md9 = $("<div></div>").addClass('col-md-9')
    md9.append($("<div></div>").append($("<label></label>").html(row.Lc_topic)))
    md9.append($("<div></div>").append($("<span></span>").html(row.Lc_schema)))
    md9.append($("<div></div>").append($("<a href='/editmeeting/" + row.Id + "'></a>").html(row.Lc_code)))
    div.append(md3).append(md9)

    return div
};

function TimeFunc(value, row, index, column) {
    return $("<div></div>").html(FormartDateTime(value))
}

function AllTimeFunc(value, row, index, column) {
    var div = $("<div></div>");
    if (row.Lc_status == 0) {
        div.html(FormartDateTime(row["Lc_planTime"]));
    } else if (row.Lc_status == 1) {
        div.html(FormartDateTime(row["Lc_beginTime"]));
    } else if (row.Lc_status == 2) {
        div.html(FormartDateTime(row["Lc_endTime"]));
    }
    return div;
}

function FormartDateTime(value) {
    if (value == undefined || value == null || value == "")
        return "";
    var datetime = new Date(value);
    return datetime.getFullYear() + "/" + (datetime.getMonth() + 1) + '/' + datetime.getDate() + " " + datetime.getHours() + ":" + datetime.getMinutes();
}

function EnterMeetingFunc(value, row, index, column) {
    var enter = $("<a></a>").html("进入会议").click(function(event) {
        alert("enter")
    });
    return $("<div></div>").append(enter)
}

function StartMeetingFunc(value, row, index, column) {
    var start = $("<a></a>").html("进入会议").click(function(event) {
        alert("start")
    });
    return $("<div></div>").append(start)
}

function ViewMeetingFunc(value, row, index, column) {
    var view = $("<a></a>").html("进入会议").click(function(event) {
        alert("view")
    });
    return $("<div></div>").append(view)
}

function AllMeetingFunc(value, row, index, column) {
    if (row.Lc_status == 0) {
        return EnterMeetingFunc(value, row, index, column);
    } else if (row.Lc_status == 1) {
        return StartMeetingFunc(value, row, index, column);
    } else if (row.Lc_status == 2) {
        return ViewMeetingFunc(value, row, index, column);
    }
}