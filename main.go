package main

import (
	"controller"
	"models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	// Maintenance History Routes
	r.GET("/maintenance", controller.GetMaintenanceHistories)
	r.GET("/maintenance/:id", controller.GetMaintenanceHistory)
	r.POST("/maintenance", controller.CreateMaintenanceHistory)
	r.PUT("/maintenance/:id", controller.UpdateMaintenanceHistory)
	r.DELETE("/maintenance/:id", controller.DeleteMaintenanceHistory)

	// Taxes Routes
	r.GET("/taxes", controller.GetTaxes)
	r.GET("/taxes/:id", controller.GetTax)
	r.POST("/taxes", controller.CreateTax)
	r.PUT("/taxes/:id", controller.UpdateTax)
	r.DELETE("/taxes/:id", controller.DeleteTax)

	// Users Routes
	r.GET("/users", controller.GetUsers)
	r.GET("/users/:id", controller.GetUser)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users/:id", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)

	// Cars Routes
	r.GET("/cars", controller.GetCars)
	r.GET("/cars/:id", controller.GetCar)
	r.POST("/cars", controller.CreateCar)
	r.PUT("/cars/:id", controller.UpdateCar)
	r.DELETE("/cars/:id", controller.DeleteCar)

	// Rentals Routes
	r.GET("/rentals", controller.GetRentals)
	r.GET("/rentals/:id", controller.GetRental)
	r.POST("/rentals", controller.CreateRental)
	r.PUT("/rentals/:id", controller.UpdateRental)
	r.DELETE("/rentals/:id", controller.DeleteRental)

	// Payments Routes
	r.GET("/payments", controller.GetPayments)
	r.GET("/payments/:id", controller.GetPayment)
	r.POST("/payments", controller.CreatePayment)
	r.PUT("/payments/:id", controller.UpdatePayment)
	r.DELETE("/payments/:id", controller.DeletePayment)

	r.Run(":8080")
}
