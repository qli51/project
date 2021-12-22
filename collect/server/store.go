package server

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/tal-tech/go-zero/core/logx"
)

func store(collectData []byte) error {
	lock.Lock()
	defer lock.Unlock()
	file, err := os.OpenFile(collectDataPATH, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logx.Errorf("open file failed: %s", err)
	}
	write := bufio.NewWriter(file)
	write.WriteString(fmt.Sprintf("%s\n", string(collectData)))
	write.Flush()

	return nil
}

func dataStore(writer http.ResponseWriter, request *http.Request) {
	if !checkAuthorization(request) {
		logx.Errorf("dataStore fail to pass checkAuthorization")
		return
	}

	collectData, err := ioutil.ReadAll(request.Body)
	if err != nil {
		logx.Errorf("get data failed: $err", err)
	}

	store(collectData)

	writer.Write([]byte("store data succeed"))
}
