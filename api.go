package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	API struct {
		router *gin.Engine
	}
)

var (
	api API
)

func (self *API) init() {
	self.router = gin.New()

	self.router.POST("/list/:u/", func(c *gin.Context) {
		pwd := c.PostForm("pwd")
		uuid, err := get_user_uuid(c.Param("u"), pwd)
		if err != nil {
			c.JSON(200, gin.H{"error": err})
			return
		}
		c.JSON(200, gin.H{"error": nil, "key_list": get_key_list(uuid)})
	})

	self.router.POST("/get/:u/:key", func(c *gin.Context) {
		pwd := c.PostForm("pwd")
		uuid, err := get_user_uuid(c.Param("u"), pwd)
		if err != nil {
			c.JSON(200, gin.H{"error": err})
			return
		}
		code := get_user_code(uuid, c.Param("key"))
		if code <= 0 {
			c.JSON(200, gin.H{"error": "key error"})
			return
		}
		_code := strconv.Itoa(code)
		if len(_code) < 6 {
			_code = "0" + _code
		}
		c.JSON(200, gin.H{"error": nil, "code": _code})
	})

	self.router.Run(conf.Path)
}
