package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received Request:")
		fmt.Println("Method:", r.Method)
		fmt.Println("URL:", r.URL)
		fmt.Println("Headers:", r.Header)

		// 读取请求主体并输出
		buf := make([]byte, 1024)
		for {
			n, err := r.Body.Read(buf)
			if err != nil {
				break
			}
			fmt.Print(string(buf[:n]))
		}
		fmt.Println("\n--- End of Request ---")

		// 发送响应
		response := "Hello, World!"
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	})

	// 启动HTTP服务器
	serverAddr := "localhost:8080"
	fmt.Printf("Server started. Listening on %s...\n", serverAddr)
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		fmt.Println("Error:", err)
	}
}
