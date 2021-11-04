package devops

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetHeaderByName(ctx *gin.Context, key string) string {
	return ctx.Request.Header.Get(key)
}

func RequestWebHooks(c *gin.Context) {
	pause, _ := c.GetQuery("pause")

	_json := make(map[string]interface{})
	c.ShouldBind(&_json)
	//c.BindJSON(&_json)
	log.Printf("json %v",&_json)

	data, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))

	// 获取 body
	//buf := make([]byte, 1024)
	//n, _ := c.Request.Body.Read(buf)
	//log.Printf("buf %v",string(buf[0:n]))

	event := GetHeaderByName(c, "X-Gitlab-Event")
	log.Printf("event %v", event)

	token := GetHeaderByName(c, "X-Gitlab-Token")
	log.Printf("token %v", token)

	client := &http.Client{}
	data, err := json.Marshal(&_json)
	if err != nil {
		panic(err.Error())
	}
	reader := bytes.NewReader(data)
	req, err := http.NewRequest("POST", "http://jenkins.aws.ops.zhangyue-ops.com/project/public/compile/webhook_build", reader)
	if err != nil {
		panic(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gitlab-Event", event)
	req.Header.Set("X-Gitlab-Token", token)

	if pause != "" {
		_pause, err := strconv.Atoi(pause)
		if err != nil {
			panic(err.Error())
		}
		time.Sleep(time.Duration(_pause) * time.Second)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()
}

func VersionControlEntrance(c *gin.Context) {
	go RequestWebHooks(c)

	var data string
	data = "successful"

	c.JSON(200, data)
}

func VersionControlExit(c *gin.Context) {

	var data string
	data = "test"
	log.Printf("%v", &data)

	c.JSON(200, data)
}
