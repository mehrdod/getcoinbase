package handler

import (
	"getcoinbase/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		coins := api.Group("/coins")
		{
			coins.GET("/:pair", h.getDataByPairs)
		}
	}

	return router
}
