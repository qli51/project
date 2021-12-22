package server

import (
	"errors"
	"strconv"

	"shop/common"
	"shop/mysql"

	"github.com/tal-tech/go-zero/core/logx"
)

func rechargeBalance(params *common.RechargeRequestParams) error {
	db, err := mysql.NewDB()
	if err != nil {
		logx.Errorf("create mysql failed: %s", err)
		return nil
	}
	defer db.DB.Close()

	userInfo, err := db.QueryUserInfo(params.ID)
	if err != nil {
		logx.Errorf("get balance failed: %s", err)
		return nil
	}

	if len(userInfo) == 0 {
		logx.Errorf("get user(%s) balance failed!", params.ID)
		return errors.New("match userInfo failed")
	}

	balance, _ := strconv.ParseFloat(userInfo[0]["balance"], 64)
	newBalance := balance + params.Value

	db.UpdateBalance(newBalance, userInfo[0]["id"])

	return nil
}
