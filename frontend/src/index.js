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
    
    var refreshBubbleChart = function(){
        rawBubbleData = []
        normalizedData = []
        
        function HeatBar(bar){
            this._bar = bar;
        }
        
        HeatBar.prototype.createTick = function(/* float{0,1} */ atOffset, opacity){
            var bw = this._bar.offsetWidth * atOffset;
            var el = document.createElement("span");
            el.style.opacity = opacity;
            el.style.left = bw + "px";
            console.log("Createtick: ", atOffset, "bw: ", bw)
            this._bar.appendChild(el);
        };
        
        fetch(apiUrl + "bub_measurements_all")
            .then(response => response.json())
            .then(data => {
                for (meas in data){
                    rawBubbleData.push(data[meas].timestamp / 1000000)
                }

                var maxValue = Math.max(...rawBubbleData)
                var minValue = Math.min(...rawBubbleData)

                for(val in rawBubbleData){
                    normalizedData.push((rawBubbleData[val] - minValue) / (maxValue - minValue))
                }
                console.log("normalizedData:", normalizedData)
                var barElem = document.querySelector(".bar"),
                    bar = new HeatBar(barElem);
                for(tick in normalizedData)
                    bar.createTick(tick, 10);

                barElem.addEventListener('mousedown', function() {
                    this.classList.add("removeGradient");
                }, false);
                
                barElem.addEventListener('mouseup', function() {
                    this.classList.remove("removeGradient");
                }, false);
            });
        
    }



var convertToDate = function(timeInNanoSecond){
    return new Date(timeInNanoSecond/1000000)
}

var refreshCharts = function(){
    refreshTempChart()
    refreshBubbleChart()
}

refreshCharts()


