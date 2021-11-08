package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

var once sync.Once

var provider *oidc.Provider
var oauth2Config *oauth2.Config

func getConfig() (*oauth2.Config, *oidc.Provider) {
	once.Do(func() {
		var err error
		provider, err = oidc.NewProvider(context.Background(), "http://localhost:8080/auth/realms/master")
		if err != nil {
			panic(err)
		}
		oauth2Config = &oauth2.Config{
			ClientID:     "testapp",
			ClientSecret: "Client Credential",
			Endpoint:     provider.Endpoint(),
			Scopes:       []string{oidc.ScopeOpenID},
			RedirectURL:  "http://localhost:8000/callback",
		}
	})
	return oauth2Config, provider
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := r.Cookie("Authorization"); err != nil {
			config, _ := getConfig()
			url := config.AuthCodeURL("")
			http.Redirect(w, r, url, http.StatusFound)
			return
		}
		io.WriteString(w, "login success")
	})

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		config, provider := getConfig()
		if err := r.ParseForm(); err != nil {
			http.Error(w, "parse form error", http.StatusInternalServerError)
			return
		}

		accessToken, err := config.Exchange(context.Background(), r.Form.Get("code"))
		if err != nil {
			http.Error(w, "Can't get access token", http.StatusInternalServerError)
			return
		}

		rawIDToken, ok := accessToken.Extra("id_token").(string)
		if !ok {
			http.Error(w, "missing token", http.StatusInternalServerError)
			return
		}
		oidcConfig := &oidc.Config{
			ClientID: "testapp",
		}
		verifier := provider.Verifier(oidcConfig)
		idToken, err := verifier.Verify(context.Background(), rawIDToken)
		if err != nil {
			http.Error(w, "id token verify error", http.StatusInternalServerError)
			return
		}

		idTokenClaims := map[string]interface{}{}
		if err := idToken.Claims(&idTokenClaims); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("%#v", idTokenClaims)
		http.SetCookie(w, &http.Cookie{
			Name:  "Authorization",
			Value: "Bearer " + rawIDToken,
			Path:  "/",
		})
		http.Redirect(w, r, "/", http.StatusFound)
	})
	log.Println("start http server :8080")
	log.Println(http.ListenAndServe(":8000", nil))
}
