/**
 * @fileOverview bs_grid is a jQuery datagrid plugin based on Twitter Bootstrap.
 *               <p>License MIT
 *               <br />Copyright Christos Pontikis <a href="http://www.pontikis.net">http://www.pontikis.net</a>
 *               <br />Project page <a href="http://www.pontikis.net/labs/bs_grid/">http://www.pontikis.net/labs/bs_grid/</a>
 * @version 0.9.1 (09 May 2014)
 * @author Christos Pontikis http://www.pontikis.net
 * @requires jquery >= 1.8, twitter bootstrap >= 2, bs_pagination plugin, jQuery UI sortable (optional), jui_filter_rules plugin >= 1.0.4 (optional)
 */

/**
 * See <a href="http://jquery.com">http://jquery.com</a>.
 * @name $
 * @class
 * See the jQuery Library  (<a href="http://jquery.com">http://jquery.com</a>) for full details.  This just
 * documents the function and classes that are added to jQuery by this plug-in.
 */

/**
 * See <a href="http://jquery.com">http://jquery.com</a>
 * @name fn
 * @class
 * See the jQuery Library  (<a href="http://jquery.com">http://jquery.com</a>) for full details.  This just
 * documents the function and classes that are added to jQuery by this plug-in.
 * @memberOf $
 */

/**
 * Pseudo-Namespace containing bs_grid private methods (for documentation purposes)
 * @name _private_methods
 * @namespace
 */

"use strict";
(function($) {

	var pluginName = "bs_grid",
		pluginGivenOptions = "bs_grid_given_options",
		pluginStatus = "bs_grid_status";

	// public methods
	var methods = {

		init: function(options) {

			var elem = this;

			return this.each(function() {

				/**
				 * store given options on first launch (in new object - no reference)
				 */
				if (typeof elem.data(pluginGivenOptions) === "undefined") {
					elem.data(pluginGivenOptions, $.extend(true, {}, options));
				}

				/**
				 * settings and defaults
				 * settings modification will affect elem.data(pluginName) and vice versa
				 */
				var settings = elem.data(pluginName);
				if (typeof settings === "undefined") {
					var bootstrap_version = "3";
					if (options.hasOwnProperty("bootstrap_version") && options["bootstrap_version"] == "2") {
						bootstrap_version = "2";
					}
					var defaults = methods.getDefaults.call(elem, bootstrap_version);
					settings = $.extend({}, defaults, options);
				} else {
					settings = $.extend({}, settings, options);
				}
				elem.data(pluginName, settings);

				// initialize plugin status
			if (typeof elem.data(pluginStatus) === "undefined") {
					elem.data(pluginStatus, {});
					//elem.data(pluginStatus)["selected_ids"] = [];
					//elem.data(pluginStatus)["filter_rules"] = [];
				} else {
					//if (!settings.row_primary_key) {
					//	elem.data(pluginStatus)["selected_ids"] = [];
					//} else {
					//	switch (settings.rowSelectionMode) {
					//		case "single":
					//			if (elem.data(pluginStatus)["selected_ids"].length > 1) {
					//				elem.data(pluginStatus)["selected_ids"] = [];
					//			}
					//			break;
					//		case false:
					//			elem.data(pluginStatus)["selected_ids"] = [];
					//			break;
					//	}
					//}
				}

				var container_id = elem.attr("id");

				// apply container style
				elem.removeClass().addClass(settings.containerClass);

				// bind events
				elem.unbind("onCellClick").bind("onCellClick", settings.onCellClick);
				elem.unbind("onRowClick").bind("onRowClick", settings.onRowClick);
				elem.unbind("onDatagridError").bind("onDatagridError", settings.onDatagridError);
				elem.unbind("onDebug").bind("onDebug", settings.onDebug);
				elem.unbind("onDisplay").bind("onDisplay", settings.onDisplay);

				// initialize plugin html
				var tools_id = create_id(settings.tools_id_prefix, container_id),
					columns_list_id = create_id(settings.columns_list_id_prefix, container_id),
					default_columns_list = "",
					sorting_list_id = create_id(settings.sorting_list_id_prefix, container_id),
					default_sorting_list = "",
					sorting_radio_name = create_id(settings.sorting_radio_name_prefix, container_id) + "_",
					startPos, newPos,
					selected_rows_id = create_id(settings.selected_rows_id_prefix, container_id),
					selection_list_id = create_id(settings.selection_list_id_prefix, container_id),
					table_container_id = create_id(settings.table_container_id_prefix, container_id),
					table_id = create_id(settings.table_id_prefix, container_id),
					no_results_id = create_id(settings.no_results_id_prefix, container_id),
					filter_toggle_id = create_id(settings.filter_toggle_id_prefix, container_id),
					custom_html1_id = create_id(settings.custom_html1_id_prefix, container_id),
					custom_html2_id = create_id(settings.custom_html2_id_prefix, container_id),
					pagination_id = create_id(settings.pagination_id_prefix, container_id),
					filter_container_id = create_id(settings.filter_container_id_prefix, container_id),
					filter_rules_id = create_id(settings.filter_rules_id_prefix, container_id),
					filter_tools_id = create_id(settings.filter_tools_id_prefix, container_id),
					elem_html = "",
					tools_html = "";

			    // create basic html structure ---------------------------------
                var elem_tools=$("<div></div>").attr("id",tools_id).addClass(settings.toolsClass)
                var elem_tableContain=$("<div></div>").attr("id",table_container_id).addClass(settings.dataTableContainerClass)
                var elem_table=$("<table></table>").attr("id",table_id).addClass(settings.dataTableClass)
                var elem_pagination=$("<div></div>").attr("id",pagination_id)
                elem.append(elem_tools).append(elem_tableContain.append(elem_table)).append(elem_pagination)
                if (settings.showPagination) {

                } else {
                    elem_pagination.hide()
                }

				// initialize grid ---------------------------------------------
				var grid_init = methods.displayGrid.call(elem, false);

				// PAGINATION --------------------------------------------------
				$.when(grid_init).then(function(data, textStatus, jqXHR) {

					var total_rows = data["total_rows"];

					var pagination_options = settings.paginationOptions,
						bs_grid_pagination_options = {
							// defined by bs_grid
							currentPage: settings.pageNum,
							rowsPerPage: settings.rowsPerPage,
							maxRowsPerPage: settings.maxRowsPerPage,
							totalPages: Math.ceil(total_rows / settings.rowsPerPage),
							totalRows: total_rows,
							bootstrap_version: settings.bootstrap_version,

							onChangePage: function(event, params) {
								settings.pageNum = params.currentPage;
								settings.rowsPerPage = params.rowsPerPage;
								methods.displayGrid.call(elem, false);
							}
						};
					$.extend(pagination_options, bs_grid_pagination_options);
					elem_pagination.bs_pagination(pagination_options);

					// custom html ---------------------------------------------
					// (not an event, but page renders better displaying custom html after grid rendering)
					if (settings.customHTMLelementID1) {
						$("#" + custom_html1_id).html($("#" + settings.customHTMLelementID1).html());
					}

					if (settings.customHTMLelementID2) {
						$("#" + custom_html2_id).html($("#" + settings.customHTMLelementID2).html());
					}

				});

			});
		},

		getVersion: function() {
			return "0.9.1";
		},

		getDefaults: function(bootstrap_version) {
			var default_settings = {
				pageNum: 1,
				rowsPerPage: 10,
				maxRowsPerPage: 100,
				row_primary_key: "",
				rowSelectionMode: "single", // "multiple", "single", false
			    showPagination:true,
				columns: [],

				sorting: [],

				paginationOptions: {//分页插件的配置变量
					containerClass: "well pagination-container",
					visiblePageLinks: 5,
					showGoToPage: true,
					showRowsPerPage: true,
					showRowsInfo: true,
					showRowsDefaultInfo: true,
					disableTextSelectionInNavPane: true
				}, // "currentPage", "rowsPerPage", "maxRowsPerPage", "totalPages", "totalRows", "bootstrap_version", "onChangePage" will be ignored

				showRowNumbers: false,


				/* STYLES ----------------------------------------------------*/
				bootstrap_version: "3",

				// bs 3
				containerClass: "grid_container",
				noResultsClass: "alert alert-warning no-records-found",

				toolsClass: "tools",

				dataTableContainerClass: "table-responsive",
				dataTableClass: "table",
				commonThClass: "th-common",
				selectedTrClass: "warning",

				tools_id_prefix: "tools_",
				columns_list_id_prefix: "columns_list_",

				selected_rows_id_prefix: "selected_rows_",
				selection_list_id_prefix: "selection_list_",
				filter_toggle_id_prefix: "filter_toggle_",

				table_container_id_prefix: "tbl_container_",
				table_id_prefix: "tbl_",

				no_results_id_prefix: "no_res_",

				custom_html1_id_prefix: "custom1_",
				custom_html2_id_prefix: "custom2_",

				pagination_id_prefix: "pag_",
				filter_container_id_prefix: "flt_container_",
				filter_rules_id_prefix: "flt_rules_",
				filter_tools_id_prefix: "flt_tools_",

				// misc
				debug_mode: "no",

				// events
				onCellClick: function() {},
				onRowClick: function() {},
				onDatagridError: function() {},
				onDebug: function() {},
				onDisplay: function() {},
				Filter:function(data){return data}
			};

			return default_settings;
		},

		getOption: function(opt) {
			var elem = this;
			return elem.data(pluginName)[opt];
		},

		getAllOptions: function() {
			var elem = this;
			return elem.data(pluginName);
		},

		destroy: function() {
			var elem = this,
				container_id = elem.attr("id"),
				pagination_container_id = create_id(methods.getOption.call(elem, "pagination_id_prefix"), container_id),
				filter_rules_id = create_id(methods.getOption.call(elem, "filter_rules_id_prefix"), container_id);

			$("#" + pagination_container_id).removeData();
			$("#" + filter_rules_id).removeData();
			elem.removeData();
		},

		selectedRows: function(action, id) {
			var elem = this,
				container_id = elem.attr("id"),
				table_id = create_id(methods.getOption.call(elem, "table_id_prefix"), container_id),
				selectedTrClass = methods.getOption.call(elem, "selectedTrClass"),
				selector_table_tr = "#" + table_id + " tbody tr",
				table_tr_prefix = "#" + table_id + "_tr_";

			switch (action) {
				case "get_ids":
					return elem.data(pluginStatus)["selected_ids"];
					break;
				case "clear_all_ids":
					elem.data(pluginStatus)["selected_ids"] = [];
					break;
				case "update_counter":
					var selected_rows_id = create_id(methods.getOption.call(elem, "selected_rows_id_prefix"), container_id);
					$("#" + selected_rows_id).text(elem.data(pluginStatus)["selected_ids"].length);
					break;
				case "selected_index":
					return $.inArray(id, elem.data(pluginStatus)["selected_ids"]);
					break;
				case "add_id":
					elem.data(pluginStatus)["selected_ids"].push(id);
					break;
				case "remove_id":
					elem.data(pluginStatus)["selected_ids"].splice(id, 1);
					break;
				case "mark_selected":
					$(table_tr_prefix + id).addClass(selectedTrClass);
					break;
				case "mark_deselected":
					$(table_tr_prefix + id).removeClass(selectedTrClass);
					break;
				case "mark_page_selected":
					$(selector_table_tr).addClass(selectedTrClass);
					break;
				case "mark_page_deselected":
					$(selector_table_tr).removeClass(selectedTrClass);
					break;
				case "mark_page_inversed":
					$(selector_table_tr).toggleClass(selectedTrClass);
					break;
			}
		},

		setPageColClass: function(col_index, headerClass, dataClass) {
			var elem = this,
				container_id = elem.attr("id"),
				data_table_selector = "#" + create_id(methods.getOption.call(elem, "table_id_prefix"), container_id);

			if (headerClass !== "") {
				$(data_table_selector + " th").eq(col_index).addClass(headerClass);
			}
			if (dataClass !== "") {
				$(data_table_selector + " tr").each(function() {
					$(this).find("td").eq(col_index).addClass(dataClass);
				});
			}
		},

		removePageColClass: function(col_index, headerClass, dataClass) {
			var elem = this,
				container_id = elem.attr("id"),
				data_table_selector = "#" + create_id(methods.getOption.call(elem, "table_id_prefix"), container_id);

			$(data_table_selector + " th").eq(col_index).removeClass(headerClass);
			$(data_table_selector + " tr").each(function() {
				$(this).find("td").eq(col_index).removeClass(dataClass);
			});

		},

		displayGrid: function(refresh_pag) {

			var elem = this,
				container_id = elem.attr("id"),
				s = methods.getAllOptions.call(elem),

				table_id = create_id(s.table_id_prefix, container_id),
				elem_table = $("#" + table_id),//获取table容器
				no_results_id = create_id(s.no_results_id_prefix, container_id),
				elem_no_results = $("#" + no_results_id),
				filter_rules_id = create_id(s.filter_rules_id_prefix, container_id),//值flt_rules_demo_grid1
				pagination_id = create_id(s.pagination_id_prefix, container_id),//值pag_demo_grid1
				elem_pagination = $("#" + pagination_id),//获取分页容器
				err_msg;

			// fetch page data and display datagrid
			var res = $.ajax({
				type: "POST",
				url: s.ajaxFetchDataURL,
				data: $.extend(true, methods.getOption.call(elem,"queryParm"), {
					page_num: s.pageNum,
					rows_per_page: s.rowsPerPage,
					debug_mode: s.debug_mode
				}),
				dataType: "json",
				success: function(data) {
					var server_error, filter_error, row_primary_key, total_rows, page_data, page_data_len, v,
						columns = s.columns,//获取table的列集合
						col_len = columns.length,//获取总共有多少列
						column, c;

					data=s.Filter.call(this,data);

					/*server_error = data["error"];
					if (server_error != null) {
						err_msg = "ERROR: " + server_error;
						elem.html('<span style="color: red;">' + err_msg + '</span>');
						elem.triggerHandler("onDatagridError", {
							err_code: "server_error",
							err_description: server_error
						});
						$.error(err_msg);
					}*/

					total_rows = data["total_rows"];//获取服务器端总共有多少条数据
					page_data = data["page_data"];//获取服务器端传输过来的行数据集合
					page_data_len = page_data.length;//获取服务器端总共传输过来多少行数据

					elem.data(pluginStatus)["total_rows"] = total_rows;

					row_primary_key = s.row_primary_key;

					// create data table
					var pageNum = parseInt(s.pageNum),
						rowsPerPage = parseInt(s.rowsPerPage),
						row_id_html, i, row, tbl_html, row_index,
						offset = ((pageNum - 1) * rowsPerPage);


					var dg_thead = $("<thead></thead");
					var dg_thead_tr=$("<tr></tr>");
					if (s.showRowNumbers){
						var dg_thead_td_index=$("<th></th>")
						dg_thead_tr.append(dg_thead_td_index)
					}
					for(i in s.columns){
						if (column_is_visible(s.columns[i])) {
							var dg_thead_td=$("<th></th>")
							dg_thead_td.html(s.columns[i].header).addClass(s.columns[i].theadStyle)
							dg_thead_tr.append(dg_thead_td)
						}
					}
                    dg_thead.append(dg_thead_tr)

					var dg_tbody=$("<tbody></tbody")
					var createTr=function(index, data,columns){
						var tr=$("<tr></tr>")
						if(s.showRowNumbers){
							var td_index=$("<td></td>")
							td_index.html(offset + parseInt(row) + 1)
							tr.append(td_index)
						}
						var createTd=function(row,column){
							if(column_is_visible(column)){
							    var td = $("<td></td>")
							    if (column.formatter !== null) {
							        var value = column.formatter.call(td, row[column.field], row, index, column)
							    	if (typeof value ==="string"){
							    		td.html(value)
							    	}else if(typeof value==="object"){
							    	    if (value instanceof jQuery) {
							    			td.append(value)
							    		}
							    	}
							    } else {

							        td.html(row[column.field])
							    }
								td.click(function(event) {
									s.onCellClick(event,$(this),index,row,column)
								});
								return td;
							}else{
								return;
							}

						}
						for(i in columns){
							tr.append(createTd(data,columns[i]))
						}
						tr.click(function(event) {
							s.onRowClick(event,$(this),index,data,columns)
						});
						return tr
					}
					for(row in page_data){
						dg_tbody.append(createTr(row, page_data[row],s.columns))
					}
					elem_table.empty().append(dg_thead).append(dg_tbody)


                    /*

					// 是否刷新分页导航条
					if (refresh_pag) {
						elem_pagination.bs_pagination({
							currentPage: s.pageNum,
							totalPages: Math.ceil(total_rows / s.rowsPerPage),
							totalRows: total_rows
						});
					}

					// 没有数据时进行的界面显示
					if (total_rows == 0) {
						elem_pagination.hide();
						elem_no_results.show();
					} else {
						elem_pagination.show();
						elem_no_results.hide();
					}

					// apply given styles ------------------------------------------
					var col_index = s.showRowNumbers ? 1 : 0,
						headerClass = "",
						dataClass = "";
					for (i in s.columns) {
						if (column_is_visible(s.columns[i])) {
							headerClass = "", dataClass = "";
							if (columns[i].hasOwnProperty("headerClass")) {
								headerClass = columns[i]["headerClass"];
							}
							if (columns[i].hasOwnProperty("dataClass")) {
								dataClass = columns[i]["dataClass"];
							}
							methods.setPageColClass.call(elem, col_index, headerClass, dataClass);
							col_index++;
						}
					}

					// apply row selections ----------------------------------------
					if (s.row_primary_key && elem.data(pluginStatus)["selected_ids"].length > 0) {

						if (s.rowSelectionMode == "single" || s.rowSelectionMode == "multiple") {
							var row_prefix_len = (table_id + "_tr_").length,
								row_id, idx;
							$("#" + table_id + " tbody tr").each(function() {
								row_id = parseInt($(this).attr("id").substr(row_prefix_len));
								idx = methods.selectedRows.call(elem, "selected_index", row_id);
								if (idx > -1) {
									methods.selectedRows.call(elem, "mark_selected", row_id);
								}
							});
						}
					}

					// update selected rows counter
					methods.selectedRows.call(elem, "update_counter");

					// trigger event onDisplay
					elem.triggerHandler("onDisplay");*/

				}
			});

			return res;

		}

	};

	var create_id = function(prefix, plugin_container_id) {
		return prefix + plugin_container_id;
	};

	var column_is_visible = function(column) {
		var visible = "visible";
		return !column.hasOwnProperty(visible) || (column.hasOwnProperty(visible) && column[visible] == "yes");
	};

	var set_column_visible = function(column, status) {
		var visible = "visible";
		if (status) {
			if (column.hasOwnProperty(visible)) {
				delete column[visible];
			}
		} else {
			column[visible] = "no";
		}
	};

	$.fn.bs_grid = function(method) {

		if (this.size() != 1) {
			var err_msg = "You must use this plugin (" + pluginName + ") with a unique element (at once)";
			this.html('<span style="color: red;">' + 'ERROR: ' + err_msg + '</span>');
			$.error(err_msg);
		}

		// Method calling logic
		if (methods[method]) {
			return methods[method].apply(this, Array.prototype.slice.call(arguments, 1));
		} else if (typeof method === "object" || !method) {
			return methods.init.apply(this, arguments);
		} else {
			$.error("Method " + method + " does not exist on jQuery." + pluginName);
		}

	};

})(jQuery);