package social_app_test

import (
	"testing"

	"github.com/mezmerizxd/zvyezda/data"
	socialAppData "github.com/mezmerizxd/zvyezda/data/social_app"
	"github.com/mezmerizxd/zvyezda/features"
	socialAppFeatures "github.com/mezmerizxd/zvyezda/features/social_app"
	"github.com/mezmerizxd/zvyezda/schemas"
)

var sad = socialAppData.New(&socialAppData.Config{})
var d = data.New(&data.Config{
		SocialApp: sad,
})

	// Features
var saf = socialAppFeatures.New(&socialAppFeatures.Config{
		Data: d,
})
var f = features.New(&features.Config{
		SocialApp: saf,
})

var newAccount = schemas.AccountForm{
	Email:    "golang@test.com",
	Password: "golang123",
}
var newProfile = schemas.ProfileForm{
	FirstName: "Golang",
	LastName:  "Gopher",
	PhoneNumber: "123456789",
	ProfilePicture: "https://www.google.com",
}

var account *schemas.AccountSchema
var profile *schemas.ProfileSchema

func TestRegisterAccount(t *testing.T) {
	var err error
	account, profile, err = f.SocialAppFeatures.RegisterAccount(newAccount, newProfile)
	if err != nil {
		t.Errorf("RegisterAccount() failed: %v", err)
	}

	accounts, err := d.SocialAppData.GetLocalAccounts()
	if err != nil {
		t.Errorf("GetLocalAccounts() failed: %v", err)
	}

	t.Log("Accounts found", accounts)
}

func TestGetAccountByID(t *testing.T) {
	acc, err2 := f.SocialAppFeatures.GetAccountByID(account.AccountID)
	if err2 != nil {
		t.Errorf("GetAccountByID() failed: %v", err2)
	}

	t.Log("Account found", acc)
} 

func TestGetProfileByID(t *testing.T) {
	prof, err2 := f.SocialAppFeatures.GetProfileByID(profile.AccountID)
	if err2 != nil {
		t.Errorf("GetProfileByID() failed: %v", err2)
	}

	t.Log("Profile found", prof)
}

func TestLoginAccount(t *testing.T) {
	acc, prof, err := f.SocialAppFeatures.LoginAccount(newAccount.Email, newAccount.Password)
	if err != nil {
		t.Errorf("LoginAccount() failed: %v", err)
	}

	t.Log("Account found", acc)
	t.Log("Profile found", prof)
}