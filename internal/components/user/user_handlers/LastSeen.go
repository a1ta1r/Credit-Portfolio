package user_handlers

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/codes"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/loans/services"
	"github.com/gin-gonic/gin"
	"gopkg.in/appleboy/gin-jwt.v2"
	"net/http"
	"time"
)

type LastSeenHandler struct {
	userService services.UserService
}

func NewLastSeenHandler(userService services.UserService) LastSeenHandler {
	return LastSeenHandler{userService: userService}
}

func (lsh LastSeenHandler) UpdateLastSeen(ctx *gin.Context) {
	claims := jwt.ExtractClaims(ctx)
	id := int(claims["user_id"].(float64))
	user := lsh.userService.GetUserByID(uint(id))
	if user.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": codes.Forbidden})
		return
	}

	user.LastSeen = time.Now()
	lsh.userService.UpdateUser(user)
}
