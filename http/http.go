package http

import (
	"github.com/gin-gonic/gin"
	"sm/service"
)

type Server struct {
	IService service.IService
}

func NewServer(service service.IService) *Server {
	return &Server{
		IService: service,
	}
}

func (s *Server) Route(r *gin.RouterGroup) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
}

func (s *Server) Run() {
	app := gin.Default()
	app.Use(gin.Recovery())
	// TODO check origin
	api := app.Group("/api")

	s.Route(api)
	app.Run(":8080")
}
