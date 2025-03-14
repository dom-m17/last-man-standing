package api

import (
	db "lms/db/sqlc"

	"github.com/gin-gonic/gin"
)

// serve http requests
type Server struct {
	store  db.Store
	router *gin.Engine
}

// creates a new http server
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// teams
	router.GET("/teams/:id", server.getTeam)

	// competitions

	// users
	router.GET("/users/:id", server.getUser)
	router.POST("/users", server.createUser)

	// matches

	// entries

	// competition_matches

	// selections

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
