package data

// WindWaterDataResponse adalah format JSON yang diinginkan untuk data kecepatan angin dan kualitas air.
type WindWaterDataResponse struct {
	WindSpeed    int    `json:"wind_speed"`
	WaterQuality int    `json:"water_quality"`
	WindStatus   string `json:"wind_status"`
	WaterStatus  string `json:"water_status"`
}

// WindWaterData struct menyimpan data terkait kecepatan angin dan kualitas air.
type WindWaterData struct {
	WindSpeed    int
	WaterQuality int
	WindStatus   string
	WaterStatus  string
}

// ToResponse mengonversi data kecepatan angin dan kualitas air ke dalam format JSON.
func (d WindWaterData) ToResponse() WindWaterDataResponse {
	return WindWaterDataResponse{
		WindSpeed:    d.WindSpeed,
		WaterQuality: d.WaterQuality,
		WindStatus:   d.WindStatus,
		WaterStatus:  d.WaterStatus,
	}
}
