package main

import (
	"fmt"

	"collect/client"

	"github.com/tal-tech/go-zero/core/logx"
)

func main() {
	userRecord, err := client.LoginProcess()
	if err != nil {
		logx.Errorf("login failed: %s", err)
		return
	}

	fmt.Println("login succeed!")

	res, err := client.RequestData(userRecord)
	if err != nil {
		logx.Errorf("request data failed: %s", err)
		return
	}

	fmt.Println(res)
}
