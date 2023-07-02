package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	go1inch "github.com/jon4hz/go-1inch"
)

func main() {
	explorer()
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

	fmt.Println(client.Tokens(context.Background(), "matic"))
}

func explorer() {
	client := go1inch.NewClient()
	res, _, err := client.ApproveTransaction(context.Background(), go1inch.ZkSync, "0x5e38cb3e6c0faafaa5c32c482864fcef5a0660ad", &go1inch.ApproveTransactionOpts{
		Amount: "5000000",
	})
	if err != nil {
		log.Fatal("Error while getting an approve calldata ", err)
	}
	fmt.Println(res)
	fmt.Println(client.Tokens(context.Background(), "zkSync"))
	fmt.Println("Request completed successfully.")
}
