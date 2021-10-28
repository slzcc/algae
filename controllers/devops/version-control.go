package devops

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func GetHeaderByName(ctx *gin.Context, key string) string {
	return ctx.Request.Header.Get(key)
}

func VersionControlEntrance(c *gin.Context) {

	_json := make(map[string]interface{})
	c.BindJSON(&_json)
	log.Printf("json %v",&_json)

	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	log.Printf("buf %v",string(buf[0:n]))

	event := GetHeaderByName(c, "X-Gitlab-Event")
	log.Printf("event %v", event)

	token := GetHeaderByName(c, "X-Gitlab-Token")
	log.Printf("token %v", token)

	client := &http.Client{}
	data, err := json.Marshal(_json)
	if err != nil {
		// handle error
	}
	reader := bytes.NewReader(data)
	req, err := http.NewRequest("POST", "http://jenkins.aws.ops.zhangyue-ops.com/project/public/compile", reader)
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gitlab-Event", event)
	req.Header.Set("X-Gitlab-Token", token)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

	c.JSON(200, _json)
}

func VersionControlExit(c *gin.Context) {

	var data string
	data = "test"
	log.Printf("%v", &data)

	c.JSON(200, data)
}
