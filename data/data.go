package data

import "github.com/mezmerizxd/zvyezda/data/social_app"

type Config struct {
	SocialApp social_app.SocialApp
}

type Data struct {
	SocialAppData social_app.SocialApp
}

func New(cfg *Config) Data {
	err := social_app.SeedLocalAccountsDatabase(cfg.SocialApp)
	if err != nil {
		panic(err)
	}

	return Data{
		SocialAppData: cfg.SocialApp,
	}
}