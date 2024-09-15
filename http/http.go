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
	r.POST("/mock", s.CreateMockApi)
	r.GET("/mock/:id", s.GetMockApi)
	r.GET("/gen/*path", s.GenerateDataFromMock)
}

func (s *Server) Run() {
	app := gin.Default()
	app.Use(gin.Recovery())
	// TODO check origin
	api := app.Group("/api")

	s.Route(api)
	app.Run(":8080")
}
