package server

import (
	"errors"
	"strconv"
	"time"

	"shop/common"
	"shop/mysql"

	"github.com/tal-tech/go-zero/core/logx"
)

func orderProduct(params *common.OrderRequestParams) error {
	db, err := mysql.NewDB()
	if err != nil {
		logx.Errorf("create mysql failed: %s", err)
		return err
	}
	defer db.DB.Close()

	// 查看商品列表
	shopList, err := db.QueryProductInfo(params.ShopID)
	if err != nil {
		logx.Errorf("get product failed: %s", err)
		return err
	}
	if len(shopList) == 0 {
		logx.Errorf("get shop products failed!")
		return errors.New("match shop products failed")
	}

	// 获取余额
	userInfo, err := db.QueryUserInfo(params.ID)
	if err != nil {
		logx.Errorf("get balance failed: %s", err)
		return nil
	}
	if len(userInfo) == 0 {
		logx.Errorf("get user(%s) balance failed!", params.ID)
		return errors.New("match balance failed")
	}

	balance, _ := strconv.ParseFloat(userInfo[0]["balance"], 64)
	price, _ := strconv.ParseFloat(shopList[0]["price"], 64)
	newBalance := balance - price

	// 判断余额是否充足
	if newBalance < 0 {
		return errors.New("blance is not enough, buy product failed")
	}

	// 更新余额
	db.UpdateBalance(newBalance, userInfo[0]["id"])

	// 执行下单
	keys := []string{"id", "product_item", "total_price", "status", "address_id", "user_id", "nick_name", "created", "updated"}
	values := []interface{}{time.Now().Unix(), params.ShopID, price, "test", "test", params.ID, "test", "test", "test"}
	db.InsertOrder(keys, values...)

	return nil
}
