package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type ResponseJson struct {
	Status int    `json:"-"`
	Code   int    `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
}

func (r ResponseJson) IsEmpty() bool {
	return reflect.DeepEqual(r, ResponseJson{})
}

func HttpRespones(c *gin.Context, status int, resp ResponseJson) {
	if resp.IsEmpty() {
		c.AbortWithStatus(status)
		return
	}
	c.AbortWithStatusJSON(status, resp)
}

func buildStatus(resp ResponseJson, DefaultStatus int) int {
	if 0 == resp.Status {
		return DefaultStatus
	}

	return resp.Status
}

func OK(c *gin.Context, resp ResponseJson) {
	HttpRespones(c, buildStatus(resp, http.StatusOK), resp)
}

func Fail(c *gin.Context, resp ResponseJson) {
	HttpRespones(c, buildStatus(resp, http.StatusBadRequest), resp)

}

func ServerFail(c *gin.Context, resp ResponseJson) {
	HttpRespones(c, buildStatus(resp, http.StatusInternalServerError), resp)
}
