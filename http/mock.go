package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sm/model"
)

func (s *Server) CreateMockApi(c *gin.Context) {
	var req model.MockApi
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := s.IService.InsertMockApi(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (s *Server) GetMockApi(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	result, err := s.IService.GetMockApi(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (s *Server) GenerateDataFromMock(c *gin.Context) {
	fmt.Println(c.Param("path"))
}
