<!-- hide script from old browsers

	//GLOBALS
	var regsMap; // Hold the registors. Indexed by registors name. ex:["L1V]

	/* 
		Inital function sets up the view with some data.
	*/ 
    function init() {
    
		reloadChart(); //set empty chart

		$today = new Date();
		$yesterday = new Date($today);
		$yesterday.setDate($today.getDate() - 1);

		$tomorrow = new Date($today);
		$tomorrow.setDate($today.getDate() + 1);

		$(function () {
			$('#beginDate').val($yesterday.getFullYear()+"/"+"02"+"/"+("0" + $yesterday.getDate()).slice(-2));
			$('#endDate').val($tomorrow.getFullYear()+"/"+"02"+"/"+("0" + $tomorrow.getDate()).slice(-2));
			$('#beginTime').val("09:00");
			$('#endTime').val("09:00");
		});
		
		$(function() {
			var url = '/locationsInfo';
			$.getJSON(url, function(data) {
				populate(data);
				submitUrlRequest(); 
			});
		});
		
	}
	
    /*
    	Grabs the selected dates/times serial and location then make JSON request.
    	JSON async block loads the json and updates the view with some data.
	*/
    function submitUrlRequest() {
        var serial = document.getElementById("slct1").value;
        var location = document.getElementById("slct1").options[document.getElementById("slct1").selectedIndex].parentNode.label;
        var beginDate = document.getElementById('beginDate').value;
        var beginTime = document.getElementById('beginTime').value;
        var endDate = document.getElementById('endDate').value;
        var endTime = document.getElementById('endTime').value;

        beginDate = beginDate.replace(/\//g, '-');
        endDate = endDate.replace(/\//g, '-');

        var urlEx = "/measurements/location/_location/serial/_serial/start/begindateTbegintime:00Z/end/enddateTendtime:00Z";

        urlEx = urlEx.replace("_location", location);
        urlEx = urlEx.replace("_serial", serial);
        urlEx = urlEx.replace("begindate", beginDate);
        urlEx = urlEx.replace("begintime", beginTime);
        urlEx = urlEx.replace("enddate", endDate);
        urlEx = urlEx.replace("endtime", endTime);

        var requestUrl = urlEx;

		$.getJSON(requestUrl, function(data) {
			loadRegs(data);
			populateRegisters(regsMap);
			showChartForKey("L1V") //FOR DEMO REMOVE OR FIX
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
    	Populate the registor's selecter.
    	regsMap: the registor's hash map a using key index. ex("L1V")
    */
    function populateRegisters(regsMap) {
        var s1 = document.getElementById("reg1");
        s1.innerHTML = "";

        for (var i in regsMap) {
            //console.log(regsMap[i]); //For testing

            var newOption = document.createElement("option");
            newOption.innerHTML = regsMap[i].name;
            newOption.value = regsMap[i].name;
            s1.appendChild(newOption);
        }
    }
    
    /*
    	Creates regsMap. A hash map for the registors.
    	measurements: raw json from /measurements call.
    */
    function loadRegs(measurements) {
        regsMap = {};
        for (var meas in measurements) {
            var time = measurements[meas].time;
            var regs = measurements[meas].registers;

            for (reg in regs) {
                if (regs[reg].name in regsMap) {
                    var regDict = regsMap[regs[reg].name];
                    var regA = regDict.regArray;
                    var regMeas = {
                        time: time,
                        data: regs[reg].data
                    };
                    regA.push(regMeas);

                } else {
                    var regA = new Array();

                    var regMeas = {
                        time: time,
                        data: regs[reg].data
                    };
                    regA.push(regMeas);

                    var regDict = {
                        name: regs[reg].name,
                        type: regs[reg].type,
                        regArray: regA
                    }

                    regsMap[regs[reg].name] = regDict
                }

            }
        }
        for (var i in regsMap) {
            console.log(regsMap[i]); //For testing
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
    	Reads data from regsMap for selected registor. Then reloadsChart with that data.
    	key: key for registor. ex("L1V")
    */
    function showChartForKey(key) {
        //console.log(key);
        var array1 = regsMap[key].regArray;
        var pointsArray = new Array();
        //alert(array1);
        array1.forEach(function(arrayItem) {
            var someDate = new Date(arrayItem.time);
            someDate = someDate.getTime();

            var x = [someDate, arrayItem.data];
            pointsArray.push(x);
        });
        yAxisName = regsMap[key].type;
        reloadChart(pointsArray, stringForType(yAxisName));
    }
        
	/*
		Reloads HighCharts Graph
		data: contains the x&y values.
		yAxisName: data type for y.
	*/
    function reloadChart(data, yAxisName) {

        $(function() {
            $('#Chart').highcharts({
                chart: {
                    type: 'line'
                },
                title: {
                    text: (yAxisName + " Chart")
                },
                xAxis: {
                    title: {
                        text: "Time"
                    },
                    type: 'datetime',
                    dateTimeLabelFormats: { // don't display the dummy year
                        month: '%e. %b',
                        year: '%b'
                    }
                },
                yAxis: {
                    title: {
                        text: yAxisName
                    }
                },
                series: [{
                    name: 'Data',
                    data: data
                }]
            });
        });

    }

    
// end hiding script from old browsers -->