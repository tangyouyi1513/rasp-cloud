package models

import (
	"github.com/astaxie/beego"
	"rasp-cloud/tools"
	"regexp"
)

var (
	user     string
	password string
)

func init() {
	user = beego.AppConfig.String("user")
	password = beego.AppConfig.String("passwd")
	hasNum := regexp.MustCompile(".*[0-9].*").Match([]byte(password))
	hasLetter := regexp.MustCompile(".*([a-z]|[A-Z]).*").Match([]byte(password))
	if len(user) == 0 || len(password) < 8 || len(password) > 50 || !hasNum || !hasLetter {
		tools.Panic("the login user and password can not be empty," +
			"the length of password can not be less than 8," +
			"the length of password can not be greater than 50," +
			"password must contain numbers and letters")
	}
}

func GetLoginUser() string {
	return user
}

func GetLoginPasswd() string {
	return password
}
