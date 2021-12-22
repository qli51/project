package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"collect/collectMethod"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func dataGet(writer http.ResponseWriter, request *http.Request) {
	if !checkLogin(request) {
		logx.Errorf("need to login first")
		writer.Write([]byte("need to login first"))
		return
	}

	if !checkAuthorization(request) {
		logx.Errorf("dataGet fail to pass checkAuthorization")
		return
	}

	var params requestParams
	if err := httpx.Parse(request, &params); err != nil {
		logx.Errorf("get data failed: %s", err)
		return
	}

	dataRes, err := dataSelect(&params)
	if err != nil {
		logx.Errorf("select data failed: %s", err)
		return
	}

	writer.Write([]byte(dataRes))
}

func isDataMatch(params *requestParams, data *collectMethod.DataRecord, currentTime int64) bool {
	if params.StartTime == 0 {
		params.StartTime = 0
	}
	if params.EndTime == 0 {
		params.EndTime = currentTime
	}
	if params.StartTime > params.EndTime {
		return false
	}

	collectTime, err := strconv.ParseInt(data.Time, 10, 64)
	if err != nil {
		logx.Errorf("transfer to int64 failed: %s", err)
		return false
	}

	if collectTime < params.StartTime || collectTime > params.EndTime {
		return false
	}

	if params.HostName != "" && params.HostName != data.HostName {
		return false
	}

	if params.OSName != "" && params.OSName != data.OSName {
		return false
	}

	return true
}

func dataSelect(params *requestParams) (string, error) {
	lock.Lock()
	defer lock.Unlock()
	dateStr := ""

	file, err := os.Open(collectDataPATH)
	if err != nil {
		logx.Errorf("open file failed: %s", err)
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		lineInfo, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		data := &collectMethod.DataRecord{}
		err = json.Unmarshal(lineInfo, data)
		if err != nil {
			continue
		}

		if isDataMatch(params, data, time.Now().Unix()) {
			dateStr = fmt.Sprintf("%s\n%s", string(lineInfo), dateStr)
		}
	}

	return dateStr, nil
}
