package input

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"my-first-golang-program/internal/domain/ports"
)

type HTTPHandler struct {
	getCdrHandler ports.GetCdrPort
}

func NewHTTPHandler(getCdrHandler ports.GetCdrPort) *HTTPHandler {
	return &HTTPHandler{
		getCdrHandler: getCdrHandler,
	}
}

func (httpHandler *HTTPHandler) Get(c *gin.Context) {
	cdrId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
		return
	}

	cdr, err := httpHandler.getCdrHandler.GetCdr(cdrId)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, cdr)
}
