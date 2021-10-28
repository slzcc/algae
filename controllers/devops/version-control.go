package devops

import (
	"github.com/gin-gonic/gin"
	"log"
)

func VersionControlEntrance(c *gin.Context) {

	json := make(map[string]interface{})
	c.BindJSON(&json)
	log.Printf("%v",&json)

	c.JSON(200, json)
}

func VersionControlExit(c *gin.Context) {

	var data string
	data = "test"
	log.Printf("%v", &data)

	c.JSON(200, data)
}
