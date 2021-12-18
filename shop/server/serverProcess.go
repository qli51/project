package server

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"shop/common"
	"shop/mysql"

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

func checkLogin(name, loginTime string) bool {
	// 为确保只有最后一次登录有效，需要检测登录的时间
	if recordTime, ok := UserReocrd.Load(name); ok {
		if recordTime.(string) == loginTime {
			return true
		}
	}

	return false
}

func serverLogin(writer http.ResponseWriter, request *http.Request) {
	var record common.UserInfo
	err := parseUserInfo(request, &record)
	if err != nil {
		logx.Errorf("get record information failed: %s", err)
		return
	}

	md5Coding := md5.New()
	md5Coding.Write([]byte(record.Passwd))
	userPasswd := hex.EncodeToString(md5Coding.Sum(nil))

	file, err := os.Open(common.UserInfoPATH)
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
		if lineInfoSlice == nil || len(lineInfoSlice) < 2 {
			continue
		}

		if record.Name == lineInfoSlice[0] && userPasswd == lineInfoSlice[1] {
			UserReocrd.Store(record.Name, record.LoginTime)
			writer.Write([]byte("password right, login succeed!"))
			return
		}
	}

	writer.Write([]byte("passwd is false!"))
}

func serverLogout(writer http.ResponseWriter, request *http.Request) {
	var record common.UserInfo
	err := parseUserInfo(request, &record)
	if err != nil {
		logx.Errorf("get record information failed: %s", err)
		return
	}

	if !checkLogin(record.Name, record.LoginTime) {
		writer.Write([]byte("user has been log out!"))
		return
	}

	UserReocrd.Delete(record.Name)

	writer.Write([]byte("log out success!"))
}

func serverDataCheck(writer http.ResponseWriter, request *http.Request) {
	var record common.UserInfo
	err := parseUserInfo(request, &record)
	if err != nil {
		logx.Errorf("get record information failed: %s", err)
		return
	}

	if !checkLogin(record.Name, record.LoginTime) {
		writer.Write([]byte("user has been log out!"))
		return
	}

	params := common.CheckRequestParams{}
	if err := httpx.Parse(request, &params); err != nil {
		logx.Errorf("get data failed: %s", err)
		return
	}

	db, err := mysql.NewDB()
	if err != nil {
		logx.Errorf("create mysql failed: %s", err)
		return
	}
	defer db.DB.Close()

	var mysqlCmd string

	switch params.Type {
	case "balance":
		mysqlCmd = fmt.Sprintf(`select * from Info where id="%s"`, params.ID)
		userBalance, err := db.Query(mysqlCmd)
		if err != nil {
			logx.Errorf("get balance failed: %s", err)
			return
		}
		if len(userBalance) == 0 {
			logx.Errorf("get user(%s) balance failed!", params.ID)
			return
		}

		data, err := json.Marshal(userBalance)
		if err != nil {
			logx.Errorf("parse balance failed: %s", err)
			return
		}
		writer.Write(data)

	case "orderList":
		mysqlCmd = fmt.Sprintf(`select * from orders where user_id="%s"`, params.ID)
		orderList, err := db.Query(mysqlCmd)
		if err != nil {
			logx.Errorf("get orderList failed: %s", err)
			return
		}
		if len(orderList) == 0 {
			logx.Errorf("get user(%s) order list failed!", params.ID)
			return
		}

		data, err := json.Marshal(orderList)
		if err != nil {
			logx.Errorf("parse orderList failed: %s", err)
			return
		}
		writer.Write(data)

	case "shopList":
		mysqlCmd = fmt.Sprintf("select * from product")
		shopList, err := db.Query(mysqlCmd)
		if err != nil {
			logx.Errorf("get shopList failed: %s", err)
			return
		}
		if len(shopList) == 0 {
			logx.Errorf("get shop list failed!")
			return
		}

		data, err := json.Marshal(shopList)
		if err != nil {
			logx.Errorf("parse orderList failed: %s", err)
			return
		}
		writer.Write(data)

	default:
		logx.Errorf("undefined type of check")
		return
	}
}

func serverDataOrder(writer http.ResponseWriter, request *http.Request) {
	var record common.UserInfo
	err := parseUserInfo(request, &record)
	if err != nil {
		logx.Errorf("get record information failed: %s", err)
		return
	}

	if !checkLogin(record.Name, record.LoginTime) {
		writer.Write([]byte("user has been log out!"))
		return
	}

	params := common.OrderRequestParams{}
	if err := httpx.Parse(request, &params); err != nil {
		logx.Errorf("get data failed: %s", err)
		return
	}

	db, err := mysql.NewDB()
	if err != nil {
		logx.Errorf("create mysql failed: %s", err)
		return
	}
	defer db.DB.Close()

	// 查看商品列表
	shopCmd := fmt.Sprintf("select * from product where id=%s", params.ShopID)
	shopList, err := db.Query(shopCmd)
	if err != nil {
		logx.Errorf("get product failed: %s", err)
		return
	}
	if len(shopList) == 0 {
		logx.Errorf("get shop products failed!")
		return
	}

	// 获取余额
	getBlanceCmd := fmt.Sprintf(`select * from Info where id="%s"`, params.ID)
	userBalance, err := db.Query(getBlanceCmd)
	if err != nil {
		logx.Errorf("get balance failed: %s", err)
		return
	}
	if len(userBalance) == 0 {
		logx.Errorf("get user(%s) balance failed!", params.ID)
		return
	}

	balance, _ := strconv.ParseFloat(userBalance[0]["balance"], 64)
	price, _ := strconv.ParseFloat(shopList[0]["price"], 64)
	newBalance := balance - price

	// 判断余额是否充足
	if newBalance < 0 {
		writer.Write([]byte("blance is not enough, buy product failed"))
		return
	}

	// 更新余额
	db.ExecUpdate("Info", "balance", "id", newBalance, userBalance[0]["id"])

	// 执行下单
	keys := []string{"id", "product_item", "total_price", "status", "address_id", "user_id", "nick_name", "created", "updated"}
	db.ExecInsert("orders", keys, time.Now().Unix(), params.ShopID, price, "test", "test", params.ID, "test", "test", "test")

	writer.Write([]byte("order success!"))
}

func serverDataRecharge(writer http.ResponseWriter, request *http.Request) {
	var record common.UserInfo
	err := parseUserInfo(request, &record)
	if err != nil {
		logx.Errorf("get record information failed: %s", err)
		return
	}

	if !checkLogin(record.Name, record.LoginTime) {
		writer.Write([]byte("user has been log out!"))
		return
	}

	params := common.RechargeRequestParams{}
	if err := httpx.Parse(request, &params); err != nil {
		logx.Errorf("get data failed: %s", err)
		return
	}

	getBlanceCmd := fmt.Sprintf(`select * from Info where id="%s"`, params.ID)
	db, err := mysql.NewDB()
	if err != nil {
		logx.Errorf("create mysql failed: %s", err)
		return
	}
	defer db.DB.Close()

	userInfo, err := db.Query(getBlanceCmd)
	if err != nil {
		logx.Errorf("get balance failed: %s", err)
		return
	}

	if len(userInfo) == 0 {
		logx.Errorf("get user(%s) balance failed!", params.ID)
		return
	}

	balance, _ := strconv.ParseFloat(userInfo[0]["balance"], 64)
	newBalance := balance + params.Value

	db.ExecUpdate("Info", "balance", "id", newBalance, userInfo[0]["id"])

	writer.Write([]byte("update balance success!"))
}
