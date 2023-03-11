package social_app

import (
	"errors"

	"github.com/mezmerizxd/zvyezda/data"
	"github.com/mezmerizxd/zvyezda/schemas"
)

type Config struct {
	Data data.Data
}

type SocialApp interface {
	GetAccountByID(id string) (*schemas.AccountSchema, error)
	GetProfileByID(id string) (*schemas.ProfileSchema, error)
	RegisterAccount(account schemas.AccountForm, profile schemas.ProfileForm) (*schemas.AccountSchema, *schemas.ProfileSchema, error)
	LoginAccount(email string, password string) (*schemas.AccountSchema, *schemas.ProfileSchema, error)
}

type socialApp struct {
	Data data.Data
}

func New(cfg *Config) SocialApp {
	return &socialApp{
		Data: cfg.Data,
	}
}

func (s *socialApp) GetAccountByID(id string) (*schemas.AccountSchema, error) {
	accounts, err := s.Data.SocialAppData.GetLocalAccounts()
	if err != nil {
		return nil, err
	}

	for _, account := range accounts {
		if account.AccountID == id {
			return &account, nil
		}
	}

	return nil, errors.New("[social_app] account not found")
}

func (s *socialApp) GetProfileByID(id string) (*schemas.ProfileSchema, error) {
	profiles, err := s.Data.SocialAppData.GetLocalProfiles()
	if err != nil {
		return nil, err
	}

	for _, profile := range profiles {
		if profile.AccountID == id {
			return &profile, nil
		}
	}

	return nil, errors.New("[social_app] profile not found")
}

func (s *socialApp) RegisterAccount(account schemas.AccountForm, profile schemas.ProfileForm) (*schemas.AccountSchema, *schemas.ProfileSchema, error) {
	/*
		Things that could be added later:
		- Email checks, check if its banned or not
		- Check for illegal characters
	*/

	// Create account
	acc, err := s.Data.SocialAppData.CreateAccountLocally(account)
	if err != nil {
		return nil, nil, err
	}

	// Create profile
	prof, err := s.Data.SocialAppData.CreateProfileLocally(acc.AccountID, profile)
	if err != nil {
		return nil, nil, err
	}

	if acc != nil && prof != nil {
		return acc, prof, nil
	}

	return nil, nil, errors.New("[social_app] failed to register account")
}

func (s *socialApp) LoginAccount(email string, password string) (*schemas.AccountSchema, *schemas.ProfileSchema, error) {
	// Get account
	acc, err := s.Data.SocialAppData.GetAccountByEmail(email)
	if err != nil {
		return nil, nil, err
	}

	// Check password
	if acc.Password != password {
		return nil, nil, errors.New("[social_app] password is incorrect")
	}

	// Get profile
	prof, err := s.GetProfileByID(acc.AccountID)
	if err != nil {
		return nil, nil, err
	}

	if acc != nil && prof != nil {
		return acc, prof, nil
	}

	return nil, nil, errors.New("[social_app] failed to login account")
}