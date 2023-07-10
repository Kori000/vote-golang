package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func vote(id int, a chan int) {
	for i := 0; i <= 10; i++ {

		// 构造请求参数
		requestData := map[string]int{"id": id}
		requestDataBytes, err := json.Marshal(requestData)
		if err != nil {
			fmt.Printf("requestDataBytes: %v\n", "构造请求参数失败")
			// 处理错误
		}
		req, err := http.NewRequest("POST", "http://127.0.0.1:4001/api/vote", bytes.NewBuffer(requestDataBytes))

		if err != nil {
			fmt.Printf("requestDataBytes: %v\n", "http POST请求失败")
			// 处理错误
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("requestDataBytes: %v\n", "http POST请求失败")
			// 处理错误
		}
		defer resp.Body.Close()
	}
	a <- 0
}

func main() {
	var id int

	// 读取用户输入的 id 值
	fmt.Printf("输入id: ")
	fmt.Scan(&id)
	a := make(chan int, 30)
	for i := 0; i < 30; i++ {
		go vote(id, a)
	}
	for b := range a {
		fmt.Println(b)
	}
}
