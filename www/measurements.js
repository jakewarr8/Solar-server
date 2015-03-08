<!-- hide script from old browsers

	//GLOBALS
	var regsMap; // Hold the registors. Indexed by registors name. ex:["L1V]

	/* 
		Inital function sets up the view with some data.
	*/ 
    function init() {
		$(function() {
			var url = '/locationsInfo';
			$.getJSON(url, function(data) {
				populate(data);
				//submitUrlRequest(); 
				updateRegistors()
			});
		});
	}
	
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


    function updateRegistors() {
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
			
			
		});
    }
    
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
			
			tablebox.appendChild(container);
			gridbox.appendChild(tablebox);
			tables.appendChild(gridbox);
			
			
		
			console.log(data);
 			var array1 = data.data;
 			var pointsArray = new Array();
 
 			array1.forEach(function(arrayItem) {
 				var someDate = new Date(arrayItem.time);
 				someDate = someDate.getTime();
				var x = [someDate, arrayItem.value];
 				pointsArray.push(x);
 			});	
 			
 			var sel = "#"+reg.name;
			$(sel).highcharts('StockChart', {
				rangeSelector : {
					//selected : 2
					enabled: true
				},

				title : {
					text : reg.name
				},
				
				series : [{
					name : 'Volts',
					data : pointsArray,
					tooltip: {
						valueDecimals: 2
					}
				}]
			});
 			
 			
		});
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

    
// end hiding script from old browsers -->
