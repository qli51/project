package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"shop/common"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest/httpx"
)

var UserReocrd sync.Map

func parseUserInfo(request *http.Request, record *common.UserInfo) error {
	userInfo, err := ioutil.ReadAll(request.Body)
	if err != nil {
		logx.Errorf("get usrInfo failed: %s", err)
		return err
	}

	err = json.Unmarshal(userInfo, record)
	if err != nil {
		logx.Errorf("decode failed: %s", err)
		return err
	}

	return nil
}

// preChecking 中间键，用于检测用户的有效性
func preChecking(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var record common.UserInfo
		err := parseUserInfo(r, &record)
		if err != nil {
			logx.Errorf("get record information failed: %s", err)
			return
		}

		// 为确保只有最后一次登录有效，需要检测登录的时间
		if recordTime, ok := UserReocrd.Load(record.Name); ok {
			if recordTime.(string) != record.LoginTime {
				w.Write([]byte("user has been log out!"))
				return
			}
		}

		f(w, r)
	}
}

func serverLogin(writer http.ResponseWriter, request *http.Request) {
	var record common.UserInfo
	err := parseUserInfo(request, &record)
	if err != nil {
		logx.Errorf("get record information failed: %s", err)
		return
	}

	err = login(&record)

	if err != nil {
		writer.Write([]byte("passwd is false!"))
		return
	}

	writer.Write([]byte("password right, login succeed!"))
}

func serverLogout(writer http.ResponseWriter, request *http.Request) {
	var record common.UserInfo
	err := parseUserInfo(request, &record)
	if err != nil {
		logx.Errorf("get record information failed: %s", err)
		return
	}

	err = logout(&record)
	if err != nil {
		writer.Write([]byte("log out failed!"))
	}

	writer.Write([]byte("log out success!"))
}

func serverGetBalance(writer http.ResponseWriter, request *http.Request) {
	params := common.CheckRequestParams{}
	if err := httpx.Parse(request, &params); err != nil {
		logx.Errorf("get data failed: %s", err)
		return
	}

	data, err := checkBalance(&params)
	if err != nil {
		failedStr := fmt.Sprintf("check balance failed: %s!", err)
		writer.Write([]byte(failedStr))
	}

	writer.Write(data)
}

func serverGetOrderList(writer http.ResponseWriter, request *http.Request) {
	params := common.CheckRequestParams{}
	if err := httpx.Parse(request, &params); err != nil {
		logx.Errorf("get data failed: %s", err)
		return
	}

	data, err := checkOrderList(&params)
	if err != nil {
		failedStr := fmt.Sprintf("check orderList failed: %s!", err)
		writer.Write([]byte(failedStr))
	}

	writer.Write(data)
}

func serverGetShopList(writer http.ResponseWriter, request *http.Request) {
	params := common.CheckRequestParams{}
	if err := httpx.Parse(request, &params); err != nil {
		logx.Errorf("get data failed: %s", err)
		return
	}

	data, err := checkShopList(&params)
	if err != nil {
		failedStr := fmt.Sprintf("check shopList failed: %s!", err)
		writer.Write([]byte(failedStr))
	}

	writer.Write(data)
}

func serverOrder(writer http.ResponseWriter, request *http.Request) {
	params := common.OrderRequestParams{}
	if err := httpx.Parse(request, &params); err != nil {
		logx.Errorf("get data failed: %s", err)
		return
	}

	err := orderProduct(&params)
	if err != nil {
		failedStr := fmt.Sprintf("order failed: %s!", err)
		writer.Write([]byte(failedStr))
	}

	writer.Write([]byte("order success!"))
}

func serverRecharge(writer http.ResponseWriter, request *http.Request) {
	params := common.RechargeRequestParams{}
	if err := httpx.Parse(request, &params); err != nil {
		logx.Errorf("get data failed: %s", err)
		return
	}

	err := rechargeBalance(&params)
	if err != nil {
		failedStr := fmt.Sprintf("recharge failed: %s!", err)
		writer.Write([]byte(failedStr))
	}

	writer.Write([]byte("recharge success!"))
}
