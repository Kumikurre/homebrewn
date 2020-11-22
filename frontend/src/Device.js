var m = require("mithril")

var apiUrl = "/api/"

var Device = {
    list: [],
    loadList: function() {
        return m.request({
            method: "GET",
            url:  apiUrl + "devices"
        })
        .then(function(result) {
            if(result === null){
                console.log("Whoopsie; no data from the endpoint: /devices")
            }
            Device.list = result
        })
    },
}

module.exports = Device
console.log("Device")