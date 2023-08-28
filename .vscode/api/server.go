package api

import (
	db "GOLANG/db/sqlc"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

//Server is serve all requests

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer Creates a New Http Server and setupn Routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//routes
	router.POST("/accounts", server.CreateAccount)
	server.router = router
	return server

}

// start the server on specific address
func (server *Server) Start(address string) error {

	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
