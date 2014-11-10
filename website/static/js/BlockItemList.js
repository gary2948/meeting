/// <reference path="static/js/jquery-2.1.1.js" />
(function ($) {
    $.fn.BlockItemList = function (option,values) {
        if (typeof option === "string") {
            return $(this[0]).data("BlockItemList")[option](values)
        } else if (typeof option === "object") {
            return $(this).each(function () {
                BlockItemList($(this),option)
            })
        }

        function BlockItemList(Dom, userSetting) {
            function DefaultOption() {
                var opt = new Object();
                opt.Items = [];
                opt.ClickFun = null;
                opt.InfoFormat = function () { };
                opt.BtnFormat = function () { };
                return opt;
            }
            var defaultSetting = DefaultOption();
            var setting = $.extend({}, defaultSetting, userSetting);
            $(Dom).data("BlockItemList", setting);

            setting.Dom = Dom;
            setting.defaultSetting = defaultSetting;
            setting.userSetting = userSetting;

            Refresh(Dom, setting.Items);
            BindFunction();


            function CreateItem(itemvalue) {
                var itemblock = $("<div></div>").addClass("btn-group");
                var info = $("<button></button>").addClass("btn disabled ").html(setting.InfoFormat(itemvalue));
                var btn = $("<button></button>").addClass("btn").html(setting.BtnFormat(itemvalue)).click(function (event) {
                    var isrefresh;
                    if (setting.ClickFun == null) {
                        isrefresh=DefaultClickFun($(this), event, itemvalue);
                    } else {
                        isrefresh=setting.ClickFun($(this), event, itemvalue);
                    }
                    if (isrefresh === true) {
                        Refresh();
                    }
                });
                itemblock.append(info).append(btn)
                return itemblock;

            }
            function Refresh() {
                $(Dom).empty();
                if (Array.isArray(setting.Items)) {
                    for (var i = 0; i < setting.Items.length; i++) {
                        $(Dom).append(CreateItem(setting.Items[i]))
                    }
                }
            }
            function DefaultClickFun(dom, event, itemvalue) {
                $(dom).parents(".btn-group").remove();
                setting.Items.splice(setting.Items.indexOf(itemvalue, 0), 1);
            }
            function BindFunction() {
                setting.GetItems = GetItems;
                setting.AddItem = AddItem;
            }
            function GetItems() {
                return setting.Items;
            }
            function AddItem(value) {
                if(value instanceof Array){
                    for(var i=0;i<value.length;i++)
                        setting.Items.push(value[i]);
                }else{
                    setting.Items.push(value);
                }
                Refresh()
            }
        }

    }
})(jQuery)