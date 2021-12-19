package server

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"strings"

	"shop/common"

	"github.com/tal-tech/go-zero/core/logx"
)

func login(record *common.UserInfo) error {
	md5Coding := md5.New()
	md5Coding.Write([]byte(record.Passwd))
	userPasswd := hex.EncodeToString(md5Coding.Sum(nil))

	file, err := os.Open(common.UserInfoPATH)
	if err != nil {
		logx.Errorf("open file failed: %s", err)
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		lineInfo, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		lineInfoSlice := strings.Split(string(lineInfo), ":")
		if lineInfoSlice == nil || len(lineInfoSlice) < 2 {
			continue
		}

		if record.Name == lineInfoSlice[0] && userPasswd == lineInfoSlice[1] {
			UserReocrd.Store(record.Name, record.LoginTime)
			return nil
		}
	}

	return errors.New("can't match the passwd")
}
