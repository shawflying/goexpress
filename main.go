package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
	"github.com/shawflying/beego-common-utils/utils/comutil"
	"goexpress/request"
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

//加密
func encode(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	fmt.Println(comutil.TransInterfaceToString(req))
	fmt.Println("请求连接：", req.URL)
	fmt.Println("req.Host:", req.Host)
	fmt.Println("req.Body:", comutil.TransInterfaceToString(req.Body))
	v, _ := req.Cookie("cookie")
	fmt.Println("req.cookie:", v)
	fmt.Println("req.GetBody:", req.GetBody)
	fmt.Println("req.Method:", req.Method)
	fmt.Println("req.Form:", req.Form)
	fmt.Println("req.RequestURI:", req.RequestURI)
	fmt.Println("req.Response:", req.Response)

	//cookie1:= make(http.Cookie{})
	//fmt.Println("req.GetBody:", req.AddCookie())

	fmt.Println(" req.Header ------------------------------------------------------")
	fmt.Println("通过 header 获取 User-Agent", req.Header.Get("User-Agent"))
	fmt.Println("使用map取值：", req.Header["User-Agent"])
	fmt.Println(comutil.TransInterfaceToString(req.Header))

	fmt.Println(" res ------------------------------------------------------")

	//fmt.Println(res)
	//fmt.Println(comutil.TransInterfaceToString(res))
	//fmt.Println(res.Header())

	fmt.Fprintf(res, `{
    "status": 200,
    "data": "36b59e39b935e1fcf05065d260177c5a"
}`) //这个写入到w的是输出到客户端的
}

// 接口请求 get post ...
func getInfo(res http.ResponseWriter, req *http.Request) {
	body, err := request.Get("http://m.sh.189.cn/service/node/crypto?data=abc123&key=express&type=0")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(res, comutil.TransInterfaceToString(body))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static/")))
	http.HandleFunc("/say", sayhelloName) //设置访问的路由
	http.HandleFunc("/encode", encode)    //code
	http.HandleFunc("/getInfo", getInfo)  //code

	fmt.Println("http://127.0.0.1:8600/")
	err := http.ListenAndServe(":8600", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
