package routes

import (
	"github.com/MSyabdewa/msib-hacktiv8-assignment3-025/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes mengatur rute-rute HTTP untuk mengakses data kecepatan angin dan kualitas air.
func SetupRoutes() {
	router := gin.Default()

	// Rute HTTP GET untuk mendapatkan data kecepatan angin dan kualitas air.
	router.GET("/", handlers.GetWindWaterData)

	// Menjalankan server pada port 5555.
	router.Run(":5555")
}
