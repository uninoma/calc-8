var WS = new WebSocket("ws://"+location.host+"/ws");
WS.onopen=function () {
    console.log("web socket open")
}
WS.onmessage=function (ev) {
    console.log(ev.data)
    var data = JSON.parse(ev.data);
    console.log(data.type)
    switch (data.type){
        case "result":
            $("#display .input").text(data.data);
        case "MR":
            // $("#display .input").text($("#display .input").text()+data.data);

    }
}