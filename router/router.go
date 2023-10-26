package router

import (
	_ "Hands/docs"
	"Hands/global"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type fnRegistRoute = func(public *gin.RouterGroup, auth *gin.RouterGroup)

var (
	fnRoutes []fnRegistRoute
)

// RegistRoute 注册路由
func RegistRoute(fn fnRegistRoute) {
	if fn == nil {
		return
	}
	fnRoutes = append(fnRoutes, fn)
}

// InitRouter 初始化路由
func InitRouter() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// gin注册路由
	r := gin.Default()

	public := r.Group("/api/v1/public")
	auth := r.Group("/api/v1")

	InitBasePlatformRoutes()

	for _, fnRegistRoute := range fnRoutes {
		fnRegistRoute(public, auth)
	}

	// swagger集成
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	prot := viper.GetString("server.port")
	if prot == "" {
		prot = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + prot,
		Handler: r,
	}

	go func() {
		global.Logger.Info("Start Listen:", prot)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
			return
		}
	}()

	<-ctx.Done()
	stop()
	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		//
		global.Logger.Error(fmt.Sprintf("server shutdown: %s\n", err.Error()))
		return
	}
	global.Logger.Info("server exiting")
}

func InitBasePlatformRoutes() {
	InitUserRoutes()
}
