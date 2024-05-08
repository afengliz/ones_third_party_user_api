package main

import (
	_ "embed"
	"net/http"
)

//go:embed third_user_data.json
var thirdPartyUserData []byte

// http server
func main() {
	// 注册路由
	http.HandleFunc("/api/sync/member/1000/depart/50", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(thirdPartyUserData)
	})
	// 监听9100端口
	_ = http.ListenAndServe("127.0.0.1:9100", http.DefaultServeMux)
}
