package ws

import (
	"net/http"
	"log"
	"github.com/gorilla/websocket"
	"strconv"
	"go/ast"
	"go/parser"
	"go/token"
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


func Serve(writer http.ResponseWriter,request *http.Request) {
	type msg struct{
		Type string `json:"type"`
		Data string `json:"data"`
	}
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
				log.Println("value calculate:"+m.Data)
				expr,err:=parser.ParseExpr(m.Data)
				if err !=nil{
					log.Println(err)
				}
				result:=Eval(expr)
				log.Println(result)
				m.Type="result"
				m.Data=strconv.Itoa(result)
				err=conn.WriteJSON(m)
				if err !=nil{
					log.Println(err)
				}
			}
		}
	}
}

func Eval(exp ast.Expr) int {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExpr(exp)
	case *ast.BasicLit:
		switch exp.Kind {
		case token.INT:
			i, _ := strconv.Atoi(exp.Value)
			return i
		}
	}

	return 0
}

func EvalBinaryExpr(exp *ast.BinaryExpr) int {
	left := Eval(exp.X)
	right := Eval(exp.Y)

	switch exp.Op {
	case token.ADD:
		return left + right
	case token.SUB:
		return left - right
	case token.MUL:
		return left * right
	case token.QUO:
		return left / right
	}

	return 0
}