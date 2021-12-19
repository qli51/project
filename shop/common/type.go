package common

// UserInfoPATH 密码文件路径
const UserInfoPATH = "./dataRecord/user_info"

type UserInfo struct {
	Name   string
	Passwd string
	LoginTime string
}

type RechargeRequestParams struct {
	ID    string  `form:"id"`
	Value float64 `form:"value"`
}

type CheckRequestParams struct {
	ID   string `form:"id"`
}

type OrderRequestParams struct {
	ID     string `form:"id"`
	ShopID string `form:"shopID"`
}
