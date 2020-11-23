var m = require("mithril")

var apiUrl = "/api/"

var TempMeasurement = {
    list: [],
    loadList: function() {
        return m.request({
            method: "GET",
            url:  apiUrl + "temp_measurements"
        })
        .then(function(result) {
            console.log("result:")
            console.log(result)
            if(result === null){
                console.log("Whoopsie; no data from the endpoint: /temp_measurements")
            }
            console.log("Tempmeasurement:")
            console.log(TempMeasurement)

            TempMeasurement.list = result
        })
    },
}

module.exports = TempMeasurement
