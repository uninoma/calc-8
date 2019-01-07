package webServer

import (
	"net/http"
	"fmt"
	"path/filepath"
	"os"
	"log"
	"github.com/uninoma/calculator/ws"
)

func Init() {
	port:=os.Getenv("PORT")
	if port == ""{
		port=":3434"
	}else{
		port=":"+port
	}
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/yo",func(res http.ResponseWriter,req *http.Request){
		fmt.Fprintf(res,"hello yo")
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dir)
	});
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		ws.Serve(writer,request)
	})
	fmt.Println("starting http server port:",port)
	log.Fatal(http.ListenAndServe(port,nil))

}
