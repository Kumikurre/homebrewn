var m = require("mithril")

var DeviceList = require("./views/DeviceList")
var TempMeasurements = require("./views/MeasurementView")

m.mount(document.getElementById("devicelist"), DeviceList)
m.mount(document.getElementById("graph"), TempMeasurements)



