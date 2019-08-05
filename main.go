package main

import (
	"fmt"

	"./e621"
)

func main() {
	baseURL := "https://e621.net"
	userName := "otyaken"
	client, _ := e621.NewClient(baseURL, userName, nil)
	post, _ := client.GetPostByID(nil, 12345)
	fmt.Printf("%v", post)
}
