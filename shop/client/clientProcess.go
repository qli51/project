package client

import (
	"common/http"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"shop/common"
	"shop/config"

	"github.com/tal-tech/go-zero/core/logx"
)

func clientLogin(userInfo *common.UserInfo) error {
	data, err := json.Marshal(userInfo)
	if err != nil {
		logx.Errorf("decode json failed: %s", err)
		return err
	}

	client := http.NewHttpClient(context.Background(), nil)
	loginHost := config.Servers.LoginServer.Host
	loginPort := config.Servers.LoginServer.Port
	url := client.BuildUrl(loginHost, loginPort, "/shop/login", nil)
	resp, err := client.Request(url, "set", data)
	if err != nil {
		logx.Errorf("get login result failed: %s", err)
		return err
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logx.Errorf("get data failed: %s", err)
		return err
	}

	if string(res) == "passwd is false!" {
		fmt.Println(string(res))
		return err
	}

	fmt.Println(string(res))
	return nil
}

func clientLogout(userInfo *common.UserInfo) error {
	data, err := json.Marshal(userInfo)
	if err != nil {
		logx.Errorf("decode json failed: %s", err)
		return err
	}

	client := http.NewHttpClient(context.Background(), nil)
	logoutHost := config.Servers.LoginServer.Host
	logoutPort := config.Servers.LoginServer.Port
	url := client.BuildUrl(logoutHost, logoutPort, "/shop/logout", nil)
	resp, err := client.Request(url, "set", data)
	if err != nil {
		logx.Errorf("get login result failed: %s", err)
		return err
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logx.Errorf("get data failed: %s", err)
		return err
	}

	fmt.Println(string(res))
	return nil
}

func clientCheck(userInfo *common.UserInfo, params map[string]interface{}, checkType string) error {
	data, err := json.Marshal(userInfo)
	if err != nil {
		logx.Errorf("decode json failed: %s", err)
		return err
	}

	client := http.NewHttpClient(context.Background(), nil)
	shopHost := config.Servers.ShopServer.Host
	shopPort := config.Servers.ShopServer.Port

	var url string
	switch checkType {
	case "balance":
		url = client.BuildUrl(shopHost, shopPort, "/shop/data/check/balance", params)
	case "shopList":
		url = client.BuildUrl(shopHost, shopPort, "/shop/data/check/shopList", params)
	case "orderList":
		url = client.BuildUrl(shopHost, shopPort, "/shop/data/check/orderList", params)
	default:
		return errors.New("unknown type")
	}

	resp, err := client.Request(url, "get", data)
	if err != nil {
		logx.Errorf("get data check result failed: %s", err)
		return err
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logx.Errorf("get data failed: %s", err)
		return err
	}

	fmt.Println(string(res))
	return nil
}

func clientRecharge(userInfo *common.UserInfo, params map[string]interface{}) error {
	data, err := json.Marshal(userInfo)
	if err != nil {
		logx.Errorf("decode json failed: %s", err)
		return err
	}

	client := http.NewHttpClient(context.Background(), nil)
	shopHost := config.Servers.ShopServer.Host
	shopPort := config.Servers.ShopServer.Port
	url := client.BuildUrl(shopHost, shopPort, "/shop/data/recharge", params)
	resp, err := client.Request(url, "set", data)
	if err != nil {
		logx.Errorf("get data recharge result failed: %s", err)
		return err
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logx.Errorf("get data failed: %s", err)
		return err
	}

	fmt.Println(string(res))
	return nil
}

func clientOrder(userInfo *common.UserInfo, params map[string]interface{}) error {
	data, err := json.Marshal(userInfo)
	if err != nil {
		logx.Errorf("decode json failed: %s", err)
		return err
	}

	client := http.NewHttpClient(context.Background(), nil)
	shopHost := config.Servers.ShopServer.Host
	shopPort := config.Servers.ShopServer.Port
	url := client.BuildUrl(shopHost, shopPort, "/shop/data/order", params)
	resp, err := client.Request(url, "set", data)
	if err != nil {
		logx.Errorf("get data order result failed: %s", err)
		return err
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logx.Errorf("get data failed: %s", err)
		return err
	}

	fmt.Println(string(res))
	return nil
}
