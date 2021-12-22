package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"collect/common"

	"github.com/tal-tech/go-zero/core/logx"
)

// ���ýӿ���ҪУ���û�����
func checkLogin(request *http.Request) bool {
	userInfo, err := ioutil.ReadAll(request.Body)
	if err != nil {
		logx.Errorf("get usrInfo failed: %s", err)
		return false
	}

	var record common.UserInfo
	err = json.Unmarshal(userInfo, &record)
	if err != nil {
		logx.Errorf("decode failed: %s", err)
		return false
	}

	// У���¼ʱ�䣬ȷ��ֻ�����µĵ�¼��Ч
	if recordTime, ok := UserReocrd.Load(record.Name); ok {
		if recordTime.(string) == record.LoginTime {
			return true
		}
	}

	return false
}

// ���ýӿ���Ҫ��Ȩ
func checkAuthorization(request *http.Request) bool {
	authorizationInfo := request.Header.Get("Authorization")
	if authorizationInfo == "" {
		return false
	}
	code := strings.Fields(authorizationInfo)[1]
	if code != authorizationCode {
		return false
	}
	return true
}
