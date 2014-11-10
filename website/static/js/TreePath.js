/// <reference path="static/js/jquery-2.1.1.js" />
(function ($) {
    $.fn.TreePath = function (option,values) {
        if (typeof option === "string") {
            return $(this[0]).data("TreePath")[option](values)
        } else if (typeof option === "object") {
            return $(this).each(function () {
                TreePath($(this), option)
            })
        }




        function TreePath(Dom, usersetting) {
            var defaultsetting=DefaultSetting();
            var Setting = $.extend({}, defaultsetting, usersetting);
            Setting.Dom = Dom;
            $(Dom).data("TreePath", Setting);

            RefreshView();

            BindFunction();


            function DefaultSetting() {
                var defaultsetting = new Object();
                defaultsetting.ItemList = [];
                defaultsetting.ItemClick = function () { };
                defaultsetting.ItemFormat = function (value) { };
                return defaultsetting;
            }

            function RefreshView() {
                $(Dom).empty();
                for (var i = 0; i < Setting.ItemList.length; i++) {
                    Dom.append(CreateItem(Setting.ItemList[i]).data("index",i));
                }

                var totalWidth = $(Dom).width()-20;
                var item = $(Dom).children();
                var width = 0;
                for (var i = item.length - 1; i >= 0; i--) {
                    width += $(item[i]).outerWidth();
                }
                if (width > totalWidth) {
                    var li = $("<li></li>");
                    var lidiv = $("<div></div>").addClass('btn-group')
                    var divspan = $("<span  data-toggle='dropdown'></span>").addClass('caret')
                    var divul = $("<ul role='menu'></ul>").addClass('dropdown-menu')
                    lidiv.append(divspan).append(divul);
                    li.append(lidiv);
                    $(Dom).prepend(li);
                    totalWidth = totalWidth - li.outerWidth();
                    li.remove();
                    for (var i = item.length - 1; i >= 0; i--) {
                        if (totalWidth - $(item[i]).outerWidth() < 0) {
                            $(item[i]).prevAll().each(function (index, el) {
                                $(this).detach()
                                divul.append(this)
                            });
                            $(item[i]).detach();
                            divul.prepend(item[i])
                            $(Dom).prepend(li);
                            break;
                        } else {
                            totalWidth -= $(item[i]).outerWidth()
                        }
                    }
                }
            }
            function CreateItem(value) {
                var item = $("<li></li>");
                var alink = $("<a href='javascript:void(0)'></a>").html(Setting.ItemFormat(value));
                item.append(alink);
                item.data("value", value).click(function (event) {
                    ItemClickFunc($(this), event, value);
                });
                return item;
            }
            function ItemClickFunc(sender, event, value) {
                Setting.ItemList.splice($(sender).data("index")+1, Setting.ItemList.length);
                RefreshView();
                Setting.ItemClick(sender, event, value);
            }
            function PushItem(value) {
                Setting.ItemList.push(value);
                RefreshView();
            }
            function GetLastItem(){
                return Setting.ItemList[Setting.ItemList.length-1];
            }
            function BindFunction() {
                Setting.PushItem = PushItem;
                Setting.GetLastItem=GetLastItem;
                Setting.Destroy=Destroy;
            }
            function Destroy(){
                $(Dom).empty();
                $(Dom).data("TreePath",null);
            }
        }
    }

})(jQuery)