package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginReq struct {
	UserName string
	Pwd      string
}

func Login(c *gin.Context) {
	req := loginReq{}
	c.Bind(&req)
	//c.BindJSON()
	//c.BindQuery()
	c.JSON(http.StatusOK, req)
}
