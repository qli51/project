package server

import (
	"encoding/json"
	"errors"

	"shop/common"
	"shop/mysql"

	"github.com/tal-tech/go-zero/core/logx"
)

func checkBalance(params *common.CheckRequestParams) ([]byte, error) {
	db, err := mysql.NewDB()
	if err != nil {
		logx.Errorf("create mysql failed: %s", err)
		return nil, err
	}
	defer db.DB.Close()

	userBalance, err := db.QueryUserInfo(params.ID)
	if err != nil {
		logx.Errorf("get balance failed: %s", err)
		return nil, err
	}

	if len(userBalance) == 0 {
		logx.Errorf("get user(%s) balance failed!", params.ID)
		return nil, errors.New("match balance in mysql failed")
	}

	data, err := json.Marshal(userBalance)
	if err != nil {
		logx.Errorf("parse balance failed: %s", err)
		return nil, err
	}

	return data, nil
}

func checkOrderList(params *common.CheckRequestParams) ([]byte, error) {
	db, err := mysql.NewDB()
	if err != nil {
		logx.Errorf("create mysql failed: %s", err)
		return nil, err
	}
	defer db.DB.Close()

	orderList, err := db.QueryOrderInfo(params.ID)
	if err != nil {
		logx.Errorf("get orderList failed: %s", err)
		return nil, err
	}
	if len(orderList) == 0 {
		logx.Errorf("get user(%s) order list failed!", params.ID)
		return nil, errors.New("match order list in mysql failed")
	}

	data, err := json.Marshal(orderList)
	if err != nil {
		logx.Errorf("parse orderList failed: %s", err)
		return nil, err
	}

	return data, nil
}

func checkShopList(params *common.CheckRequestParams) ([]byte, error) {
	db, err := mysql.NewDB()
	if err != nil {
		logx.Errorf("create mysql failed: %s", err)
		return nil, err
	}
	defer db.DB.Close()

	shopList, err := db.QueryProductInfo("all")
	if err != nil {
		logx.Errorf("get shopList failed: %s", err)
		return nil, err
	}
	if len(shopList) == 0 {
		logx.Errorf("get shop list failed!")
		return nil, errors.New("match shop list in mysql failed")
	}

	data, err := json.Marshal(shopList)
	if err != nil {
		logx.Errorf("parse orderList failed: %s", err)
		return nil, err
	}

	return data, nil
}
