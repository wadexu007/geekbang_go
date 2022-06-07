package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"httpserver/utilities"

	"github.com/golang/glog"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		//当访问 localhost/healthz 时，应返回 200
		glog.Info("Entering healthz handler")
		io.WriteString(w, "ok\n")
	case "POST":
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	default:
	}
}

func RootHandler(rw http.ResponseWriter, req *http.Request) {
	glog.Info("Entering root handler")

	// fmt.Println(os.Environ())
	os.Setenv("Version", "1.0.0")
	version := os.Getenv("Version")

	//输出request header
	io.WriteString(rw, "================Details of the http request header:============\n")
	for k, v := range req.Header {
		io.WriteString(rw, fmt.Sprintf("%s=%s\n", k, v))
		//接收客户端 request，并将 request 中带的 header 写入 response header

		//use Add instead of Set which will not override
		for _, vv := range v {
			fmt.Printf("Header key: %s, Header value: %s \n", k, vv)
			rw.Header().Add(k, vv)
		}
	}

	//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	rw.Header().Add("Version", version)

	//输出 response header
	io.WriteString(rw, "================Details of the http response header:============\n")
	for k, v := range rw.Header() {
		io.WriteString(rw, fmt.Sprintf("%s=%s\n", k, v))
	}

	//Server 端记录访问日志包括客户端 IP，输出到 server 端的标准输出
	glog.Info("Client IP is ", utilities.GetClientIP(req))
	glog.Info("URL Path is ", req.URL.Path) // how to prevent favicon.ico request url when browser access?
}
