package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/deeprecession/chat-web-app/api/db"
)

var (
	IncorrectCredentialsError = errors.New("incorrect user credentials")
	UsernameIsTakenError      = errors.New("username is taken")
)

type User struct {
	Username string
	Password string
}

type UserCreds struct {
	Username string
	Password string
}

type UserAuthJWT struct {
	secretKey []byte
	storage   db.MongoStorage
}

func NewUserAuthJWT(secretKey []byte, storage db.MongoStorage) UserAuthJWT {
	return UserAuthJWT{
		secretKey: secretKey,
		storage:   storage,
	}
}

func (userAuth UserAuthJWT) Login(creds UserCreds) (string, error) {
	user, err := userAuth.storage.GetUser(creds.Username)
	if err != nil {
		if errors.Is(err, db.UserNotExist) {
			return "", IncorrectCredentialsError
		}

		return "", err
	}

	if user.Password != creds.Password {
		return "", IncorrectCredentialsError
	}

	token, err := userAuth.generateJWT(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (userAuth UserAuthJWT) Signup(user User) (string, error) {
	err := userAuth.storage.InsertUser(db.User{Username: user.Username, Password: user.Password})
	if err != nil {
		if errors.Is(err, db.UserAlreadyExist) {
			return "", UsernameIsTakenError
		}

		return "", err
	}

	token, err := userAuth.generateJWT(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (userAuth UserAuthJWT) GetUsernameFromToken(tokenStr string) (string, error) {
	token, err := userAuth.verifyToken(tokenStr)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("failed to get claims")
	}

	username, ok := claims["sub"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get username from claims")
	}

	return username, nil
}

func (userAuth UserAuthJWT) generateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(userAuth.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (userAuth UserAuthJWT) verifyToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return userAuth.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
