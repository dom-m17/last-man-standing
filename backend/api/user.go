package api

import (
	"database/sql"
	db "lms/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type createUserRequest struct {
	Username      string         `json:"username" binding:"required,alphanum"`
	Password      string         `json:"password" binding:"required,min=6"`
	FirstName     string         `json:"first_name" binding:"required"`
	LastName      string         `json:"last_name" binding:"required"`
	Email         string         `json:"email" binding:"required,email"`
	PhoneNumber   sql.NullString `json:"phone_number"`
	FavouriteTeam sql.NullInt64  `json:"favourite_team"`
}

type createuserResponse struct {
	Username      string         `json:"username"`
	FirstName     string         `json:"first_name"`
	LastName      string         `json:"last_name"`
	Email         string         `json:"email"`
	PhoneNumber   sql.NullString `json:"phone_number"`
	FavouriteTeam sql.NullInt64  `json:"favourite_team"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// TODO: Password to be hashed here

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: req.Password,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		FavouriteTeam:  req.FavouriteTeam,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createuserResponse{
		Username:      user.Username,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Email:         user.Email,
		PhoneNumber:   user.PhoneNumber,
		FavouriteTeam: user.FavouriteTeam,
	}

	ctx.JSON(http.StatusOK, rsp)
}
