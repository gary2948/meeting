/// <reference path="static/js/jquery-2.1.1.js" />

(function ($) {
    $.fn.DropDownDatePick = function (option, value) {
        if (typeof option === "string") {
            if (value == undefined) {
                return $(this[0]).data("DropDownDatePick")[option]()
            }
            return $(this[0]).data("DropDownDatePick")[option](values)
        } else if (typeof option === "object") {
            return $(this).each(function () {
                DropDownDatePick($(this), option)
            })
        }
    }
    function DropDownDatePick(Dom, value) {

        function DefaultSetting() {
            var option = new Object();
            option.Date = new Date();
            option.DownHeight = 300;
            return option;
        }

        var defaultSetting = DefaultSetting();
        var UserSetting = value;
        var Setting = $.extend({}, defaultSetting, UserSetting);

        DivYear = CreateDropDown(Setting.Date.getFullYear(), "年", CreateYearItemList(), YearItemClick);
        DivMonth = CreateDropDown(Setting.Date.getMonth()+1, "月", CreateMonthItemList(), MonthItemClick);
        DivDay = CreateDropDown(Setting.Date.getDate(), "日", CreateDayItemList(Setting.Date), DayItemClick);
        $(Dom).append(DivYear).append(DivMonth).append(DivDay);
        BindFunction();
        $(Dom).data("DropDownDatePick", Setting);

        function CreateYearItemList(){
            var start = new Date().getFullYear();
            var item=[];
            for(var i=start;i>=start-120;i--){
                item[item.length]=i;
            }
            return item;
        }
        function CreateMonthItemList() {
            var item = [];
            for (var i = 1; i <= 12; i++) {
                item[item.length] = i;
            }
            return item;
        }
        function CreateDayItemList(date) {
            var item = [];
            var days = GetDaysOfMonth(date.getFullYear(), date.getMonth() + 1);
            for (var i = 1; i <= days; i++) {
                item[item.length] = i;
            }
            return item;
        }
        function GetDaysOfMonth(year, month) {
            var date = new Date(year, month - 1);
            date.setMonth(month);
            date.setDate(0);
            return date.getDate();
        }
        function YearItemClick(sender, event) {
            var day = Setting.Date.getDate();
            Setting.Date.setDate(1);
            Setting.Date.setYear($(sender).data("ItemValue"));
            var days = GetDaysOfMonth(Setting.Date.getFullYear(), Setting.Date.getMonth() + 1);
            if (day > days) {
                
            } else {
                Setting.Date.setDate(day)
            }
            DivDay.remove();
            DivDay = CreateDropDown(Setting.Date.getDate(), "日", CreateDayItemList(Setting.Date), DayItemClick);
            $(Dom).append(DivDay);
            
        }
        function MonthItemClick(sender, event) {
            var day = Setting.Date.getDate();
            Setting.Date.setDate(1);
            Setting.Date.setMonth(parseInt($(sender).data("ItemValue"))-1);
            var days = GetDaysOfMonth(Setting.Date.getFullYear(), Setting.Date.getMonth() + 1);
            if (day > days) {
                
            } else {
                Setting.Date.setDate(day)
            }
            DivDay.remove();
            DivDay = CreateDropDown(Setting.Date.getDate(), "日", CreateDayItemList(Setting.Date), DayItemClick);
            $(Dom).append(DivDay);
        }
        function DayItemClick(sender, event) {
            Setting.Date.setDate($(sender).data("ItemValue"));
        }

        function CreateDropDown(value,title,dropdownlist,liclick) {
            var div = $("<div></div>").addClass("btn-group");
            var divButton = $("<button></button>").addClass("btn btn-danger dropdown-toggle").attr("data-toggle", "dropdown");
            var divButtonValue = $("<span></span>").html(value).data("value",value).addClass("pickValue");
            var divButtonTitle = $("<span><span>").html(title).addClass("pickTitle");
            var divButtonCaret = $("<span></span>").addClass("caret").html("");
            var divUl = $("<ul></ul>").addClass("dropdown-menu").css({ "height": "300px", "overflow-y": "auto" });
            divUl.attr("role", "menu");
            var li;
            for (var i = 0; i < dropdownlist.length; i++) {
                li = $("<li></li>").append($("<a href='javascript:void(0)'></a>").html(dropdownlist[i])).data("ItemValue", dropdownlist[i]).click(function (event) {
                    var iv=$(this).data("ItemValue");
                    divButtonValue.html(iv);
                    divButtonValue.data("value", iv);
                    liclick(this, event);
                });
                divUl.append(li);
            }
            divButton.append(divButtonValue).append(divButtonTitle).append(divButtonCaret);
            div.append(divButton).append(divUl);
            return div;
        }

        function BindFunction() {
            Setting.GetDate = GetDate;
            Setting.Destroy=Destroy;
        }
        function GetDate() {
            return Setting.Date;
        }

        function Destroy(){
            $(Dom).remove();
            $(Dom).removeData('DropDownDatePick')
        }
    }
})(jQuery)