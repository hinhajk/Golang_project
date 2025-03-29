package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "create",
	})
}

func Edit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "edit",
	})
}

func DELETE(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "delete",
	})
}

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get",
	})
}
