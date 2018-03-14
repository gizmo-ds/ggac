package main

import (
	"log"
)

var (
	_config_path = "./config.json"
	_users_path  = "./users.json"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	reload_config(_config_path)
	reload_users(_users_path)

	api.init()
}

func main() {
	user_uuid := get_user_uuid("cunyu")
	if user_uuid == "nil" {
		return
	}
	log.Println(get_user_code(user_uuid, "test1"))
}
