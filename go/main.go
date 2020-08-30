package main

import (
	"net/http"
)

// Docker動作検証用にHello Worldを返すサーバを立てる
func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8888", nil)
}

// HelloHandler サーバーの処理内容を記述 --- (*3)
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
