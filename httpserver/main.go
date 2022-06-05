package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/golang/glog"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	// WriteHeader(int) is not called if our response implicitly returns 200 OK, so
	// we default to that status code.
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
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

func getClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
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
	glog.Info("Client IP is ", getClientIP(req))
	glog.Info("URL Path is ", req.URL.Path) // how to prevent favicon.ico request url?
}
