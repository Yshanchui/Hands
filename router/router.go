package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type fnRegistRoute = func(public *gin.RouterGroup, auth *gin.RouterGroup)

var (
	fnRoutes []fnRegistRoute
)

func RegistRoute(fn fnRegistRoute) {
	if fn == nil {
		return
	}
	fnRoutes = append(fnRoutes, fn)
}

func InitRouter() {
	r := gin.Default()

	public := r.Group("/api/v1/public")
	auth := r.Group("/api/v1")

	InitBasePlatformRoutes()

	for _, fnRegistRoute := range fnRoutes {
		fnRegistRoute(public, auth)
	}

	prot := viper.GetString("server.port")
	if prot == "" {
		prot = "8080"
	}
	err := r.Run(":" + prot)
	if err != nil {
		panic("Start Server Error: " + err.Error())
	}
}

func InitBasePlatformRoutes() {
	InitUserRoutes()
}
