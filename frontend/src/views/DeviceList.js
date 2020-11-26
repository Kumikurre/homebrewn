var m = require("mithril")
var Device = require("../models/Device")

module.exports = {
    oninit: Device.loadList,
    view: function() {
        return m(".device-list", Device.list.map(function(device) {
            return m(".device-list-item", [
                m(".device-name", device.name),
                m(".device-sensors", device.sensors.map(function(sensor){
                    return m(".sensor", sensor)
                }))
            ])
        }))
    }
}
