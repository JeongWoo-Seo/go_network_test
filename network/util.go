package network

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *Network) okResponse(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, result)
}

func (u *Network) failResponse(c *gin.Context, result interface{}) {
	c.JSON(http.StatusBadRequest, result)
}
