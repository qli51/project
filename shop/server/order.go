package server

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"shop/common"
	"shop/mysql"

	"github.com/tal-tech/go-zero/core/logx"
)

func dataOrder(params *common.OrderRequestParams) error {
	db, err := mysql.NewDB()
	if err != nil {
		logx.Errorf("create mysql failed: %s", err)
		return err
	}
	defer db.DB.Close()

	// 查看商品列表
	shopCmd := fmt.Sprintf("select * from product where id=%s", params.ShopID)
	shopList, err := db.Query(shopCmd)
	if err != nil {
		logx.Errorf("get product failed: %s", err)
		return err
	}
	if len(shopList) == 0 {
		logx.Errorf("get shop products failed!")
		return errors.New("match shop products failed")
	}

	// 获取余额
	getBlanceCmd := fmt.Sprintf(`select * from Info where id="%s"`, params.ID)
	userBalance, err := db.Query(getBlanceCmd)
	if err != nil {
		logx.Errorf("get balance failed: %s", err)
		return nil
	}
	if len(userBalance) == 0 {
		logx.Errorf("get user(%s) balance failed!", params.ID)
		return errors.New("match balance failed")
	}

	balance, _ := strconv.ParseFloat(userBalance[0]["balance"], 64)
	price, _ := strconv.ParseFloat(shopList[0]["price"], 64)
	newBalance := balance - price

	// 判断余额是否充足
	if newBalance < 0 {
		return errors.New("blance is not enough, buy product failed")
	}

	// 更新余额
	db.ExecUpdate("Info", "balance", "id", newBalance, userBalance[0]["id"])

	// 执行下单
	keys := []string{"id", "product_item", "total_price", "status", "address_id", "user_id", "nick_name", "created", "updated"}
	db.ExecInsert("orders", keys, time.Now().Unix(), params.ShopID, price, "test", "test", params.ID, "test", "test", "test")

	return nil
}
