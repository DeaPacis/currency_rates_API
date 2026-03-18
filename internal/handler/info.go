package handler

import (
	"app/internal/config"
	"app/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg config.Config
}

func NewHandler(cfg config.Config) *Handler {
	return &Handler{cfg: cfg}
}

func (h *Handler) Info(c *gin.Context) {
	response := model.InfoResponse{
		Version: h.cfg.Version,
		Service: h.cfg.Service,
		Author:  h.cfg.Author,
	}

	c.JSON(http.StatusOK, response)
}
