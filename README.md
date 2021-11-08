# docker-compose-sample-login-keycloak 
http server: Go  
Authentication infrastructure: keycloak

# How to run
1. `docker-compose up`を行い、keycloakを起動する
2. keycloakにログインする
3. クライアント -> testapp -> クレデンシャル -> シークレットをコピー
4. `server/main.go`を書き換える
```
		oauth2Config = &oauth2.Config{
			ClientID:     "testapp",
			ClientSecret: "ここを書き換える",
			Endpoint:     provider.Endpoint(),
			Scopes:       []string{oidc.ScopeOpenID},
			RedirectURL:  "http://localhost:8000/callback",
		}
```
5. serverディレクトリに移動し、 `go run main.go`をする
6. `http://localhost:8000`にアクセスする