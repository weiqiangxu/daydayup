//brief_intro/echo.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func echo(wr http.ResponseWriter, r *http.Request) {
	// ReadAll从r读取，直到出现错误或EOF并返回它读取的数据。
	// 成功的调用返回err==nil，而不是err==EOF。
	// 因为ReadAll定义为从src读取直到EOF，所以它不会将EOF from read视为要报告的错误。
	// 从Go 1.16开始，这个函数只调用io.ReadAll文件
	// 从request之中读取 body

	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wr.Write([]byte("echo error"))
		return
	}
	// Write将数据作为HTTP回复的一部分写入连接。
	// 如果尚未调用WriteHeader，Write将调用WriteHeader(
	// http.StatusOK状态)在写入数据之前。如果标头不包含内容类型行，
	// Write会将内容类型集添加到将写入数据的初始512字节传递给DetectContentType的结果中。
	// 此外，如果所有写入数据的总大小小于几KB，并且没有刷新调用，则会自动添加Content-Length头。
	// 根据HTTP协议版本和客户机的不同，调用Write或WriteHeader可能会阻止将来对服务器进行读取请求。
	// 正文. 对于HTTP/1.x请求，处理程序应该在写入响应之前读取任何需要的请求体数据。一旦头被刷新（
	// 由于冲洗器冲洗调用或写入足够的数据以触发刷新），请求正文可能不可用。
	// 对于HTTP/2请求，gohttp服务器允许处理程序继续读取请求主体，同时并发地写入请求
	// 我只是想测试一下打字的感觉而已哈哈哈
	fmt.Println("msg = ", msg)
	var uu []byte
	uu = []byte("{\"hell\":\"world\"}")
	// string 和 []byte 的强转换\标准转换 - https://segmentfault.com/a/1190000037679588
	// go test做性能对比 - 用法
	// 如何理解 golang的字符串只能被替换不能被修改
	// 通过unsafe和reflect包 - 这两个包是啥子东西
	// 设置返回json格式
	// 第一种 header('Content-type: application/json');
	// 另一种 header('Content-type: text/json');
	// application/json 和 text/json有什么区别 - 没啥区别 一个官方一个非官方

	wr.Header().Set("Content-type", "application/json")
	// golang 设置web请求状态码 - net/http包
	//设置 http请求状态 为500
	wr.WriteHeader(500)
	// 设置响应为json
	// wr.Header()

	writeLen, err := wr.Write(uu)
	if err != nil || writeLen != len(msg) {
		log.Println(err, "write len:", writeLen)
		// 在go中，byte是uint8的别名
	}
	// byte的修改操作是允许的
	// b := []byte("Hello Gopher!")
	// b [1] = 'T'

	// 字符串这样是错误的 - 修改操作是禁止的
	// s := "Hello Gopher!"
	// s[1] = 'T'

	// anxios 的 500 为啥没有处理到
}

func main() {
	// http的包绑定一个路由 - 'router string','name func'
	// handler func(http.ResponseWriter, *http.Request) -
	// 必须是一个传递参数 responseWriter \ http.request 的方式
	http.HandleFunc("/user", echo)
	// 监听一个端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
