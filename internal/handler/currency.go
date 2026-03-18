package handler

import (
	"app/internal/model"
	"app/internal/service"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Currency(c *gin.Context) {
	currency := strings.ToUpper(c.Query("currency"))
	date := c.Query("date")

	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	rates, err := service.GetCurrencyRates(date)
	if err != nil {
		log.Printf("Error: %s", err)
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: "internal server error",
		})
		return
	}

	data := map[string]float64{}

	if currency != "" {
		value, ok := rates[currency]
		if !ok {
			log.Printf("Error: currency %s not found", currency)
			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Error: "currency not found",
			})
			return
		}

		data[currency] = value
	} else {
		data = rates
	}

	response := model.CurrencyResponse{
		Service: h.cfg.Service,
		Data:    data,
	}

	c.JSON(http.StatusOK, response)
}
