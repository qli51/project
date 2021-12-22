package client

import (
	"context"
	"fmt"
	"io/ioutil"

	newHttp "net/http"

	"github.com/tal-tech/go-zero/core/logx"

	"common/http"
)

func RequestData(userInfo []byte) (string, error) {
	var hostName, osName, startTime, endTime string
	params := map[string]interface{}{}

	fmt.Println("Please enter your hostName:")
	fmt.Scanln(&hostName)
	params["hostName"] = hostName

	fmt.Println("Please enter your osName:")
	fmt.Scanln(&osName)
	params["osName"] = osName

	fmt.Println("Please enter your startTime:")
	fmt.Scanln(&startTime)
	params["startTime"] = startTime

	fmt.Println("Please enter your endTime:")
	fmt.Scanln(&endTime)
	params["endTime"] = endTime

	header := newHttp.Header{"Authorization": []string{"APPCODE 666"}}
	client := http.NewHttpClient(context.Background(), header)
	url := client.BuildUrl("127.0.0.1", "9090", "/monitorcenter/data/trend", params)
	resp, err := client.Request(url, "get", userInfo)
	if err != nil {
		logx.Errorf("get data failed: %s", err)
		return "", err
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logx.Errorf("get data failed: %s", err)
		return "", err
	}

	return string(res), nil
}
