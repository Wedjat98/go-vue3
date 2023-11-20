package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    // 定义一个处理器函数
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    })

    // 启动 HTTP 服务器
    log.Println("Listening on http://localhost:8080/")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
