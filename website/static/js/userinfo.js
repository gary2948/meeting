$(document).ready(function() {

	$("#editPersonInfo").click(function(event) {
		ControlEditView(true);
	});
	$("#commitPersonInfo").click(function(event) {
		CommitPersonInfoClick()
	}).hide();
	$("#giveUpPersonInfo").click(function(event) {
		ControlEditView(false);
	}).hide();
	ControlEditView(false);
	InitEditValue();
});

function CommitPersonInfoClick() {
	var parm=new Object();
	 $("#UserDetailInfo input,#UserDetailInfo textarea").each(function(index, el) {
	 	parm[$(this).data("name")]=$(this).val();
	 });
	 parm["sex"]=$("#sex").LcDropDown("GetValue");
	 parm["birthday"]=$("#birthday").DropDownDatePick("GetDate").getTime();
	 parm["hometown"]="中国 "+$("#hometown").ProvinceCitySelect("GetCombo"," ")
	 parm["ActionType"]="EditUserInfo";
	 parm["language"]=$("#selectFirstLanguage").LcDropDown("GetValue").value;
	 parm["language2"]=$("#selectSecondLanguage").LcDropDown("GetValue").value;
	 //parm["list"]=[{name:"test1",id:{id1:1,id2:2}},{name:"test1",id:{id1:1,id2:2}}] ;//[{name:"test1",id:1},{name:"test2",id:2}]  [1,2,3]; 

	 $.ajax({
	 	url: '/dataserver',
	 	type: 'POST',
	 	dataType: 'json',
	 	data: parm,
	 	success:function(value){
	 		location.reload();
	 	},
	 	error:function(value){}
	 });
	 
}

function InitEditValue(){
	$("#birthday").DropDownDatePick({Date:new Date($("#birthday").data('value')*1000)})
	var value=$("#hometown").data('value').split(" ");
	$("#hometown").ProvinceCitySelect({Province:value[1],City:value[2]});
	$("#sex").LcDropDown({ItemList:["男","女"],Select:$("#sex").data("value")})
	var value=LanguageList($("#selectFirstLanguage").data("language"));
	$("#selectFirstLanguage").LcDropDown({
		ItemList:LanguageList(),
		Select:value,
		ItemFormat:function(value){ return value.title;}
	})//语言
	value=LanguageList($("#selectSecondLanguage").data("language"));
	$("#selectSecondLanguage").LcDropDown({
		ItemList:LanguageList(),
		Select:value,
		ItemFormat:function(value){ return value.title;}
	})//语言
}

function LanguageList(value){
	var arr=[
	{title:"中文",value:"zh"},
	{title:"日语",value:"jp"},
	{title:"西班牙语",value:"spa"},
	{title:"泰语",value:"th"},
	{title:"俄罗斯语",value:"ru"},
	{title:"粤语",value:"yue"},
	{title:"英语",value:"en"},
	{title:"韩语",value:"kor"},
	{title:"法语",value:"fra"},
	{title:"阿拉伯语",value:"ara"},
	{title:"葡萄牙语",value:"pt"},
	{title:"文言文",value:"wyw"}];
	if(value==undefined || value==null){
		return arr.concat();
	}else{
		for (var i = 0; i < arr.length; i++) {
			if(arr[i].value==value){
				return arr[i];
			}
		};
	}
	return $.extend({}, arr[0]);
}

function ControlEditView(value) {
	if (value) {
		$(".staticInfo").hide();
		$(".editInfo").show();
		$("#commitPersonInfo,#giveUpPersonInfo").show();
		$("#editPersonInfo").hide();
	} else {
		$(".staticInfo").show();
		$(".editInfo").hide();
		$("#commitPersonInfo,#giveUpPersonInfo").hide();
		$("#editPersonInfo").show();
	}
}