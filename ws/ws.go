package ws

import (
	"net/http"
	"log"
	"github.com/gorilla/websocket"
	"github.com/Knetic/govaluate"
	"fmt"
	"strings"
)
var lastId=0;
var clients=make(map[int]*Client)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


type Client struct {
	conn  *websocket.Conn
}

type msg struct{
	Type string `json:"type"`
	Data string `json:"data"`
}

func Serve(writer http.ResponseWriter,request *http.Request) {

	conn, err := upgrader.Upgrade(writer,request,nil)
	log.Println("ws request")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("ws connected")
	client :=&Client{conn}
	var id int
	id=lastId+1
	lastId=id

	log.Println("ws client:",id)
	clients[id]=client
	log.Println("clients len:",len(clients))
	m:=msg{}
	memory:=""
	for {
		err=conn.ReadJSON(&m)
		if err!=nil{
			log.Println("readJSON err:",err)
			delete(clients,id)
			break
		}else{
			switch m.Type{
			case "alive?":
				m.Data="yes"
				conn.WriteJSON(m)
			case "MC":
				memory=""
			case "MR":
				m.Type="MR"
				m.Data=memory
				conn.WriteJSON(m)
			case "M+":
				memory=m.Data
			case "M-":
				memory="-"+m.Data
			case "calculator":
				print(m.Data)
			case "calculate":
				err:=conn.WriteJSON(calculate(m))
				if err !=nil{
					log.Println(err)
				}
			}
		}
	}
}
func calculate(m msg) msg{
	log.Println("value calculate:"+m.Data)
	resultStr:=""
	if m.Data!="" && !strings.Contains(m.Data,"e"){
		expression, err := govaluate.NewEvaluableExpression(string(m.Data))
		result, err := expression.Evaluate(nil)
		if err !=nil{
			fmt.Println(err)
		}
		resultStr = fmt.Sprintf("%v", result)
	}else{
		resultStr="0"
	}
	log.Println(resultStr)
	m.Data=resultStr
	m.Type="result"
	return m
}