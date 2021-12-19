package server

import (
	"errors"
	"fmt"
	"strconv"

	"shop/mysql"
	"shop/common"

	"github.com/tal-tech/go-zero/core/logx"
)

func dataRecharge(params *common.RechargeRequestParams) error {
	getBlanceCmd := fmt.Sprintf(`select * from Info where id="%s"`, params.ID)
	db, err := mysql.NewDB()
	if err != nil {
		logx.Errorf("create mysql failed: %s", err)
		return nil
	}
	defer db.DB.Close()

	userInfo, err := db.Query(getBlanceCmd)
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

	db.ExecUpdate("Info", "balance", "id", newBalance, userInfo[0]["id"])

	return nil
}
