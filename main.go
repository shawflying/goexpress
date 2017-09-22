package main

import (
	"fmt"
	"goexpress/controllers"
	"log"
	"net/http"
	_ "net/http/pprof"
)

type customHandler struct {
}

func (cb *customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("customHandler!!")
	w.Write([]byte("customHandler!!"))
}

//中间件
func middlewaresHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("-------------ok")
		next.ServeHTTP(w, r)
	})
}

//返回结果
func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
func main() {
	http.Handle("/img/", http.FileServer(http.Dir("static")))
	http.Handle("/js/", http.FileServer(http.Dir("static")))
	http.Handle("/css/", http.FileServer(http.Dir("static")))

	http.HandleFunc("/say", controllers.SayhelloName)  //设置访问的路由
	http.HandleFunc("/encode", controllers.Encode)     //code
	http.HandleFunc("/getInfo", controllers.GetInfo)   //code
	http.HandleFunc("/postInfo", controllers.PostInfo) //code
	http.HandleFunc("/p/index", controllers.Home)      //进入首页

	HomeHandle := http.HandlerFunc(final)               //中间处理
	http.Handle("/mid", middlewaresHandler(HomeHandle)) //中间件中间包裹一层方法

	fmt.Println("http://127.0.0.1:8600/")
	err := http.ListenAndServe(":8600", nil) //&customHandler{}
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
