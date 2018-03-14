package main

import (
	"log"

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
	self.router.GET("/get/:uuid/:key", func(c *gin.Context) {
		log.Println(c.Param("uuid"), c.Param("key"))
		c.String(200, "%v", 666)
	})

	self.router.Run(conf.Path)
}
