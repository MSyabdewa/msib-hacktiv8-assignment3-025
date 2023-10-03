package services

import (
	"math/rand"
	"time"

	"github.com/MSyabdewa/msib-hacktiv8-assignment3-025/data"
)

// generateRandomValue mengembalikan nilai acak antara 0 dan 15.
func GenerateRandomValue() int {
	return rand.Intn(16)
}

// generateWindWaterData menghasilkan data terkini tentang kecepatan angin, kualitas air, dan status keamanannya.
func GenerateWindWaterData() data.WindWaterData {
	waterQuality := GenerateRandomValue()
	windSpeed := GenerateRandomValue()

	waterStatus := "Aman"
	if waterQuality < 5 {
		waterStatus = "Aman"
	} else if waterQuality >= 5 && waterQuality <= 8 {
		waterStatus = "Siaga"
	} else {
		waterStatus = "Bahaya"
	}

	windStatus := "Aman"
	if windSpeed < 6 {
		windStatus = "Aman"
	} else if windSpeed >= 6 && windSpeed <= 15 {
		windStatus = "Siaga"
	} else {
		windStatus = "Bahaya"
	}

	return data.WindWaterData{
		WindSpeed:    windSpeed,
		WaterQuality: waterQuality,
		WindStatus:   windStatus,
		WaterStatus:  waterStatus,
	}
}

// StartUpdatingData secara terus-menerus memperbarui data setiap 15 detik.
func StartUpdatingData() {
	for {
		GenerateWindWaterData()
		time.Sleep(15 * time.Second)
	}
}
