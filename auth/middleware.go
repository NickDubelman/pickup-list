package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// UserFromContext takes a context and returns the UserInfo for the user making the
// request
func UserFromContext(ctx context.Context) (UserInfo, error) {
	tokenStr, err := AccessTokenFromContext(ctx)
	if err != nil {
		return UserInfo{}, err
	}

	claims := &UserInfo{}
	tkn, err := parseToken(tokenStr, claims)
	if err != nil {
		return UserInfo{}, err
	}

	now := time.Now()
	expires := time.Unix(claims.IssuedAt, 0).Add(accessTokenDuration)

	if !tkn.Valid {
		return UserInfo{}, NotAuthorized{} // Token invalid
	}

	if now.After(expires) {
		return UserInfo{}, TokenExpired{} // Token expired
	}

	return *claims, nil
}

// ContextWithAccessToken takes a context and an access token and returns a new
// context with the access token attached
func ContextWithAccessToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, contextKey{"accessToken"}, token)
}

// AccessTokenFromContext takes a context and returns the attached accessToken
func AccessTokenFromContext(ctx context.Context) (string, error) {
	ctxValue := ctx.Value(contextKey{"accessToken"})
	if ctxValue == nil {
		return "", NotAuthorized{}
	}

	accessToken, ok := ctxValue.(string)
	if !ok {
		return "", fmt.Errorf("Expected accessToken to have type string")
	}

	return accessToken, nil
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authHeader := strings.Split(req.Header.Get("Authorization"), "Token ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		}

		req = req.WithContext(ContextWithAccessToken(req.Context(), authHeader[1]))

		// Do stuff
		next.ServeHTTP(w, req)
	})
}

type contextKey struct{ name string }
