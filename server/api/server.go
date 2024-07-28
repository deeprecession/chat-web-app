package api

import (
	"crypto/ed25519"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/deeprecession/chat-web-app/api/auth"
	"github.com/deeprecession/chat-web-app/api/db"
)

type Server struct {
	userAuth auth.UserAuthJWT
	storage  db.MongoStorage
}

func NewServer(storage db.MongoStorage) (Server, error) {
	_, privateKey, _ := ed25519.GenerateKey(nil)
	userAuth := auth.NewUserAuthJWT(privateKey, storage)

	return Server{
		userAuth: userAuth,
		storage:  storage,
	}, nil
}

func getTokenFromContext(ctx *gin.Context) (string, error) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		return "", fmt.Errorf("no token provided")
	}

	bearerToken := strings.Split(token, " ")
	if len(bearerToken) != 2 {
		return "", fmt.Errorf(`token must be in format "Bearer {token}`)
	}

	return bearerToken[1], nil
}
