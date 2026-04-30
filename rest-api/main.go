package main

import (
	"log"

	// Frameworks & Drivers
	// "github.com"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// Local Packages (Replace 'your-project' with your module name)
	"example.com/domain"
	"example.com/repository"
	"example.com/service"
	"example.com/web"
)

func main() {
	// 1. Setup Database Connection
	dsn := "host=localhost user=fleet_db_user password=pass dbname=fleet port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 2. Run Migrations (Creates the 'drivers' table automatically)
	db.AutoMigrate(&domain.Driver{})

	// 3. Dependency Injection
	// Repository -> Service -> Handler
	repo := repository.NewDriverRepository(db)
	svc := service.NewDriverService(repo)
	handler := web.NewDriverHandler(svc)

	// 4. Setup Gin Router
	r := gin.Default()

	// 5. Routes
	driverRoutes := r.Group("/drivers")
	{
		// You can add handler.Update and handler.Delete similarly
		driverRoutes.POST("/", handler.Create)      // Create
		driverRoutes.GET("/", handler.GetAll)      // Get All (New)
		driverRoutes.GET("/:id", handler.Get)      // Get One
		driverRoutes.PATCH("/:id", handler.Update)  // Update (New)
		driverRoutes.DELETE("/:id", handler.Delete)
	}

	// 6. Start Server
	log.Println("Server starting on :8082...")
	r.Run(":8082")
}

