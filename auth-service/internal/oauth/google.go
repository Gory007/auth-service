package oauth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GoogleOAuthConfig = &oauth2.Config{
		ClientID:     "629066593447-62bahlmhe6efc61s4ooh3jl4pl6h51qb.apps.googleusercontent.com", // Замени на свой
		ClientSecret: "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",                                      // Замени на свой
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	// Случайная строка для защиты от CSRF
	RandomState = "random"
)
