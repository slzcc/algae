package devops

import (
	"github.com/gin-gonic/gin"
	"log"
)

func GetHeaderByName(ctx *gin.Context, key string) string {
	return ctx.Request.Header.Get(key)
}

func VersionControlEntrance(c *gin.Context) {

	json := make(map[string]interface{})
	c.BindJSON(&json)
	log.Printf("json %v",&json)

	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	log.Printf("buf %v",string(buf[0:n]))

	event := GetHeaderByName(c, "X-Gitlab-Event")
	log.Printf("event %v", event)

	token := GetHeaderByName(c, "X-Gitlab-Token")
	log.Printf("event %v", token)

	c.JSON(200, json)
}

func VersionControlExit(c *gin.Context) {

	var data string
	data = "test"
	log.Printf("%v", &data)

	c.JSON(200, data)
}
