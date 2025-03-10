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

	// Teams
	router.GET("/teams/:id", server.getTeam)
	// router.GET("/teams/", server.listTeams)
	// router.POST("/team/:id", server.createTeam) // Is this even needed? Teams get created once a season

	// Users
	router.GET("/users/:id", server.getUser)
	router.POST("/users", server.createUser)
	// competitions
	// matches
	//

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
