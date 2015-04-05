
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
		populate(data);
		updateTables()
	});
	
});

/*
	Populate location/serial selecter
	obj: raw json from the /locationsInfo call
*/
function populate(obj) {
	var s1 = document.getElementById("slct1");
	s1.innerHTML = "";

	for (var loc in obj) {
		var location = obj[loc].location;
		var serials = obj[loc].serials;

		var newOptGroup = document.createElement("optgroup");
		newOptGroup.label = location;
		s1.appendChild(newOptGroup);

		for (var ser in serials) {
			var serial = serials[ser];

			var newOption = document.createElement("option");
			newOption.value = serial;
			newOption.innerHTML = serial;
			newOptGroup.appendChild(newOption);
		}
	}
}

/*
	Returns the scientific unit for a type. ex("F->Hz")
	type: type of registor ("F")
*/
function stringForType(type) {
	if (type == 'I') {
		return "Current (Amps)";
	}
	else if (type == 'F') {
		return "Frequency (Hz)";
	}
	else if (type == 'V') {
		return "Voltage (V)";
	}
}

/*
	Loads all the tables for a selected serial and location
*/
function updateTables() {
	//registersInfo/location/TxState/serial/0001
	var serial = document.getElementById("slct1").value;
	var location = document.getElementById("slct1").options[document.getElementById("slct1").selectedIndex].parentNode.label;
	var url = "/registersInfo/location/" + location + "/serial/" + serial;
	$.getJSON(url, function(regs) {
		console.log(regs);
		var tables = document.getElementById("tables");
		tables.innerHTML = "";
		
		var s1 = document.getElementById("reg1");
		s1.innerHTML = "";
		for (var reg in regs) {
			var newOption = document.createElement("option");
			newOption.innerHTML = regs[reg].name;
			newOption.value = regs[reg].name;
			s1.appendChild(newOption);
			
			loadTableForReg(regs[reg]);
		}
		//setInterval(cycleTables,3000);
		
		
	});
}

/*
	Loads a table for a single registor. Called by updateTables().
*/
function loadTableForReg(reg) {
	var location = document.getElementById("slct1").options[document.getElementById("slct1").selectedIndex].parentNode.label;
	var serial = document.getElementById("slct1").value;
	var url = "/measurements/location/"+location+"/serial/"+serial+"/reg/"+reg.name+"/start/2014-12-16T05:07:00Z/end/2015-12-17T14:07:00Z";
	$.getJSON(url, function(data) { 
	
		//CREATE HTML
		var tables = document.getElementById("tables");
		var gridbox = document.createElement('div');
		gridbox.setAttribute('class', 'pure-u-1 pure-u-md-1-2 pure-u-lg-1-3');
		
		var tablebox = document.createElement('div');
		tablebox.setAttribute('class', 'gr');
		
		var container = document.createElement('div');
		container.setAttribute('class', 'tc');
		container.setAttribute('id', reg.name);
		
		var btn = document.createElement('BUTTON');
		btn.setAttribute('class', 'opener');
		$(btn).click(openDialog);
		var t = document.createTextNode("Show History");
		btn.appendChild(t); 
		tablebox.appendChild(btn);
		
		tablebox.appendChild(container);
		gridbox.appendChild(tablebox);
		tables.appendChild(gridbox);
		
		console.log(data);
		
		//get chart then set
		var sel = "#"+reg.name;
		registerCharts[reg.name] = $(sel).highcharts('StockChart', {
			rangeSelector : {
				//selected : 2
				enabled: true
			},
			title : {
				text : reg.name
			},
			series : [{
				name : 'Volts',
				data : data.data,
				tooltip: {
					valueDecimals: 2
				}
			}]
		});
		
		
		
	});
}

function cycleTables() {
	$.each( registerCharts, function(index,value){
		console.log(value); 
		requestData(index,value);
	})
}

function requestData(reg,chart) {
    $.ajax({
        url: 'http://txsolar.mooo.com/lastmeasurement/loc/TxState/ser/0001/reg/'+reg,
        success: function(point) {
        	
        	//var chart = registerCharts[reg.name];
        	
        	var someDate = new Date(point.time);
			someDate = someDate.getTime();
			var x = [someDate, point.value];
        	
        	
            var series = chart.series[0],
                shift = series.data.length > 20; // shift if the series is 
                                                 // longer than 20

            // add the point
            chart.series[0].addPoint(x, true, shift);
            
            // call it again after one second
            //setTimeout(function(){requestData(reg,chart)}, 1000);    
        },
        cache: false
    });
}
	
function openDialog(event) {	
	$( "#dialog" ).dialog( "open" );
	console.log(event);
}



function pastMonth() {
//Get data from past month
//Reload Chart with data from the past month
//
}
function pastWeek() {
//Same thing, just with past week
}
function pastDay() {
//Same thing with past day
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

    