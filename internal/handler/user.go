package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitUserHandler(v1 *gin.RouterGroup) {
	v1.POST("/create-user", (h.CreateUser))
}

func(h *Handler) CreateUser(c *gin.Context) error {

}