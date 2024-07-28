package api

import "github.com/gin-gonic/gin"

// create a chat
// (POST /api/chat)
func (s Server) PostApiChat(ctx *gin.Context) {
	_ = PostApiChat201JSONResponse{}.VisitPostApiChatResponse(ctx.Writer)
}
