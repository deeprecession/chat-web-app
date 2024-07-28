package api

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/deeprecession/chat-web-app/api/auth"
)

// signup a new chatter
// (POST /api/auth/signup)
func (s Server) PostApiAuthSignup(ctx *gin.Context) {
	var user CreateUserRequest

	err := ctx.BindJSON(&user)
	if err != nil {
		_ = PostApiAuthSignup400JSONResponse{
			BadRequestJSONResponse: BadRequestJSONResponse{
				Message: fmt.Sprintf("failed to parse json: %s", err),
			},
		}.VisitPostApiAuthSignupResponse(ctx.Writer)

		return
	}

	token, err := s.userAuth.Signup(auth.User{Username: user.Username, Password: user.Password})
	if errors.Is(err, auth.UsernameIsTakenError) {
		_ = PostApiAuthSignup409JSONResponse{
			Message: err.Error(),
		}.VisitPostApiAuthSignupResponse(ctx.Writer)

		return
	}
	if err != nil {
		_ = PostApiAuthSignup500JSONResponse{
			ServerErrorJSONResponse{
				Message: err.Error(),
			},
		}.VisitPostApiAuthSignupResponse(ctx.Writer)

		return
	}

	bearerToken := fmt.Sprintf("Bearer %s", token)

	_ = PostApiAuthSignup200Response{
		Headers: PostApiAuthSignup200ResponseHeaders{
			Authorization: bearerToken,
		},
	}.VisitPostApiAuthSignupResponse(ctx.Writer)
}

// login a new chatter
// (POST /api/auth/login)
func (s Server) PostApiAuthLogin(ctx *gin.Context) {
	userCreds := LoginUserRequest{}

	err := ctx.BindJSON(&userCreds)
	if err != nil {
		_ = PostApiAuthLogin400JSONResponse{
			BadRequestJSONResponse: BadRequestJSONResponse{
				Message: fmt.Sprintf("failed to parse json: %s", err),
			},
		}.VisitPostApiAuthLoginResponse(ctx.Writer)

		return
	}

	token, err := s.userAuth.Login(
		auth.UserCreds{Username: userCreds.Username, Password: userCreds.Password},
	)
	if errors.Is(err, auth.IncorrectCredentialsError) {
		_ = PostApiAuthLogin401JSONResponse{
			UnauthorizedJSONResponse{
				Message: err.Error(),
			},
		}.VisitPostApiAuthLoginResponse(ctx.Writer)

		return
	}
	if err != nil {
		_ = PostApiAuthLogin500JSONResponse{
			ServerErrorJSONResponse{
				Message: err.Error(),
			},
		}.VisitPostApiAuthLoginResponse(ctx.Writer)

		return
	}

	bearerToken := fmt.Sprintf("Bearer %s", token)

	_ = PostApiAuthLogin200Response{
		Headers: PostApiAuthLogin200ResponseHeaders{Authorization: bearerToken},
	}.VisitPostApiAuthLoginResponse(ctx.Writer)
}
