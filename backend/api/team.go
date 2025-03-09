package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getTeamRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getTeam(ctx *gin.Context) {
	var req getTeamRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	team, err := server.store.GetTeam(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, team)
}
