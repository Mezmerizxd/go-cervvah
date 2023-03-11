package messaging

import "github.com/gin-gonic/gin"

type Messaging interface {
	DeleteMessage(c *gin.Context)
	EditMessage(c *gin.Context)
}

type messaging struct{}

func New() Messaging {
	return &messaging{}
}

func (m *messaging) DeleteMessage(c *gin.Context) {}

func (m *messaging) EditMessage(c *gin.Context) {}