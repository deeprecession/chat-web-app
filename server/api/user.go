package api

import (
	"github.com/gin-gonic/gin"
)

// get user info
// (GET /api/user/me)
func (s Server) GetApiUserMe(ctx *gin.Context) {
	token, err := getTokenFromContext(ctx)
	if err != nil {
		_ = GetApiUserMe401JSONResponse{
			UnauthorizedJSONResponse{
				Message: err.Error(),
			},
		}.VisitGetApiUserMeResponse(ctx.Writer)

		return
	}

	username, err := s.userAuth.GetUsernameFromToken(token)
	if err != nil {
		_ = GetApiUserMe401JSONResponse{
			UnauthorizedJSONResponse{
				Message: err.Error(),
			},
		}.VisitGetApiUserMeResponse(ctx.Writer)

		return
	}

	_ = GetApiUserMe200JSONResponse{
		Username: username,
	}.VisitGetApiUserMeResponse(ctx.Writer)
}
