package server

import (
	"fmt"
	"net/http"
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
	http.HandleFunc("/shop/data/balance", preChecking(serverGetBalance))
	http.HandleFunc("/shop/data/orderList", preChecking(serverGetOrderList))
	http.HandleFunc("/shop/data/shopList", preChecking(serverGetShopList))
	http.HandleFunc("/shop/order", preChecking(serverOrder))
	http.HandleFunc("/shop/recharge", preChecking(serverRecharge))
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
