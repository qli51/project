package client

import (
	"context"
	"encoding/json"
	"time"
	newHttp "net/http"
	"fmt"
	"io/ioutil"

	"github.com/tal-tech/go-zero/core/logx"

	"collect/collectMethod"
	"common/http"
)

func formatData() ([]byte, error) {
	collector := collectMethod.NewCollector()
	collector.CollectData()

	data, err := json.Marshal(collector.DataRecord)
	if err != nil {
		logx.Errorf("decode to json failed: %s", err)
		return nil, err
	}

	return data, nil
}

func StartCollectClient() {
	header := newHttp.Header{"Authorization" : []string{"APPCODE 666"}}
	client := http.NewHttpClient(context.Background(), header)
	url := client.BuildUrl("127.0.0.1", "8080", "/monitorcenter/data/store", nil)

	// 创建定时打点器
	tk := time.NewTicker(time.Minute)
	defer tk.Stop()

	for {
		collectData, err := formatData()
		if err != nil {
			logx.Errorf("collect data failed: %s", err)
			return
		}

		resp, err := client.Request(url, "set", collectData)
		if err != nil {
			client.Logger.Errorf("create request failed: %s", err)
			return
		}

		res, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logx.Errorf("get data failed: %s", err)
		}

		if string(res) == "need to login first" {
			fmt.Println("need to login first")
			return
		}

		logx.Infof("collect data succeed")

		<-tk.C
	}
}
