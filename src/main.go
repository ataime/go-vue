package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Item struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	fmt.Println("run....")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	}) //设置访问的路由
	http.HandleFunc("/list", sayhelloName)   //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm() //解析参数，默认是不会解析的
	cont := r.URL.Query().Get("content")
	fmt.Println("content:", cont)
	var items = []Item{
		{
			Title:   "aaa",
			Content: "AAA",
		},
		{
			Title:   "bbb",
			Content: "BBB",
		},
	}
	if cont != "" {
		items = items[:1]
	}
	fmt.Println(items)
	v, _ := json.Marshal(items)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有源跨域请求，生产环境中应谨慎使用
	w.Write(v)
}
