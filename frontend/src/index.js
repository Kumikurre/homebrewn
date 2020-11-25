var m = require("mithril")
var c3 = require("c3")

var DeviceList = require("./views/DeviceList")
var TempMeasurements = require("./views/MeasurementView")

m.mount(document.getElementById("devicelist"), DeviceList)
m.mount(document.getElementById("graphdata"), TempMeasurements)

var apiUrl = "/api/"
var refreshTempChart = function(){
    output = {
                xs: {},
                columns:[]
            }

    deviceIndex = 1
    tempValues = {}
    fetch(apiUrl + "temp_measurements_all")
        .then(response => response.json())
        .then(data => {
            for(val in data){

                if(!(data[val].device in output.xs)){
                    output.xs[data[val].device] = deviceIndex.toString()

                    tempValues[deviceIndex.toString()] = []
                    tempValues[data[val].device] = []

                    deviceIndex = deviceIndex + 1
                }
                tempValues[data[val].device].push(data[val].value)
                timeStampName = output.xs[data[val].device]
                tempValues[timeStampName].push(convertToDate(data[val]["timestamp"]))

            }

            for (const [key, value] of Object.entries(output.xs)) {
                console.log(`${key}: ${value}`);
                keyList = [key]
                valueList = [value]

                keyListFinal = keyList.concat(tempValues[key])
                valueListFinal = valueList.concat(tempValues[value])

                output.columns.push(keyListFinal)
                output.columns.push(valueListFinal)
              }
            console.log(output)

            var chart = c3.generate({
                bindto: '#graph',
                data: output,
                axis : {
                    x : {
                        type : 'timeseries',
                        tick: {
                            format: function (x) { return x.toLocaleString('en-GB'); }
                        }
                    }
                }
            });
            
        });
    }

var convertToDate = function(timeInNanoSecond){
    return new Date(timeInNanoSecond/1000000)
}

refreshTempChart()


