package cmd

import (
	"Hands/conf"
	"fmt"
)

func Start() {
	conf.InitConfig()
}

func Clean() {
	fmt.Println("--------------clean--------------")
}
