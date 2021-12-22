package server

import (
	"net/http"
	"sync"
	
	"github.com/tal-tech/go-zero/core/threading"
)

func startListenLogin() {
	http.HandleFunc("/monitorcenter/login", login)
	http.ListenAndServe("0.0.0.0:7070", nil)
}

func startListenDataCollect() {
	http.HandleFunc("/monitorcenter/data/store", dataStore)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func startListenDataGet() {
	http.HandleFunc("/monitorcenter/data/trend", dataGet)
	http.ListenAndServe("0.0.0.0:9090", nil)
}

func StartServer() {
	pool := threading.NewTaskRunner(interfaceNum)
	var wg sync.WaitGroup
	wg.Add(interfaceNum)

	pool.Schedule(func() {
		defer wg.Done()
		startListenLogin()
	})

	pool.Schedule(func() {
		defer wg.Done()
		startListenDataCollect()
	})

	pool.Schedule(func() {
		defer wg.Done()
		startListenDataGet()
	})

	wg.Wait()
}
