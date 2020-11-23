var m = require("mithril")
var TempMeasurement = require("../models/TempMeasurement")

module.exports = {
    oninit: TempMeasurement.loadList,
    view: function() {
        return m(".measurements", TempMeasurement.list.map(function(meas) {
            [
                m(".meas-name", meas.device.name),
                m(".meas-val", meas.value),
                m(".meas-time", meas.timestamp)
            ]
        }))
    }
}


