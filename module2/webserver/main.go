package main

import (
	"fmt"
	"github.com/lvxiaohu/cloudnative-stu/module2/webserver/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

//接收客户端 request，并将 request 中带的 header 写入 response header
//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
//当访问 localhost/healthz 时，应返回 200

func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("AppName", "K8S")
	for k, v := range r.Header {
		w.Header().Add(k, fmt.Sprint(v))
		//fmt.Println(k, v)
	}
	// 获取客户端IP
	clientIP := r.Header.Get("X-Real-Ip")
	if clientIP == "" {
		clientIP = strings.Split(r.RemoteAddr, ":")[0]
	}
	log.Printf("客户端IP是: %s, 状态：%d", clientIP, 200)
}

// ClientIP 根据不同的网络情况可以获取IP，例如前端加了Nginx、Haproxy
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err != nil {
		return ip
	}
	return ""
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// 给 Response 写入数据
	io.WriteString(w, "200")

}

func GetOSENV(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, os.Getenv("VERSION"))
}

func images(w http.ResponseWriter, r *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	randInt := rand.Intn(2000)
	time.Sleep(time.Millisecond * time.Duration(randInt))
	w.Write([]byte(fmt.Sprintf("<h1>%d<h1>", randInt)))
}

func main() {
	// 给系统设置一个 VERSION 变量
	os.Setenv("VERSION", "Linux-4.5")

	http.HandleFunc("/", root)
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/version", GetOSENV)
	http.HandleFunc("/images", images)
	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Web Server failed!")
	}
}
