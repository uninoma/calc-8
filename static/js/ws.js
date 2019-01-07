var WS;
function wsInit(callback){
    if (location.protocol.match("https")){
        WS = new WebSocket("wss://"+location.host+"/ws");
    }else{
        WS = new WebSocket("ws://"+location.host+"/ws");
    }
    callback();
}
wsInit(function () {
    console.log("WS init done")
});

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
            break;
        case "MR":
            $("#display .input").text($("#display .input").text()+data.data);
            break;

    }
}

WS.onclose=function (ev) {
    console.log("socket closed",ev)
    setTimeout(function () {
        wsInit(function () {
            console.log("WS init after close DONE")
        });
    },2000);
}

WS.onerror=function (ev) {
    setTimeout(function () {
        wsInit(function () {
            console.log("WS init after error DONE")
        });
    },2000);
}