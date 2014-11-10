/// <reference path="static/js/jquery-2.1.1.js" />
(function ($) {
    $.fn.CheckBox = function (option,value) {
        if (typeof option == "string") {
            return $(this[0]).data("CheckBox")[option](value);
        } else {
            return $(this).each(function () {
                CheckBox(this, option);
            })
        }
    }

    function CheckBox(Dom, option) {

        var defaultSetting = DefaultOption();
        var Setting = $.extend({}, defaultSetting, option);
        Setting.defaultSetting = defaultSetting;
        Setting.userSetting = option;

        Create();
        BindFunction();
        $(Dom).data("CheckBox", Setting);


        function DefaultOption() {
            var option = {
                Click: function () { },
            }
            return option;
        }

        function Create() {
            var outer = $("<div></div>").addClass('mycheckboxouter');
            var label = $("<label></label>").addClass('mycheckbox');
            var parent = $(Dom).parent();
            outer.append(label).append($(Dom).detach());
            parent.append(outer);


            if ($(Dom).prop('checked')) {

            } else {
                label.hide();
            }
            $(Dom).hide();
            Setting.Border = outer;
            Setting.Container = label;
            Setting.Parent = parent;

            outer.click(function (event) {
                Click(this, event);
            })

            $(Dom).unbind('change').change(function () {
                Onchange();
            })
        }

        function Click(sender, event) {
            if ($(Dom).prop("checked")) {
                $(Dom).prop("checked",false)
            } else {
                $(Dom).prop("checked",true)
            }
            $(Dom).change();
            Setting.Click(sender,event,$(sender).find("input"))
        }

        function SetValue(value) {
            if ($(Dom).prop("checked") != value) {
                $(Dom).prop("checked", value);
                $(Dom).change();
            }
        }

        function GetValue() {
            return $(Dom).prop("checked");
        }

        function Onchange() {
            if ($(Dom).prop("checked"))
                Setting.Container.show();
            else
                Setting.Container.hide();
            //Setting.Onchange($(Dom).parent(".mycheckboxouter"), event, $(Dom))
        }

        function BindFunction() {
            Setting.SetValue = SetValue;
            Setting.GetValue = GetValue;
        }
    }

})(jQuery)