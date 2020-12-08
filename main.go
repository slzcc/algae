package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/routers/istio"
	"github.com/slzcc/algae/routers/k8s"
	"github.com/slzcc/algae/routers/metrics"
)

func main() {
	/* 路由初始化*/
	engine := gin.Default()

	//跨域
	engine.Use(Cors())

	istio.LRouter(engine)
	k8s.LRouter(engine)
	metrics.LRouter(engine)

	engine.Run(":5801")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		//请求方法
		method := c.Request.Method

		//请求头部
		origin := c.Request.Header.Get("Origin")

		// 声明请求头keys
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			// 这是允许访问所有域
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")

			// 服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")

			// header 的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")

			// 允许跨域设置可以返回其他子段
			// 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")

			// 缓存请求信息 单位为秒
			c.Header("Access-Control-Max-Age", "172800")

			// 跨域请求是否需要带cookie信息 默认设置为true
			c.Header("Access-Control-Allow-Credentials", "false")

			// 设置返回格式是json
			c.Set("content-type", "application/json")
		}

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()
	}
}
