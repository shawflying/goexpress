# net.http

ResponseWriter： 生成Response的接口

Handler： 处理请求和生成返回的接口

ServeMux： 路由，后面会说到ServeMux也是一种Handler

Conn : 网络连接


##　配置优化
```text
	http.Handle("/img/", http.FileServer(http.Dir("static")))


	http.HandleFunc("/say", controllers.SayhelloName)  //设置访问的路由

	HomeHandle := http.HandlerFunc(final)               //中间处理
	http.Handle("/mid", middlewaresHandler(HomeHandle)) //中间件中间包裹一层方法

	err := http.ListenAndServe(":8600", nil) //&customHandler{}

```
Server实现的接口如下：
```text
func (srv *Server) ListenAndServe() error // 监听server，监听到有请求时，调用handler。
func (srv *Server) Serve(l net.Listener) error   //对某个端口进行监听，里面就是调用for进行accept的处理了
func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error //开启https server服务，内部调用Serve
```

http包也提供了外部使用的几个方法。但实际上是调用Server的内部方法

```text
func ListenAndServe(addr string, handler Handler) error   //开启Http服务
func ListenAndServeTLS(addr string, certFile string, keyFile string, handler Handler) error //开启HTTPs服务
```