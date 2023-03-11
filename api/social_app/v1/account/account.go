package account

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mezmerizxd/zvyezda/features"
)

type LoginForm struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Config struct {
	Features features.Features
}

type Account interface{
	Login(c *gin.Context)
	Signup(c *gin.Context)
	VerifyEmail(c *gin.Context)
}

type account struct{
	Features features.Features
}

func New(cfg *Config) Account {
	return &account{
		Features: cfg.Features,
	}
}

func (a *account) Login(c *gin.Context) {
	// Get the login form
	var r LoginForm
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Println("Login form: ", r)

	// // Verify the login details
	// account, err := a.Features.SocialApp.Account.VerifyLoginDetails(r.Email, r.Password)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }
	account, profile, err := a.Features.SocialAppFeatures.LoginAccount(r.Email, r.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Println("Account: ", account)
	log.Println("Profile: ", profile)

	c.JSON(200, gin.H{"message": "Login successful", "token": profile.Token})
}

func (a *account) Signup(c *gin.Context) {
	// f := feature.New()

	// f.Data.SocialAppData.Account.CreateAccountLocally(account_data.AccountSchema{
	// 		Email: "test@test.com",
	// 		Password: "abc123",
	// })
}

func (a *account) VerifyEmail(c *gin.Context) {}