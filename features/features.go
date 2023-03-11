package features

import "github.com/mezmerizxd/zvyezda/features/social_app"

type Config struct {
	SocialApp social_app.SocialApp
}

type Features struct {
	SocialAppFeatures social_app.SocialApp
}

func New(cfg *Config) Features {
	return Features{
		SocialAppFeatures: cfg.SocialApp,
	}
}