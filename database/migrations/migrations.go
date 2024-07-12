package migrations

import (
	"fmt"

	"github.com/Iris-GH5/BE-Iris/database"
	"github.com/Iris-GH5/BE-Iris/model/entity"

	"log"
)

func RunMigrations() {
	if database.DB == nil {
		fmt.Printf("Database connection: %v\n", database.DB)
		log.Fatal("Database connection is nil")
	}

	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}

	fmt.Println("Migration run successfully")
}
