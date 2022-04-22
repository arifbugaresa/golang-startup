package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-startup/handler"
	"golang-startup/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type users struct {
	ID             int
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type campaigns struct {
	ID               int
	UserId           int
	Name             string
	ShortDescription string
	Description      string
	GoalAmount       int
	CurrentAmount    int
	Perks            string
	BackerCount      int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type campaignImages struct {
	ID          int
	Campaign_id int
	FileName    string
	IsPrimary   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func main() {

	//Connect to database
	dsn := "host=localhost user=postgres password=paramadaksa dbname=startup port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("DB Connected")
	}

	db.AutoMigrate(&users{}, &campaigns{}, &campaignImages{})

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()
}
