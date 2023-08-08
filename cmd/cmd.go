package cmd

import (
	"Hands/conf"
	"Hands/router"
	"fmt"
)

func Start() {
	conf.InitConfig()
	router.InitRouter()
}

func Clean() {
	fmt.Println("--------------clean--------------")
}
