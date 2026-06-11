package handler

import (
	"net/http"

	"github.com/gaboliveiradev/api-jobs/internal/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "failed to retrieve openings")
		return
	}

	sendSuccess(ctx, "list openings", openings)
}
