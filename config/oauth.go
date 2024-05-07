package config

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

type OAuth struct {
	AuthURL string
}

func newOAuth() *OAuth {
	oauthClientID := os.Getenv("OAUTH_CLIENT_ID")
	baseUrl := "https://accounts.google.com/o/oauth2/v2/auth"
	redirectUrl := os.Getenv("OAUTH_GOOGLE_REDIRECT_URL")
	scopeOpenId := "openid"
	scopeProfile := "https://www.googleapis.com/auth/userinfo.profile"
	scopeEmail := "https://www.googleapis.com/auth/userinfo.email"
	scope := scopeOpenId + " " + scopeProfile + " " + scopeEmail
	combinedUrl := "%s?redirect_uri=%s&client_id=%s&scope=%s&response_type=code&access_type=offline"
	authUrl := fmt.Sprintf(combinedUrl, baseUrl, url.QueryEscape(redirectUrl), oauthClientID, strings.ReplaceAll(url.QueryEscape(scope), "+", "%20"))

	return &OAuth{
		AuthURL: authUrl,
	}
}
