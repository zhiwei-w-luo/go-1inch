package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	go1inch "github.com/jon4hz/go-1inch"
)

func main() {
	ordi()
}

func ordi(){
	client := go1inch.NewClient()
	res, _, err := client.ApproveTransaction(context.Background(), go1inch.ZkSync, "0x5e38cb3e6c0faafaa5c32c482864fcef5a0660ad", &go1inch.ApproveTransactionOpts{
		Amount: "5000000",
	})
	if err != nil {
		log.Fatal("Error while getting an approve calldata ", err)
	}
	fmt.Println(res)

	fmt.Println(client.Tokens(context.Background(), "zkSync"))
}

func explorer() {
	client := &http.Client{}

	url := "https://api.1inch.io/v5.0/324/approve/transaction?tokenAddress=0x3355df6d4c9c3035724fd0e3914de96a5a83aaf4&amount=5000000"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}



	// 设置请求头，模拟浏览器行为
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	var  expRes interface{}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Request Error:%v\n.",err)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		if expRes != nil {
			err = json.Unmarshal(body, expRes)
			if err != nil {
				fmt.Printf("Request Error:%v\n.",err)
			}
		}
		fmt.Printf("Request completed successfully.Code:%d\n", resp.StatusCode)
		fmt.Printf("Request resp.Header:%v\n",resp.Header)
		return

	default:
		fmt.Printf("%s\n", body)
		return
	}


	// 处理响应结果
	// ...


}
