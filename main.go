package main

import (
	"fmt"
	"log"
	"os"
)

var (
	_config_path = "./config.json"
	_users_path  = "./users.json"
)

func init() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "c": // config
			if len(os.Args) < i+1 {
				fmt.Println("参数错误")
				return
			}
			_config_path = os.Args[i+1]

		case "u": // user-list
			if len(os.Args) < i+1 {
				fmt.Println("参数错误")
				return
			}
			_users_path = os.Args[i+1]

		case "h": // help
			fmt.Println(`xD`)
		}
	}

	reload_config(_config_path)
	reload_users(_users_path)
}

func main() {
	api.init()
}
