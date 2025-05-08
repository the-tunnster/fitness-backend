package login

import (
	"log"
	"context"
	
	"fitness-tracker/internal/config"

	"golang.org/x/oauth2"
	"github.com/coreos/go-oidc/v3/oidc"
)

type OIDCConfig struct {
	Provider     *oidc.Provider
	Verifier     *oidc.IDTokenVerifier
	OAuth2Config oauth2.Config
}

var OIDC OIDCConfig

func InitOIDC() {
	cfg := config.AppConfig

	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	if err != nil {
		log.Fatalf("Failed to create OIDC provider: %v", err)
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: cfg.OIDCClientID})

	oauth2Config := oauth2.Config{
		ClientID:     cfg.OIDCClientID,
		ClientSecret: cfg.OIDCClientSecret,
		RedirectURL:  cfg.OIDCRedirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	OIDC = OIDCConfig{
		OAuth2Config : oauth2Config,
		Verifier : verifier,
		Provider : provider,
	}
}
