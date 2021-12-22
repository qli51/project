package server

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"collect/common"

	"github.com/tal-tech/go-zero/core/logx"
)

func login(writer http.ResponseWriter, request *http.Request) {
	userInfo, err := ioutil.ReadAll(request.Body)
	if err != nil {
		logx.Errorf("get usrInfo failed: $err", err)
	}

	var record common.UserInfo
	err = json.Unmarshal(userInfo, &record)
	if err != nil {
		logx.Errorf("decode failed: %s", err)
		return
	}

	md5Coding := md5.New()
	md5Coding.Write([]byte(record.Passwd))
	userPasswd := hex.EncodeToString(md5Coding.Sum(nil))

	file, err := os.Open(userInfoPATH)
	if err != nil {
		logx.Errorf("open file failed: %s", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		lineInfo, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		lineInfoSlice := strings.Split(string(lineInfo), ":")
		if lineInfoSlice == nil || len(lineInfoSlice) == 0 {
			continue
		}

		if record.Name == lineInfoSlice[0] && userPasswd == lineInfoSlice[1] {
			UserReocrd.Store(record.Name, record.LoginTime)
			writer.Write([]byte("1"))
			return
		}
	}
}

