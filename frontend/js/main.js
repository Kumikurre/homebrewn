var root = document.body

m.render(root, [
    m("main", [
        m("h1", {class: "title"}, "This text is rendered from javascript"),
        m("button", "A button"),
    ])
])


var get_devices = function() {
    m.request({
        method: "GET",
        url: "/api/devices"
    })
    .then(function(data) {
        m.render(root.output, [
            m("output", [
                m("h2", {class: "title"}, data)
            ])
        ])
    })
}

get_devices()