package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

type UserInput struct {
	Data string `json:"data"`
}

var db *gorm.DB

func main() {
	dsn := "host=localhost user=postgres password=firman dbname=test port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&User{})

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:4200", "http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/api/users", createUser)

	r.Run(":8080")
}

func createUser(c *gin.Context) {
	var input UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	parts := strings.Fields(input.Data)

	if len(parts) < 3 {
		c.JSON(400, gin.H{"error": "Invalid data format"})
		return
	}

	nameParts := []string{}
	age := 0
	for _, part := range parts {
		if _, err := strconv.Atoi(part); err == nil {
			age = parseAge(part)
			break
		} else {
			nameParts = append(nameParts, part)
		}
	}

	name := formatText(strings.Join(nameParts, " "))

	cityParts := []string{}
	for _, part := range parts[len(nameParts)+1:] {
		part = strings.ToUpper(part)
		if part == "TAHUN" || part == "THN" || part == "TH" {
			continue
		}
		cityParts = append(cityParts, part)
	}

	city := formatText(strings.Join(cityParts, " "))

	user := User{
		Name: name,
		Age:  age,
		City: city,
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(201, user)
}
