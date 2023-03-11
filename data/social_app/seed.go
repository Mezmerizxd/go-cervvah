package social_app

import (
	"errors"
	"log"
	"strconv"

	"github.com/mezmerizxd/zvyezda/schemas"
)

var _seedLocalAccounts []schemas.AccountForm = []schemas.AccountForm{
	{
		Email:    "test@test.com",
		Password: "abc123",
	},
}

var _seedLocalProfiles []schemas.ProfileForm = []schemas.ProfileForm{
	{
		FirstName: "Test",
		LastName:  "Test",
		PhoneNumber: "1234567890",
		ProfilePicture: "https://i.imgur.com/3ZQ3Z0M.png",
	},
}

func SeedLocalAccountsDatabase(s SocialApp) error {
	for i, acc := range _seedLocalAccounts {
		s.CreateAccountLocally(acc)
		s.CreateProfileLocally(strconv.Itoa(i), _seedLocalProfiles[i])
	}

	// Check accounts were created
	database, err := s.GetLocalDatabase()
	if err != nil {
		return errors.New("[social_app] failed to seed local accounts database")
	}

	for i, acc := range database.Accounts {
		if acc.Email != _seedLocalAccounts[i].Email {
			return errors.New("[social_app] failed to seed local accounts database")
		}
	}
	for i, prof := range database.Profiles {
		if prof.FirstName != _seedLocalProfiles[i].FirstName {
			return errors.New("[social_app] failed to seed local accounts database")
		}
	}

	log.Println("[social_app] successfully seeded local accounts database")
	return nil
}