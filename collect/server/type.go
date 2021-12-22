package server

import (
	"sync"
)

// lock ���ݲ�����Ҫ��������ֹ�������µ�������
var lock = &sync.RWMutex{}
var UserReocrd sync.Map
var name, passwd string

const (
	// ������Ϣ�洢·��
	collectDataPATH = "./dataRecord/collect_data"
	// �˻���Ϣ�洢·��
	userInfoPATH = "./dataRecord/user_info"
	// �ṩ�Ľӿ�����
	interfaceNum = 3
	// ֻ�м�Ȩ��Ϊ666���ſ���ͨ����Ȩ
	authorizationCode = "666"
)

type requestParams struct {
	HostName  string `form:"hostName,optional"`
	OSName    string `form:"osName,optional"`
	StartTime int64  `form:"startTime,optional"`
	EndTime   int64  `form:"endTime,optional"`
}