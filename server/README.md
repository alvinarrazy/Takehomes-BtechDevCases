# How to start

run locally on PORT 8080

```shell
go run ./cmd/rest/main.go
```

# Features

Provides simple auth login, logout, and get user email

## Login

- route `/auth/login`
- body needed are email and password
- Token is set on cookies['auth_token'] by set_cookie response header
- Harcoded users

```go
var users = map[string]string{
	"admin@email.com": "password123",
	"user@email.com":  "userpass",
}
```

## Logout

- route `/auth/logout`
- Logout will remove the cookies using set_cookie

## Get User Email

- route `/user`
- Will only return email
