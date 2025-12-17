package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", "0.0.0.0:8080", "监听地址:监听端口")
}

func main() {

	flag.Parse()
	
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, http.FileServer(http.Dir("."))))
	log.Printf("HTTP服务器已启动")
	log.Printf("监听地址: %s", port)
	log.Printf("提示: 可以使用 -port 参数修改监听地址和端口")
	log.Printf("例如: -port 127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(port, nil))

}
