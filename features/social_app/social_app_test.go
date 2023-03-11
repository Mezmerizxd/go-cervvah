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

var newAccount1 = schemas.AccountForm{
	Email:    "test@test.com",
	Password: "abc123",
}
var newProfile1 = schemas.ProfileForm{
	FirstName: "Albert",
	LastName:  "Einstein",
	PhoneNumber: "123456789",
	ProfilePicture: "https://www.google.com",
}

var newAccount2 = schemas.AccountForm{
	Email:    "golang@test.com",
	Password: "golang123",
}
var newProfile2 = schemas.ProfileForm{
	FirstName: "Golang",
	LastName:  "Gopher",
	PhoneNumber: "123456789",
	ProfilePicture: "https://www.google.com",
}

var account1 *schemas.AccountSchema
var profile1 *schemas.ProfileSchema

func TestRegisterAccount(t *testing.T) {
	var err error
	account1, profile1, err = f.SocialAppFeatures.RegisterAccount(newAccount1, newProfile1)
	_, _, err2 := f.SocialAppFeatures.RegisterAccount(newAccount2, newProfile2)
	if err != nil {
		t.Errorf("RegisterAccount() failed: %v", err)
	}
	if err2 != nil {
		t.Errorf("RegisterAccount() failed: %v", err2)
	}

	accounts, err := d.SocialAppData.GetLocalAccounts()
	if err != nil {
		t.Errorf("GetLocalAccounts() failed: %v", err)
	}

	t.Log("Accounts found", accounts)
}

func TestGetAccountByID(t *testing.T) {
	acc, err2 := f.SocialAppFeatures.GetAccountByID(account1.AccountID)
	if err2 != nil {
		t.Errorf("GetAccountByID() failed: %v", err2)
	}

	t.Log("Account found", acc)
} 

func TestGetProfileByID(t *testing.T) {
	prof, err2 := f.SocialAppFeatures.GetProfileByID(profile1.AccountID)
	if err2 != nil {
		t.Errorf("GetProfileByID() failed: %v", err2)
	}

	t.Log("Profile found", prof)
}

func TestLoginAccount(t *testing.T) {
	acc, prof, err := f.SocialAppFeatures.LoginAccount(newAccount1.Email, newAccount1.Password)
	if err != nil {
		t.Errorf("LoginAccount() failed: %v", err)
	}

	t.Log("Account found", acc)
	t.Log("Profile found", prof)
}