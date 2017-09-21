# go中提供了pprof包来做代码的性能监控，在两个地方有包：


参考：http://www.cnblogs.com/ghj1976/p/5473693.html

````text
net/http/pprof
runtime/pprof
````

其实net/http/pprof中只是使用runtime/pprof包来进行封装了一下，并在http端口上暴露出来。

使用 net/http/pprof 做WEB服务器的性能监控
如果你的go程序是用http包启动的web服务器，想要查看自己的web服务器的状态。这个时候就可以选择net/http/pprof。
```text

   import _ "net/http/pprof"
```
然后就可以在浏览器中使用http://localhost:port/debug/pprof/ 直接看到当前web服务的状态，包括CPU占用情况和内存使用情况等。
当然，非WEB的也可以用下面方式启动WEB。
在 main 方法中增加
```go
func main() {
    go func() {
        http.ListenAndServe("localhost:6060", nil)
    }()
}
```
下图就是访问该网址的一次截图：


CPU消耗分析
使用 runtime/pprof 做应用程序性能监控
关键代码：

```go
import  "runtime/pprof"

func main() {
    f, err := os.OpenFile("./tmp/cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
}
```
注意，有时候 defer f.Close()， defer pprof.StopCPUProfile() 会执行不到，这时候我们就会看到 prof 文件是空的， 我们需要在自己代码退出的地方，增加上下面两行，确保写文件内容了。

```go
pprof.StopCPUProfile()

f.Close()
```

 

对产生的文件进行分析：
我们可以使用 go tool pprof (应用程序) （应用程序的prof文件） 方式来对这个 prof 文件进行分析。

```text
$ go tool pprof HuaRongDao ./tmp/cpu.prof 
Entering interactive mode (type "help" for commands)
(pprof) 

```
一些常用 pprof 的命令：

top

在默认情况下，top命令会输出以本地取样计数为顺序的列表。我们可以把这个列表叫做本地取样计数排名列表。

```text
(pprof) top
2700ms of 3200ms total (84.38%)
Dropped 58 nodes (cum <= 16ms)
Showing top 10 nodes out of 111 (cum >= 80ms)
      flat  flat%   sum%        cum   cum%
     670ms 20.94% 20.94%      670ms 20.94%  runtime.mach_semaphore_signal
     580ms 18.12% 39.06%      590ms 18.44%  runtime.cgocall
     370ms 11.56% 50.62%      370ms 11.56%  runtime.mach_semaphore_wait
     360ms 11.25% 61.88%      360ms 11.25%  runtime.memmove
     210ms  6.56% 68.44%      580ms 18.12%  golang.org/x/mobile/gl.(*context).DoWork
     120ms  3.75% 72.19%      120ms  3.75%  runtime.usleep
     110ms  3.44% 75.62%      110ms  3.44%  image/png.filterPaeth
     100ms  3.12% 78.75%      160ms  5.00%  compress/flate.(*decompressor).huffSym
     100ms  3.12% 81.88%      100ms  3.12%  image/draw.drawNRGBASrc
      80ms  2.50% 84.38%       80ms  2.50%  runtime.memclr
(pprof) 
```

参考： https://github.com/hyper-carrot/go_command_tutorial/blob/master/0.12.md 

默认情况下top命令会列出前10项内容。但是如果在top命令后面紧跟一个数字，那么其列出的项数就会与这个数字相同。

 

web

与gv命令类似，web命令也会用图形化的方式来显示概要文件。但不同的是，web命令是在一个Web浏览器中显示它。如果你的Web浏览器已经启动，那么它的显示速度会非常快。如果想改变所使用的Web浏览器，可以在Linux下设置符号链接/etc/alternatives/gnome-www-browser或/etc/alternatives/x-www-browser，或在OS X下改变SVG文件的关联Finder。

mac 下 修改默认打开方式： 右键一个想处理的文件，按alt 键（lion）出现always open with，然后打开，整个过程中， 先右键，然后一直按 alt， 一直到打开为止。