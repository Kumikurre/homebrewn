var m = require("mithril")
var Device = require("./Device")

module.exports = {
    oninit: Device.loadList,
    view: function() {
        return m(".device-list", Device.list.map(function(device) {
            return m(".device-list-item", device.name + "     " + device.censors)
        }))
    }
}
console.log("DeviceList")