package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/techschool/simplebank/sqlc"
)

type Server struct{
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server{
     server :=  &Server{store:  store}
	 router := gin.Default()

	 router.POST("/account",  server.createAccount)
	 router.GET("/account/:id",  server.getAccount)
	 router.GET("/account",  server.listAccount)

	 server.router = router
	 return server
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
	}

func errorResponse(err error) gin.H {
	return  gin.H{"error": err.Error()}
}