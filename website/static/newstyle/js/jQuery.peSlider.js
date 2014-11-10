/**
 * --------------------------------------------------------------------
 * jQuery peSlider plugin
 * Author: Scott Jehl, scott@filamentgroup.com
 * Copyright (c) 2009 Filament Group 
 * licensed under MIT (filamentgroup.com/examples/mit-license.txt)
 * --------------------------------------------------------------------
 */
jQuery.fn.peSlider = function(settings){

	//configurable options (none so far)
	var o = jQuery.extend({},settings);
	
	if( !jQuery('body').is('[role]') ){ jQuery('body').attr('role','application'); }
	
	return jQuery(this).each(function(){
		var thisLabel = jQuery('label[for=' + jQuery(this).attr('id') + ']').attr('id', jQuery(this).attr('id') + '-label').attr('id');
		var thisUnits = jQuery(this).attr('data-units') || '';
		var slider = jQuery('<div></div>');
		
		if( jQuery(this).is('input') ){
			var input = jQuery(this);
			var thisUnits = input.attr('data-units');
			var friendlyVal = input.val() + ' ' + thisUnits;
			var sliderOptions = jQuery.extend(o,{
				min: parseFloat(input.attr('min')),
				max: parseFloat(input.attr('max')),
				value: parseFloat(input.val())
			});
			
			slider
				.insertBefore(input)
				.slider(sliderOptions)
				.bind('slide', function(e, ui){ 
					input.val(ui.value); 
					friendlyVal = input.val() + ' ' + thisUnits;
					slider.find('a').attr({
						'aria-valuenow': ui.value,
						'aria-valuetext': friendlyVal,
						'title': friendlyVal
					});
				})
				.find('a')
				.attr({
					'role': 'slider',
					'aria-valuemin': input.attr('min'),
					'aria-valuemax': input.attr('max'),
					'aria-valuenow': input.val(),
					'aria-valuetext': friendlyVal,
					'title': friendlyVal,
					'aria-labelledby': thisLabel
				});
			
			input
				.keyup(function(){ 
					var inVal = parseFloat(input.val());
					if( !isNaN(inVal) ){ 
						slider.slider('value', inVal); 
						input.val(slider.slider('value')); 
					}				
				})
				.blur(function(){
					var inVal = parseFloat(input.val());
					if( isNaN(inVal) ){ 
						input.val(0);
					}				
				});
			
			if( !settings.step ){
				var step = Math.round( parseFloat(input.attr('max')) / slider.width());
				if(step > 1){ slider.slider('option','step',step); }
			}
			
		}
		else {		
			var select = jQuery(this);
			var friendlyVal = select.find('option').eq( select[0].selectedIndex ).text() + ' ' + thisUnits;
			var sliderOptions = jQuery.extend(o,{
				min: 0,
				max: select.find('option').length-1,
				value: select[0].selectedIndex	
			});
			
			slider
				.insertBefore(select)
				.slider(sliderOptions)
				.bind('slide', function(e, ui){ 
					select[0].selectedIndex = ui.value; 
					friendlyVal = select.find('option').eq( select[0].selectedIndex ).text() + ' ' + thisUnits;
					slider.find('a').attr({
						'aria-valuenow': ui.value,
						'aria-valuetext': friendlyVal,
						'title': friendlyVal
					});
				})
				.find('a')
				.attr({
					'role': 'slider',
					'aria-valuemin': 0,
					'aria-valuemax': select.find('option').length-1,
					'aria-valuenow': select[0].selectedIndex,
					'aria-valuetext': friendlyVal,
					'title': friendlyVal,
					'aria-labelledby': thisLabel
				});		
			
			select.bind('keyup change', function(){ 
				slider.slider('value',  select[0].selectedIndex); 
			});			
		}
		
		
	});	
};