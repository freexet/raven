package graph

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/freexet/raven/auth"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx context.Context, callback func(ctx *gin.Context, user *auth.User) (interface{}, error)) (interface{}, error) {
	gc := ctx.Value(ContextKey{Name: "ginCtx"}).(*gin.Context)

	authHeader := gc.Request.Header.Get("authorization")
	if len(authHeader) == 0 {
		return nil, errors.New("no token provided")
	}

	t := strings.Split(authHeader, " ")

	if t[0] != "Bearer" {
		return nil, errors.New("invalid token")
	}

	token, err := jwt.ParseWithClaims(
		t[1],
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("SECRET_KEY")), nil
		},
	)
	if err != nil {
		return nil, errors.New("could not parse token")
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, errors.New("unauthorized: invalid token")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("token is expired")
	}

	a, _ := gc.Get("auth")
	user, err := a.(auth.Service).GetUserByID(claims.Id)
	if err != nil {
		return nil, err
	}

	return callback(gc, user)
}
