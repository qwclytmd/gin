package main

import (
	"bcw/app/common/public"
	"bcw/config"
	"bcw/routes"
	"bcw/server"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	gin.SetMode(config.GetViper().GetString("server.run_mode"))

	StartServers()
	r := routes.SetupRouter()

	srv := &http.Server{
		Addr:    ":" + config.GetViper().GetString("server.port"),
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func StartServers() {
	//服务初始化
	server.Initialize()
	//同步权限
	public.Permissions{}.SyncCasbinRules(server.HttpServers.Casbin)
}
