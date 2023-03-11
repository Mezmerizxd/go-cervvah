package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mezmerizxd/zvyezda/api/social_app/v1/account"
	"github.com/mezmerizxd/zvyezda/api/social_app/v1/messaging"
	"github.com/mezmerizxd/zvyezda/api/social_app/v1/profile"
	features "github.com/mezmerizxd/zvyezda/features"
)

type Config struct {
	Features features.Features
}

func New(handler *gin.Engine, cfg *Config) {
	v1 := handler.Group("/api/social-app/v1")
	{
		// Controllers
		account := account.New(&account.Config{
			Features: cfg.Features,
		})	
		messaging := messaging.New()
		profile := profile.New()
		
		// Routes
		v1_account := v1.Group("/account")
		{
			v1_account.GET("/login", account.Login)
			v1_account.POST("/signup", account.Signup)
			v1_account.POST("/verify-email", account.VerifyEmail)
		}
		v1_messaging := v1.Group("/messaging")
		{
			v1_messaging.POST("/delete-message", messaging.DeleteMessage)
			v1_messaging.POST("/edit-message", messaging.EditMessage)
		}
		v1_profile := v1.Group("/profile")
		{
			v1_profile.GET("/friends", profile.Friends)
			v1_profile.GET("/friend-requests", profile.FriendRequests)
			v1_profile.POST("/handle-friend-request", profile.HandleFriendRequest)
			v1_profile.POST("/send-friend-request", profile.SendFriendRequest)
		}
	}
}