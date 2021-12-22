package server

import (
	"sync"
)

// lock 数据操作需要加锁，防止并发导致的数据损坏
var lock = &sync.RWMutex{}
var UserReocrd sync.Map
var name, passwd string

const (
	// 数据信息存储路径
	collectDataPATH = "./dataRecord/collect_data"
	// 账户信息存储路径
	userInfoPATH = "./dataRecord/user_info"
	// 提供的接口数量
	interfaceNum = 3
	// 只有鉴权码为666，才可以通过鉴权
	authorizationCode = "666"
)

type requestParams struct {
	HostName  string `form:"hostName,optional"`
	OSName    string `form:"osName,optional"`
	StartTime int64  `form:"startTime,optional"`
	EndTime   int64  `form:"endTime,optional"`
}