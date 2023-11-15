package api

import (
	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

func NewUserApi() *UserApi {
	return &UserApi{}
}

// Login @Summary User Login
// @Tags User
// @Description User Login Detailed Description
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Success 200 {string} json "{"msg": "Login Success"}"
// @Failure 400 {string} json "{"msg": "Login Failed"}"
// @Router /api/v1/public/user/login [post]
func (u *UserApi) Login(c *gin.Context) {
	//fmt.Println("Loing 执行了")
	//c.AbortWithStatusJSON(http.StatusOK, gin.H{
	//	"msg": "Login Success",
	//})

	OK(c, ResponseJson{
		Msg: "Login Success",
	})
}
