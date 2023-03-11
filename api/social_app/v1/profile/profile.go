package profile

import "github.com/gin-gonic/gin"

type Profile interface {
	Friends(c *gin.Context)
	FriendRequests(c *gin.Context)
	HandleFriendRequest(c *gin.Context)
	SendFriendRequest(c *gin.Context)
}

type profile struct{}

func New() Profile {
	return &profile{}
}

func (p *profile) Friends(c *gin.Context) {}

func (p *profile) FriendRequests(c *gin.Context) {}

func (p *profile) HandleFriendRequest(c *gin.Context) {}

func (p *profile) SendFriendRequest(c *gin.Context) {}