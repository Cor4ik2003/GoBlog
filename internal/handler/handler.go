package handler

import (
	"github.com/Corik2003/GoBlog/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func (h *Handler)  Init(api *gin.RouterGroup) {
	v1 := api.Group("/api")
	{
		h.InitUserHandler(v1)
	}
}