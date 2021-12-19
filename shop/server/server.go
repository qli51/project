package server

import (
	"net/http"
	"fmt"
	"sync"

	"shop/config"

	"github.com/tal-tech/go-zero/core/threading"
)

const (
	// api数目
	interfaceNum = 3
)

func startListenLogin() {
	loginHost := config.Servers.LoginServer.Host
	loginPort := config.Servers.LoginServer.Port
	http.HandleFunc("/shop/login", serverLogin)
	server := fmt.Sprintf("%s:%s", loginHost, loginPort)
	http.ListenAndServe(server, nil)
}

func startListenLogout() {
	logoutHost := config.Servers.LogOutServer.Host
	logoutPort := config.Servers.LogOutServer.Port
	http.HandleFunc("/shop/logout", serverLogout)
	server := fmt.Sprintf("%s:%s", logoutHost, logoutPort)
	http.ListenAndServe(server, nil)
}

func startListenShopProcess() {
	shopHost := config.Servers.ShopServer.Host
	shopPort := config.Servers.ShopServer.Port
	http.HandleFunc("/shop/data/check/balance", logCheck(serverCheckBalance))
	http.HandleFunc("/shop/data/check/orderList", logCheck(serverCheckOrderList))
	http.HandleFunc("/shop/data/check/shopList", logCheck(serverCheckShopList))
	http.HandleFunc("/shop/data/order", logCheck(serverDataOrder))
	http.HandleFunc("/shop/data/recharge", logCheck(serverDataRecharge))
	server := fmt.Sprintf("%s:%s", shopHost, shopPort)
	http.ListenAndServe(server, nil)
}

func startServer() {
	pool := threading.NewTaskRunner(interfaceNum)
	var wg sync.WaitGroup
	wg.Add(interfaceNum)

	pool.Schedule(func() {
		defer wg.Done()
		startListenLogin()
	})

	pool.Schedule(func() {
		defer wg.Done()
		startListenShopProcess()
	})

	pool.Schedule(func() {
		defer wg.Done()
		startListenLogout()
	})

	wg.Wait()

}

func Run() {
	config.LoadMetricsConf()
	startServer()
}
