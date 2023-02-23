package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func initServer(address string, router *gin.Engine) server {
	// 运行一个服务
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
