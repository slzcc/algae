package v1

import (
	"github.com/slzcc/algae/controllers/devops"
	"github.com/gin-gonic/gin"
)

// https://books.studygolang.com/advanced-go-programming-book/ch6-web/ch6-02-router.html
func VersionControl(router *gin.RouterGroup){
	router.POST("", devops.VersionControlEntrance)
	router.GET("", devops.VersionControlExit)
}