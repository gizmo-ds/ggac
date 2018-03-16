package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type (
	Config struct {
		UserListsPath string
		Path          string
	}
)

var (
	conf Config
	us   []interface{}
	uis  map[string]interface{}
)

func reload_config(path string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(bytes, &conf); err != nil {
		panic(err)
	}
}

func reload_users(path string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(bytes, &us); err != nil {
		panic(err)
	}
}

func get_user_uuid(user, pwd string) (uuid string, err error) {
	uis = make(map[string]interface{})
	uis = us[0].(map[string]interface{})

	t1, err := json.Marshal(uis[user])
	if err != nil {
		return
	}
	var t2 struct {
		UUID string
		PWD  string
	}
	if err = json.Unmarshal(t1, &t2); err != nil {
		return
	}
	if t2.PWD != pwd {
		err = errors.New("password error")
		return
	}
	uuid = t2.UUID
	return
}

func get_user_code(user_uuid string, user_key string) int {
	bytes, err := ioutil.ReadFile(conf.UserListsPath + user_uuid + ".json")
	if err != nil {
		if os.IsExist(err) {
			ioutil.WriteFile(conf.UserListsPath+user_uuid+".json", []byte("[{}]"), 0644)
		} else {
			panic(err)
		}
	}
	var _conf []interface{}
	if err := json.Unmarshal(bytes, &_conf); err != nil {
		log.Println(err)
		return -1
	}
	if keys, ok := _conf[0].(map[string]interface{}); ok {
		key, ok := keys[user_key]
		if !ok {
			log.Printf("not exist \"%v\"", user_key)
			return -1
		}
		return get_google_auth(time.Now(), key.(string))
	}
	return -1
}

func get_key_list(user_uuid string) (key_list []string) {
	bytes, err := ioutil.ReadFile(conf.UserListsPath + user_uuid + ".json")
	if err != nil {
		if os.IsExist(err) {
			ioutil.WriteFile(conf.UserListsPath+user_uuid+".json", []byte("[{}]"), 0644)
		} else {
			panic(err)
		}
	}
	var _conf []interface{}
	if err := json.Unmarshal(bytes, &_conf); err != nil {
		log.Println(err)
		return
	}
	if keys, ok := _conf[0].(map[string]interface{}); ok {
		for k, _ := range keys {
			key_list = append(key_list, k)
		}
	}
	return
}
