package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/golang/glog"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func wrapHandlerWithLogging(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		lrw := NewLoggingResponseWriter(w)
		wrappedHandler.ServeHTTP(lrw, req)

		statusCode := lrw.statusCode
		//输出 status code
		glog.Info("HTTP Status code ", statusCode)
	})
}

func main() {
	glog.Info("Starting http server")

	roothandler := wrapHandlerWithLogging(http.HandlerFunc(rootHandler))
	healthzhandler := wrapHandlerWithLogging(http.HandlerFunc(healthz))

	http.Handle("/", roothandler)
	http.Handle("/healthz", healthzhandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func healthz(w http.ResponseWriter, r *http.Request) {
	//当访问 localhost/healthz 时，应返回 200
	glog.Info("Entering healthz handler")
	io.WriteString(w, "ok\n")
}

func rootHandler(rw http.ResponseWriter, req *http.Request) {
	glog.Info("Entering root handler")

	// fmt.Println(os.Environ())
	term_program_version := os.Getenv("TERM_PROGRAM_VERSION")

	//输出request header
	io.WriteString(rw, "================Details of the http request header:============\n")
	for k, v := range req.Header {
		io.WriteString(rw, fmt.Sprintf("%s=%s\n", k, v))
		//接收客户端 request，并将 request 中带的 header 写入 response header
		rw.Header().Add(k, v[0])
	}

	//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	rw.Header().Add("Term_program_version", term_program_version)

	//输出 response header
	io.WriteString(rw, "================Details of the http response header:============\n")
	for k, v := range rw.Header() {
		io.WriteString(rw, fmt.Sprintf("%s=%s\n", k, v))
	}

	//Server 端记录访问日志包括客户端 IP，输出到 server 端的标准输出
	glog.Info("Client IP is ", req.RemoteAddr)
	glog.Info("URL Path is ", req.URL.Path) // how to prevent favicon.ico request url?
}
