package tools

import (
	"github.com/astaxie/beego/logs"
	"log"
)

func Panic(message string) {
	logs.Error(message)
	log.Panic(message)
}

func Fatal(message string) {
	logs.Error(message)
	log.Fatal(message)
}
