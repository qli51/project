package server

import (
	"shop/common"
)

func logout(record *common.UserInfo) error {
	UserReocrd.Delete(record.Name)
	return nil
}
