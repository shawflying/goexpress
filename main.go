package main

import (
	"fmt"
	"goexpress/controllers"
	"log"
	"net/http"
)

func main() {

	http.Handle("/img/", http.FileServer(http.Dir("static")))
	http.Handle("/js/", http.FileServer(http.Dir("static")))
	http.Handle("/css/", http.FileServer(http.Dir("static")))

	http.HandleFunc("/say", controllers.SayhelloName) //设置访问的路由
	http.HandleFunc("/encode", controllers.Encode)    //code
	http.HandleFunc("/getInfo", controllers.GetInfo)  //code

	http.HandleFunc("/p/index", controllers.Home) //进入首页

	fmt.Println("http://127.0.0.1:8600/")
	err := http.ListenAndServe(":8600", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
