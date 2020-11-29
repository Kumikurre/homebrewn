var m = require("mithril")

var apiUrl = "/api/"

var TempManagement = {
    list: [],

    loadList: function() {
        return m.request({
            method: "GET",
            url:  apiUrl + "device_target_temps"
        })
        .then(function(result) {

            if(result === null){
                console.log("Whoopsie; no data from the endpoint: /device_target_temps")
            }
            TempManagement.list = result

            console.log(TempManagement.list)
        })
    },
    save: function(dev) {
        console.log("Save called!", dev)
        console.log("url....", apiUrl + "device_target_temp/" + dev.device)
        dev.value = parseFloat(dev.target)
        delete dev.target
        console.log("Dev again...", dev)
        return m.request({
            method: "POST",
            url: apiUrl + "device_target_temp/" + dev.device,
            body: dev,
        })
    },
}

module.exports = TempManagement
