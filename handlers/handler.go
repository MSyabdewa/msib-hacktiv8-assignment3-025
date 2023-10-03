package handlers

import (
	"net/http"

	"github.com/MSyabdewa/msib-hacktiv8-assignment3-025/services"
	"github.com/gin-gonic/gin"
)

// GetWindWaterData mengembalikan data kecepatan angin dan kualitas air dalam format JSON.
func GetWindWaterData(c *gin.Context) {
	windWaterData := services.GenerateWindWaterData()
	responseData := windWaterData.ToResponse()

	// Menggunakan gin.IndentedJSON untuk menghasilkan JSON yang terstruktur.
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": responseData,
	})
}
