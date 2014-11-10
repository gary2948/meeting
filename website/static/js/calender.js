jQuery(document).ready(function($) {
    $("#meetinglink").addClass('active');

    $('#calendar').fullCalendar({
        header: {
            left: 'prev,next today',
            center: 'title',
            right: ''
        },
        dayClick: function(date, jsEvent, view) {
            $("#meetingdate").val(date.format())
            CommitInfoCallBack = function() {
                $('#calendar').fullCalendar('refetchEvents')
            }
            CreateMeetingDlg();
        },
        eventClick: function(calEvent, jsEvent, view) {
            EventClick(calEvent, jsEvent, view);
        },
        events: function(start, end, timezone, callback) {
            var view = $('#calendar').fullCalendar('getView');
            $.ajax({
                url: "/dataserver",
                type: "post",
                data: {
                    ActionType: "TimeSpanMeetingList",
                    StartTime: view.start.unix(),
                    EndTime: view.end.unix(),
                },
                dataType: "json",
                success: function(data) {

                    var events = [];
                    if (data.RequestResult) {
                        if (!(data.Data == null || data.Data == undefined)) {
                            for (var i = 0; i < data.Data.length; i++) {
                                data.Data[i].title = data.Data[i].Lc_topic
                                data.Data[i].start = data.Data[i].Lc_planTime
                                events.push(data.Data[i])
                            }
                        }
                    } else {
                        if (data.Data == 0)
                            location.href = "/";
                    }
                    callback(events);
                },
                error: function(data) {},
            })
        }
    });


    $("#tooltip").remove().appendTo('body').hide().click(function(event) {
        event.stopPropagation();
    });

    $("body").click(function(event) {
        if ($("#tooltip").hasClass('show')) {
            $("#tooltip").removeClass('show').hide()
        }
    });

    $("#deleteMeetingtool").click(function(event) {
        var data = $("#tooltip").data('value')
        $.ajax({
            url: '/dataserver',
            type: 'POST',
            dataType: 'json',
            data: {
                ActionType: 'DeleteMeeting',
                meetingID: data.Id
            },
            success: function(value) {
                if(!value.ResultModel && value.Data==0)
                    location.href="/"
                $("#tooltip").removeClass('show').hide()
                $('#calendar').fullCalendar('refetchEvents')
            },
            error: function() {
                $("#tooltip").removeClass('show').hide()
            }
        })

    });

    $("#editMeetingtool").click(function(event) {
        var data = $("#tooltip").data('value')
        window.location.href = "/editmeeting/" + data.Id
    });

})


function EventClick(calEvent, jsEvent, view) {

    var tip = $("#tooltip");
    var tipValue = tip.data("value")
    if (tipValue != null && tipValue != undefined && tipValue.Id == calEvent.Id) {
        if ($("#tooltip").hasClass('show')) {
            $("#tooltip").hide().removeClass('show')
            jsEvent.stopPropagation();
            return;
        }
    }
    tip.data("value", calEvent)
    tip.children('.popover-title').empty().append(function() {
        var title = $("<div></div>")
        title.css({
            "text-align": 'center',
            color: 'red'
        });
        title.html(calEvent.Lc_topic)
        return title;
    })
    tip.children('.popover-content').empty().append(function() {
        var content = $("<div></div>").css({
            width: '250px'
        });
        content.html(calEvent.Lc_schema);
        return content;
    })

    var width = tip.outerWidth(); //框宽度
    var height = tip.outerHeight(); //框高度
    var clientX = jsEvent.clientX; //屏幕x
    var clientY = jsEvent.clientY; //屏幕y
    var offsetX = jsEvent.offsetX;
    var offsetY = jsEvent.offsetY;
    var windowX = $(window).width(); //工作区w
    var windowY = $(window).height(); //工作区y



    var positionX, positionY, argX, argY;

    positionY = jsEvent.pageY - offsetY - height;
    if (clientX - width / 2 < 0) {
        positionX = 0;
        argX = clientX - positionX
    } else if (clientX + width / 2 > windowX) {
        positionX = windowX - clientX
    } else {
        positionX = clientX - width / 2;
    }
    tip.css({
        top: positionY,
        left: positionX
    });
    tip.find('arrow').css({
        left: argX
    });
    tip.addClass('show').show();
    jsEvent.stopPropagation();
}