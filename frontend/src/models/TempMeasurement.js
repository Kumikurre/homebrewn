var m = require("mithril")

var apiUrl = "/api/"

var TempMeasurement = {
    list: [],
    loadList: function() {
        return m.request({
            method: "GET",
            url:  apiUrl + "temp_measurements_all"
        })
        .then(function(result) {

            if(result === null){
                console.log("Whoopsie; no data from the endpoint: /temp_measurements")
            }
            TempMeasurement.list = result

            console.log(TempMeasurement.list)
        })
    },
}

module.exports = TempMeasurement
