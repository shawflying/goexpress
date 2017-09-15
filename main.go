package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

func sayhelloName(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(req.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", req.URL.Path)
	fmt.Println("scheme", req.URL.Scheme)
	fmt.Println(req.Form["url_long"])
	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(res, "Hello goexpress!") //这个写入到w的是输出到客户端的
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static/")))
	http.HandleFunc("/say", sayhelloName) //设置访问的路由
	err := http.ListenAndServe(":8600", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
