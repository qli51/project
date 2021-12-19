package client

import (
	"fmt"
	"time"
	"strconv"
	"shop/config"
	"shop/common"

	"github.com/tal-tech/go-zero/core/logx"
)

var userInfo common.UserInfo

func loginProcess() error {
	fmt.Println("please login first")
	var name, passwd string

	fmt.Println("Please enter your account:")
	fmt.Scanln(&name)
	userInfo.Name = name

	fmt.Println("Please enter your passwd:")
	fmt.Scanln(&passwd)
	userInfo.Passwd = passwd

	userInfo.LoginTime = strconv.FormatInt(time.Now().Unix(), 10)

	err := clientLogin(&userInfo)
	if err != nil {
		return err
	}

	return nil
}

func logoutProcess() error {
	err := clientLogout(&userInfo)
	if err != nil {
		logx.Errorf("log out failed: %s", err)
		return err
	}

	return nil
}

func checkProcess() error {
	params := make(map[string]interface{}, 1)
    
	var id, checkType string
	fmt.Println("Please enter your ID of checking:")
	fmt.Scanln(&id)
	params["id"] = id

	fmt.Println("Please enter your type of checking:")
	fmt.Scanln(&checkType)

	err := clientCheck(&userInfo, params, checkType)
	if err != nil {
		return err
	}

	return nil
}

func rechargeProcess() error {
	params := make(map[string]interface{}, 2)

	var id, value string
	fmt.Println("Please enter your ID of recharge:")
	fmt.Scanln(&id)
	params["id"] = id

	fmt.Println("Please enter your value of recharge:")
	fmt.Scanln(&value)
	params["value"] = value


	err := clientRecharge(&userInfo, params)
	if err != nil {
		return err
	}

	return nil
}

func orderProcess() error {
	params := make(map[string]interface{}, 2)

	var id, shopID string
	fmt.Println("Please enter your userID of order:")
	fmt.Scanln(&id)
	params["id"] = id

	fmt.Println("Please enter your shopID of order:")
	fmt.Scanln(&shopID)
	params["shopID"] = shopID


	err := clientOrder(&userInfo, params)
	if err != nil {
		return err
	}

	return nil
}

func Run() {
	config.LoadMetricsConf()
	err := loginProcess()
	if err != nil{
		return
	}

	for {
		fmt.Println("please enter your operate:")
		var operate string
		fmt.Scanln(&operate)
		switch operate {

		case "logout":
			err := logoutProcess()
			if err != nil {
				logx.Errorf("log out failed: $err", err)
				continue
			}
			return

		case "check":
			err := checkProcess()
			if err != nil {
				logx.Errorf("check process failed: $err", err)
				continue
			}
			continue

		case "recharge":
			err := rechargeProcess()
			if err != nil {
				logx.Errorf("recharge failed: $err", err)
				continue
			}
			continue


		case "order":
			err := orderProcess()
			if err != nil {
				logx.Errorf("order failed: $err", err)
				continue
			}
			continue

		default:
			fmt.Println("operate should be logout,check,recharge or order")
			continue
		}
	}
}
