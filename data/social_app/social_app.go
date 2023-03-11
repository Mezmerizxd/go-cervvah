package social_app

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"
	"strconv"

	"github.com/mezmerizxd/zvyezda/schemas"
	"golang.org/x/crypto/bcrypt"
)

type Config struct {}

type SocialApp interface {
	GetLocalDatabase() (schemas.LocalDatabaseSchema, error)
	GetLocalAccounts() ([]schemas.AccountSchema, error)
	CreateAccountLocally(account schemas.AccountForm) (*schemas.AccountSchema, error)
	ChangeEmailLocally(accountID string, newEmail string) error
	ChangePasswordLocally(accountID string, newPassword string) error
	GetLocalProfiles() ([]schemas.ProfileSchema, error)
	CreateProfileLocally(accountID string, profile schemas.ProfileForm) (*schemas.ProfileSchema, error)
	GetAccountByEmail(email string) (*schemas.AccountSchema, error)
}

type socialApp struct {}

func New(cfg *Config) SocialApp {
	return &socialApp{}
}

var _localAccounts []schemas.AccountSchema
var _localProfiles []schemas.ProfileSchema

func (s *socialApp) GetLocalDatabase() (schemas.LocalDatabaseSchema, error) {
	return schemas.LocalDatabaseSchema{
		Accounts: _localAccounts,
		Profiles: _localProfiles,
	}, errors.New("[social_app] no data inside local database")
}

func (s *socialApp) GetLocalAccounts() ([]schemas.AccountSchema, error) {
	if len(_localAccounts) > 0 {
		return _localAccounts, nil
	}
	return nil, errors.New("[social_app] no local accounts found")
}

func (s *socialApp) CreateAccountLocally(account schemas.AccountForm) (*schemas.AccountSchema, error) {
	// Check account doesnt exist
	for _, acc := range _localAccounts {
		if acc.Email == account.Email {
			return nil, errors.New("[social_app] account already exists")
		}
	}

	accountID := strconv.Itoa(len(_localAccounts) + 1)

	newAccount := schemas.AccountSchema{
		AccountID: accountID,
		Email:     account.Email,
		Password:  account.Password,
	}

	_localAccounts = append(_localAccounts, newAccount)
	return &newAccount, nil
}

func (s *socialApp) ChangeEmailLocally(accountID string, newEmail string) error {
	// Check email isnt in use
	for _, acc := range _localAccounts {
		if acc.Email == newEmail {
			return errors.New("[social_app] email already in use")
		}
	}
	// Change email
	for i, acc := range _localAccounts {
		if acc.AccountID == accountID {
			_localAccounts[i].Email = newEmail
		}
	}
	return nil
}

func (s *socialApp) ChangePasswordLocally(accountID string, newPassword string) error {
	// Change password
	for i, acc := range _localAccounts {
		if acc.AccountID == accountID {
			_localAccounts[i].Password = newPassword
		}
	}
	return nil
}

func (s *socialApp) GetLocalProfiles() ([]schemas.ProfileSchema, error) {
	if len(_localProfiles) > 0 {
		return _localProfiles, nil
	}
	return nil, errors.New("[social_app] no local profiles found")
}

func (s *socialApp) CreateProfileLocally(accountID string, profile schemas.ProfileForm) (*schemas.ProfileSchema, error) {
	// Check profile doesnt exist
	for _, prof := range _localProfiles {
		if prof.AccountID == accountID {
			return nil, errors.New("[social_app] profile already exists")
		}
	}

	// Get account
	var account schemas.AccountSchema
	for _, acc := range _localAccounts {
		if acc.AccountID == accountID {
			account = acc
		}
	}

	// Generate jwt token
	hash, err := bcrypt.GenerateFromPassword([]byte(account.Email), bcrypt.DefaultCost)
  if err != nil {
      log.Fatal(err)
  }
	hasher := md5.New()
  hasher.Write(hash)
  jwt := hex.EncodeToString(hasher.Sum(nil))


	newProfile := schemas.ProfileSchema{
		AccountID: 		accountID,
		Token: 				jwt,
		FirstName: 		profile.FirstName,
		LastName: 		profile.LastName,
		PhoneNumber: 	profile.PhoneNumber,
		ProfilePicture: profile.ProfilePicture,
	}

	_localProfiles = append(_localProfiles, newProfile)
	return &newProfile, nil
}

func (s *socialApp) GetAccountByEmail(email string) (*schemas.AccountSchema, error) {
	for _, acc := range _localAccounts {
		if acc.Email == email {
			return &acc, nil
		}
	}
	return nil, errors.New("[social_app] account not found")
}