package router

import (
	"Hands/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitUserRoutes() {
	RegistRoute(func(public, auth *gin.RouterGroup) {
		userApi := api.NewUserApi()
		publicUser := public.Group("/user")
		{
			publicUser.POST("/login", userApi.Login)
		}
		authUser := auth.Group("user")
		{
			authUser.GET("", func(c *gin.Context) {
				// {data: [{id:1,name: "test"}]}
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"data": []map[string]interface{}{
						{"id": 1, "name": "test"},
						{"id": 2, "name": "test2"},
					},
				})
			})

			authUser.GET("/:id", func(c *gin.Context) {
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"id":   1,
					"name": "test",
				})
			})
		}
	})
}
