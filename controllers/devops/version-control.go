package devops

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetHeaderByName(ctx *gin.Context, key string) string {
	return ctx.Request.Header.Get(key)
}

func VersionControlEntrance(c *gin.Context) {

	pause, _ := c.GetQuery("pause")
	if pause != "" {
		_pause, err := strconv.Atoi(pause)
		if err != nil {
			panic(err.Error())
		}
		time.Sleep(time.Duration(_pause) * time.Second)
	}

	_json := make(map[string]interface{})
	c.BindJSON(&_json)
	//log.Printf("json %v",&_json)

	// 获取 body
	//buf := make([]byte, 1024)
	//n, _ := c.Request.Body.Read(buf)
	//log.Printf("buf %v",string(buf[0:n]))

	event := GetHeaderByName(c, "X-Gitlab-Event")
	//log.Printf("event %v", event)

	token := GetHeaderByName(c, "X-Gitlab-Token")
	//log.Printf("token %v", token)

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

	resp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err.Error())
	//}

	//log.Println(string(body))

	c.JSON(200, _json)
}

func VersionControlExit(c *gin.Context) {

	var data string
	data = "test"
	log.Printf("%v", &data)

	c.JSON(200, data)
}
