package main

import (
	"github.com/MSyabdewa/msib-hacktiv8-assignment3-025/routes"
	"github.com/MSyabdewa/msib-hacktiv8-assignment3-025/services"
)

func main() {
	// Menjalankan goroutine untuk memperbarui data secara terus-menerus.
	go services.StartUpdatingData()

	// Mengatur rute-rute HTTP dan menjalankan server.
	routes.SetupRoutes()
}
