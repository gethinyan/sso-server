package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetToken 获取 token
func GetToken(c *gin.Context) {
	token, err := c.Cookie("jsonWebToken")
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"token": ""})
		return
	}
	c.JSONP(http.StatusOK, token)
}
