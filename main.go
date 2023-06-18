package main

import (
	"FinalProject_Rental-Car-Management/controller"
	"FinalProject_Rental-Car-Management/database"
	"FinalProject_Rental-Car-Management/middleware"
	"FinalProject_Rental-Car-Management/utils/initializer"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.EnvLoader()
	database.ConnectDatabase()
}

func main() {
	r := gin.Default()

	// Cars Routes
	r.GET("/api/cars", controller.CarsIndex)
	r.GET("/api/cars/:id", controller.CarsShow)

	// Users Routes
	r.POST("/api/users/login", controller.HandleLogin)
	r.POST("/api/users/logout", controller.HandleLogout)
	r.POST("/api/users", controller.UserCreate)

	authMiddleware := middleware.RequireAuth

	auth := r.Group("/api")
	auth.Use(authMiddleware)
	{
		// cars
		auth.POST("/cars", controller.CarsCreate)
		auth.PUT("/cars/:id", controller.CarsUpdate)
		auth.DELETE("/cars", controller.CarsDelete)

		// Maintenance History Routes
		auth.GET("/maintenance", controller.MaintenanceHistoryIndex)
		auth.GET("/maintenance/:id", controller.MaintenanceHistoryShow)
		auth.POST("/maintenance", controller.MaintenanceHistoryCreate)
		auth.PUT("/maintenance/:id", controller.MaintenanceHistoryUpdate)
		auth.DELETE("/maintenance", controller.MaintenanceHistoryDelete)
		auth.GET("/maintenance/car/:id", controller.MaintenanceHistoryShowByCarID)

		// Payments Routes
		auth.GET("/payments", controller.PaymentIndex)
		auth.GET("/payments/:id", controller.PaymentShow)
		auth.POST("/payments", controller.PaymentCreate)
		auth.PUT("/payments/:id", controller.PaymentUpdate)
		auth.DELETE("/payments", controller.PaymentDelete)

		auth.GET("/users", controller.UserIndex)
		auth.GET("/users/:id", controller.UserShow)
		auth.PUT("/users/:id", controller.UserUpdate)
		auth.DELETE("/users", controller.UserDelete)

		// Rentals Routes
		auth.GET("/rentals", controller.RentalIndex)
		auth.GET("/rentals/:id", controller.RentalShow)
		auth.POST("/rentals", controller.RentalCreate)
		auth.PUT("/rentals/:id", controller.RentalUpdate)
		auth.DELETE("/rentals", controller.RentalDelete)

		// Taxes Routes
		auth.GET("/taxes", controller.TaxIndex)
		auth.GET("/taxes/:id", controller.TaxShow)
		auth.POST("/taxes", controller.TaxCreate)
		auth.PUT("/taxes/:id", controller.TaxUpdate)
		auth.DELETE("/taxes", controller.TaxDelete)
		auth.GET("/taxes/car/:id", controller.GetTaxesByCarID)
	}

	r.Run(os.Getenv("PORT"))
}
