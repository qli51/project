package server

import (
	"testing"
	"encoding/json"
	"errors"
	"io/ioutil"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"fmt"
	"database/sql"

	"shop/common"

	. "github.com/agiledragon/gomonkey"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func TestParseUserInfo(t *testing.T) {
	Convey("解析成功", t, func() {
		patcheOne := ApplyFunc(ioutil.ReadAll, func(io.Reader) ([]byte, error) {
			return []byte("read succeed"), nil
		})
		defer patcheOne.Reset()

		patcheTwo := ApplyFunc(json.Unmarshal, func([]byte, interface{}) error {
			return nil
		})
		defer patcheTwo.Reset()

		err := parseUserInfo(&http.Request{}, &common.UserInfo{})
		So(err, ShouldBeNil)
	})
	Convey("读取返回信息失败", t, func() {
		patcheOne := ApplyFunc(ioutil.ReadAll, func(io.Reader) ([]byte, error) {
			return []byte("read failed"), errors.New("read failed")
		})
		defer patcheOne.Reset()

		err := parseUserInfo(&http.Request{}, &common.UserInfo{})
		So(err, ShouldNotBeNil)
	})
	Convey("解析失败", t, func() {
		patcheOne := ApplyFunc(ioutil.ReadAll, func(io.Reader) ([]byte, error) {
			return []byte("read succeed"), nil
		})
		defer patcheOne.Reset()

		patcheTwo := ApplyFunc(json.Unmarshal, func([]byte, interface{}) error {
			return errors.New("parse failed")
		})
		defer patcheTwo.Reset()

		err := parseUserInfo(&http.Request{}, &common.UserInfo{})
		So(err, ShouldNotBeNil)
	})
}

func TestCheckLogin(t *testing.T) {
	Convey("未登录，返回错", t, func() {
		res := checkLogin("test", "test")
		So(res, ShouldEqual, false)
	})
	Convey("模拟登录场景", t, func() {
		UserReocrd.Store("test", "test")
		res := checkLogin("test", "test")
		So(res, ShouldEqual, true)
	})
}

func TestServerLogin(t *testing.T) {
	w := httptest.NewRecorder()
	Convey("成功登入", t, func() {
		patcheOne := ApplyFunc(parseUserInfo, func(*http.Request, *common.UserInfo) error {
			return nil
		})
		defer patcheOne.Reset()

		file, _ := os.Open("../dataRecord/userInfo")
		fmt.Println(common.UserInfoPATH)
		patcheTwo := ApplyFunc(os.Open, func(string) (*os.File, error) {
			return file, nil
		})
		defer patcheTwo.Reset()

		serverLogin(w, &http.Request{})
	})
	Convey("解析失败", t, func() {
		patcheOne := ApplyFunc(parseUserInfo, func(*http.Request, *common.UserInfo) error {
			return errors.New("parse user information failed")
		})
		defer patcheOne.Reset()
		serverLogin(w, &http.Request{})
	})
	Convey("读文件失败", t, func() {
		patcheOne := ApplyFunc(parseUserInfo, func(*http.Request, *common.UserInfo) error {
			return nil
		})
		defer patcheOne.Reset()

		patcheTwo := ApplyFunc(os.Open, func(string) (*os.File, error) {
			return &os.File{}, errors.New("read file failed")
		})
		defer patcheTwo.Reset()
		serverLogin(w, &http.Request{})
	})
}

func TestServerLogout(t *testing.T) {
	w := httptest.NewRecorder()
	Convey("成功登出", t, func() {
		patcheOne := ApplyFunc(parseUserInfo, func(*http.Request, *common.UserInfo) error {
			return nil
		})
		defer patcheOne.Reset()

		patcheTwo := ApplyFunc(checkLogin, func(string, string) bool {
			return false
		})
		defer patcheTwo.Reset()
		serverLogout(w, &http.Request{})

	})
	Convey("解析用户信息失败", t, func() {
		patcheOne := ApplyFunc(parseUserInfo, func(*http.Request, *common.UserInfo) error {
			return errors.New("parse user information failed")
		})
		defer patcheOne.Reset()
		serverLogout(w, &http.Request{})
	})
	Convey("已不是最新用户信息", t, func() {
		patcheOne := ApplyFunc(parseUserInfo, func(*http.Request, *common.UserInfo) error {
			return nil
		})
		defer patcheOne.Reset()
		patcheTwo := ApplyFunc(checkLogin, func(string, string) bool {
			return false
		})
		defer patcheTwo.Reset()
		serverLogout(w, &http.Request{})
	})
}

func TestServerDataOrder(t *testing.T) {
	w := httptest.NewRecorder()
	Convey("成功下单", t, func() {
		patcheOne := ApplyFunc(parseUserInfo, func(*http.Request, *common.UserInfo) error {
			return nil
		})
		defer patcheOne.Reset()

		patcheTwo := ApplyFunc(checkLogin, func(string, string) bool {
			return false
		})
		defer patcheTwo.Reset()

		patcheThree := ApplyFunc(httpx.Parse, func(*http.Request, interface{}) error {
			return nil
		})
		defer patcheThree.Reset()

		patcheFour := ApplyFunc(sql.Open, func(string, string) (*sql.DB,error) {
			return &sql.DB{}, nil
		})
		defer patcheFour.Reset()
		
		serverDataOrder(w, &http.Request{})
	})
}

func TestServerDataRecharge(t *testing.T) {
	w := httptest.NewRecorder()
	Convey("成功充值", t, func() {
		patcheOne := ApplyFunc(parseUserInfo, func(*http.Request, *common.UserInfo) error {
			return nil
		})
		defer patcheOne.Reset()

		patcheTwo := ApplyFunc(checkLogin, func(string, string) bool {
			return false
		})
		defer patcheTwo.Reset()

		patcheThree := ApplyFunc(httpx.Parse, func(*http.Request, interface{}) error {
			return nil
		})
		defer patcheThree.Reset()

		patcheFour := ApplyFunc(sql.Open, func(string, string) (*sql.DB,error) {
			return &sql.DB{}, nil
		})
		defer patcheFour.Reset()

		serverDataRecharge(w, &http.Request{})
	})
}



