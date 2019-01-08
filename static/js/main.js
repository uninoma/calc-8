$(document).ready(
    function() {
        $("button").click(function(event) {
            console.log(event.target.innerText)
            displayController(event.target.innerText,textDecrease)
        });
    }
);

function displayController(val,callback) {
    if($("#display .input").text()[0]==0 && $("#display .input").text()[1] != "." && val!="."){
        $("#display .input").text($("#display .input").text().slice(1))
    }
    $("#display .input").text($("#display .input").text().replace(/^\+/,"0+"))
    $("#display .input").text($("#display .input").text().replace(/^\-/,"0-"))
    $("#display .input").text($("#display .input").text().replace(/^\//,"0/"))
    $("#display .input").text($("#display .input").text().replace(/^\*/,"0*"))
    switch(val){
        case "C":
            $("#display .input").html("");
            break;
        case "⇜"://backspace
            $("#display .input").text($("#display .input").text().slice(0,-1))
            break;
        case "MC":
            console.log(val);
            var dataToSend = JSON.stringify({"type":"MC","value":"MC"});
            WS.send(dataToSend);
            break;
        case "MR":
            sendData("MR","")
            break;
        case "M+":
            sendData("M+",$("#display .input").text())
            break;
        case "M-":
            sendData("M-",$("#display .input").text())
            break;
        case "=":
            sendData("calculate",$("#display .input").text())
            break;
        default:
            $("#display .input").html($("#display .input").html()+val);

            break;
    }
    callback()
}

function sendData(type,data) {
    var data = data;



    var dataToSend = JSON.stringify({"type":type,"data":data});
    WS.socket.send(dataToSend,function (ev) {
        if(!ev){
            wsInit();
        }
    });
}
function textDecrease(){
    var inputWidth=$(".input").width()
    var displayWidth=$("#display").width();
    console.log("input width:",inputWidth);
    console.log("display width:",displayWidth);
    var diffrence=inputWidth-displayWidth;
    console.log("raznica:",diffrence)
    if(diffrence>-40){
        $(".input").css("font-size", $(".input").css("font-size").replace("px","") - 12);
    }
    if($(".input").text().length<18){
           $(".input").css("font-size",window.getComputedStyle(document.getElementById("display"))["font-size"])
    }
}
