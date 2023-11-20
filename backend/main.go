package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    // 创建一个新的路由器
    r := mux.NewRouter()

    // 定义一个处理器函数
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    })

    // 启动 HTTP 服务器
    log.Println("Listening on http://localhost:8080/")
    log.Fatal(http.ListenAndServe(":8080", r)) // 使用路由器
}
