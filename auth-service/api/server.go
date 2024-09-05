package api

import (
	"net/http"

	database "github.com/DEVunderdog/auth-service/database/sqlc"
	"github.com/DEVunderdog/auth-service/utils"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config utils.Config
	store database.Store
	router *gin.Engine
}

func NewServer(config utils.Config, store database.Store) (*Server, error) {

	server := &Server{
		config: config,
		store: store,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	server.router = router
}

func (server *Server) Start(address string) *http.Server {
	srv := &http.Server{
		Addr: address,
		Handler: server.router,
	}

	return srv
}

