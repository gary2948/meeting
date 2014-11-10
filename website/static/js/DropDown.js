/// <reference path="static/js/jquery-2.1.1.js" />
(function ($) {
    $.fn.LcDropDown = function (option, value) {
        if (typeof option === "string") {
            if (value == undefined) {
                return $(this[0]).data("LcDropDown")[option]()
            }
            return $(this[0]).data("LcDropDown")[option](values)
        } else if (typeof option === "object") {
            return $(this).each(function () {
                LcDropDown($(this), option)
            })
        }
    }
    function LcDropDown(Dom, option) {
        var defaultsetting = defaultSetting();
        var Setting = $.extend({}, defaultsetting, option);
        var UserSetting = option;

        $(Dom).append(CreateItems(Setting.ItemList, Setting.Select));
        BindFunction();
        $(Dom).data("LcDropDown", Setting);


        function defaultSetting() {
            var set = new Object();
            set.ItemList = [];
            set.Select = "";
            set.Enable=true;
            set.ItemFormat = function (value) { return value;}
            return set;
        }

        function CreateItems(ItemList, value, ItemClickFunction) {
            var div = $("<div></div>");//.addClass("btn-group");
            var divButton = $("<button></button>").addClass("btn btn-danger dropdown-toggle").attr("data-toggle", "dropdown");
            var divButtonValue = $("<span></span>").html(Setting.ItemFormat(value)).addClass("pickValue");
            var divButtonCaret = $("<span></span>").addClass("caret").html("");
            var divUl = $("<ul></ul>").addClass("dropdown-menu").css({ "height": "300px", "overflow-y": "auto" });
            divUl.attr("role", "menu");
            var li;

            for (var i = 0; i < ItemList.length ; i++) {
                li = $("<li></li>").append($("<a href='javascript:void(0)'></a>").html( Setting.ItemFormat( ItemList[i]))).data("ItemValue", ItemList[i]).click(function (event) {
                    var iv = $(this).data("ItemValue");
                    divButtonValue.html(Setting.ItemFormat(iv)).data("value", iv);
                    Setting.Select = iv;
                });
                if (ItemList[i].name == value) {
                    li.addClass("DefaultSelect");
                }
                divUl.append(li);
            }
            divButton.append(divButtonValue).append(divButtonCaret);
            div.append(divButton).append(divUl);
            return div;
        }

        function GetValue() {
            return Setting.Select;
        }

        function BindFunction() {
            Setting.GetValue = GetValue;
        }
    }
})(jQuery)