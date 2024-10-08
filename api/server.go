package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/techschool/simplebank/sqlc"
)

type Server struct{
	store db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server{
     server :=  &Server{store:  store}
	 router := gin.Default()

	 if v, ok :=  binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
		}

	 router.POST("/users",  server.createUser)

	 router.POST("/account",  server.createAccount)
	 router.GET("/account/:id",  server.getAccount)
	 router.GET("/account",  server.listAccount)
	 
	 router.POST("/transfers",  server.createTransfer)

	 server.router = router
	 return server
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
	}

func errorResponse(err error) gin.H {
	return  gin.H{"error": err.Error()}
}