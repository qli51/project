package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/tal-tech/go-zero/core/logx"

	"collect/common"
	"common/http"
)

func LoginProcess() ([]byte, error) {
	fmt.Println("Please login first!")
	dataReocrd := &common.UserInfo{}
	var name, passwd string

	fmt.Println("Please enter your account:")
	fmt.Scanln(&name)
	dataReocrd.Name = name

	fmt.Println("Please enter your passwd:")
	fmt.Scanln(&passwd)
	dataReocrd.Passwd = passwd

	dataReocrd.LoginTime = strconv.FormatInt(time.Now().Unix(), 10)

	data, err := json.Marshal(dataReocrd)
	if err != nil {
		logx.Errorf("decode json failed: %s", err)
		return nil, err
	}

	client := http.NewHttpClient(context.Background(), nil)
	url := client.BuildUrl("127.0.0.1", "7070", "/monitorcenter/login", nil)
	resp, err := client.Request(url, "set", data)
	if err != nil {
		logx.Errorf("get login result failed: %s", err)
		return nil, err
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logx.Errorf("get data failed: %s", err)
		return nil, err
	}

	if string(res) != "1" {
		return nil, errors.New("login failed, try again!")
	}

	return data, nil
}
