
//GLOBALS
//var regsMap; // Unused atm
var registerCharts = {};

/* 
	Inital function sets up the view with some data.
*/ 
$(document).ready(function () {
	
	$( "#dialog" ).dialog({
		autoOpen: false,
	});
	
	var url = '/locationsInfo';
	$.getJSON(url, function(data) {
		sel_pop(data);
		//populate(data);
		
	});
	
});

function sel_pop(obj) {
	for (var loc in obj) {
		var location = obj[loc].location;
 		var serials = obj[loc].serials;
 		
		for (var ser in serials) {
			var serial = serials[ser];
			var name = serial.serial
			var regs =  serial.regs;
			
			$("#loc-ser_selector").append('<div class="pure-u-1"><h5>'+location+' - '+name+'</h5></div>');
			
			for (var reg in regs) {
 				var register = regs[reg];
 				var is = location+"-"+serial.serial+"-"+register.name+"-"+register.type;
 				//console.log("hehe");
 				$("#loc-ser_selector").append('<div class="pure-u-1-2 pure-u-md-1-3"><div class="center"><label for="'+is+'" class=" pure-checkbox"><div class="pure-button pure-u-1 checkButton"><input id="'+is+'" type="checkbox" value="" checked>'+register.name+'</div></label></div></div>');
 				//$("#"+is).change(checkboxclicked);
 				
 				loadTableForReg(location,serial.serial,register.name,register.type);
 			}
		}
	}
}

/*
	Populate location/serial selecter
	obj: raw json from the /locationsInfo call
*/
function populate(obj) {
	//var s1 = document.getElementById("slct1");
	//s1.innerHTML = "";

	for (var loc in obj) {
		var location = obj[loc].location;
 		var serials = obj[loc].serials;
 /*
 		var newOptGroup = document.createElement("optgroup");
 		newOptGroup.label = location;
 		s1.appendChild(newOptGroup);
 */
 		for (var ser in serials) {
 			var serial = serials[ser];
 			var regs =  serial.regs;
 /*
 			var newOption = document.createElement("option");
 			newOption.value = serial.serial;
 			newOption.innerHTML = serial.serial;
 			newOptGroup.appendChild(newOption);
*/ 			
 			for (var reg in regs) {
 				loadTableForReg(location,serial.serial,regs[reg].name,regs[reg].type);
 			}
 			
 		}
	}
}

/*
	Returns the scientific unit for a type. ex("F->Hz")
	type: type of registor ("F")
*/
function stringForType(type) {
	if (type == 'I') {
		return "Ampère (I)";
	}
	else if (type == 'F') {
		return "Frequency (Hz)";
	}
	else if (type == 'V') {
		return "Voltage (V)";
	} 
	else if (type == 'Ee') {
		return "Irradiance (W/m²)";
	}
	else if (type == 'T'){
		return "Temperature (C)";
	}
}

/*
	Loads a table for a single register. Called by updateTables().
*/
function loadTableForReg(loc,ser,reg,type) {
	var url = "/measurements/location/"+loc+"/serial/"+ser+"/reg/"+reg+"/start/2014-12-16T05:07:00Z/end/2015-12-17T14:07:00Z";
	$.getJSON(url, function(data) { 
	
		var sel = ser+reg;
		sel = sel.replace(/\|/g, "r")
		
		//CREATE HTML
		var tables = document.getElementById("tables");
		var gridbox = document.createElement('div');
		gridbox.setAttribute('class', 'pure-u-1 pure-u-md-1-2 pure-u-lg-1-2');
		gridbox.setAttribute('id', sel+'g');
		
		var tablebox = document.createElement('div');
		tablebox.setAttribute('class', 'gr');
		
		var container = document.createElement('div');
		container.setAttribute('class', 'tc');
		container.setAttribute('id', sel);
		
// 		var btn = document.createElement('BUTTON');
// 		btn.setAttribute('class', 'opener');
// 		$(btn).click(openDialog);
// 		var t = document.createTextNode("Show History");
// 		btn.appendChild(t); 
// 		tablebox.appendChild(btn);
		
		tablebox.appendChild(container);
		gridbox.appendChild(tablebox);
		tables.appendChild(gridbox);
		
		console.log(data);
		sel = "#"+sel
		console.log(sel)
		//get chart then set
		registerCharts[reg] = $(sel).highcharts('StockChart', {
			chart : {
				events : {
					load : function () {

						// set up the updating of the chart each second
						var series = this.series[0];
						setInterval(function () {
							var url = 'http://txsolar.mooo.com/lastmeasurement/loc/'+loc+'/ser/'+ser+'/reg/'+reg;
							$.getJSON(url, function(point) { 
									var someDate = new Date(point.time);
									someDate = someDate.getTime();
									var x = [someDate, point.value];
									series.addPoint(x, true, true);
							});
						}, 60000);
					}
				}
			},
			
			credits: {
            	enabled: false
        	},
        	exporting: {
        		buttons: {
					contextButton: {
						menuItems: [{
							text: 'Export to PNG (small)',
							onclick: function() {
								this.exportChart({
									width: 250
								});
							}
						}, {
							text: 'Export to PNG (large)',
							onclick: function() {
								this.exportChart(); // 800px by default
							}
						}, {
							text: 'Export to CSV',
							onclick: function() {
								window.open('/getcsv/loc/'+loc+'/ser/'+ser+'/reg/'+reg);
							}
						}, 
						null
						]
					}
				}
        	},
			rangeSelector : {
				//selected : 2
				enabled: true
			},
			title : {
				text : ser+": "+reg
			},
			series : [{
				name : stringForType(type),
				data : data.data,
				tooltip: {
					valueDecimals: 2
				}
			}]
			
		});
		
// 		Highcharts.getOptions().exporting.buttons.contextButton.menuItems.push({
// 			text: 'My new button',
// 			onclick: function () {
// 				alert('OK');
// 			}
// 		});
		
		//requestData(reg,container);
		
		
		
	});
}

function openDialog(event) {	
	$( "#dialog" ).dialog( "open" );
	console.log(event);
}

function checkboxclicked(event) {	
	//$( "#dialog" ).dialog( "open" );
	console.log(event.target.checked + event.target.id);
	var res = event.target.id.split("-");
	if (event.target.checked) {
		loadTableForReg(res[0],res[1],res[2],res[3])
	} else {
		$("#"+res[1]+res[2]+"g").remove();
	}
}

//This is just some code for the functionality of the checkboxes
$(function CheckBoxes() {
	// Apparently click is better chan change? Cuz IE?
	$('input[type="checkbox"]').change(function(e) {
		var checked = $(this).prop("checked"),
		container = $(this).parent(),
		siblings = container.siblings();

	  	container.find('input[type="checkbox"]').prop({
			indeterminate: false,
			checked: checked
		});

		function checkSiblings(el) {
			var parent = el.parent().parent(), all = true;

			el.siblings().each(function() {
				return all = ($(this).children('input[type="checkbox"]').prop("checked") === checked);
			});

		  	if (all && checked) {
				parent.children('input[type="checkbox"]').prop({
					indeterminate: false,
					checked: checked
				});
				checkSiblings(parent);
		  	} else if (all && !checked) {
			  	parent.children('input[type="checkbox"]').prop("checked", checked);
			  	parent.children('input[type="checkbox"]').prop("indeterminate", (parent.find('input[type="checkbox"]:checked').length > 0));
			  	checkSiblings(parent);
		  	} else {
			  	el.parents("li").children('input[type="checkbox"]').prop({
					indeterminate: true,
					checked: false
			  	});
		  	}
		}
		checkSiblings(container);
	});
});

    