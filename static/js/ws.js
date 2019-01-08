var WS={};

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


function wsInit(callback){
    var ws;
    if (location.protocol.match("https")){
        ws = new WebSocket("wss://"+location.host+"/ws");
    }else{
        ws = new WebSocket("ws://"+location.host+"/ws");
    }
    WS.socket=ws;
    WS.socket.onopen=WS.onopen;
    WS.socket.onclose=WS.onclose;
    WS.socket.onmessage=WS.onmessage;
    WS.socket.onerror=WS.onerror;
    callback();
}
wsInit(function () {
    console.log("WS init done")
});