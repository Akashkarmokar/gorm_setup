package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Akashkarmokar/gorm_setup/cmd/models"
	"github.com/Akashkarmokar/gorm_setup/cmd/operatoins"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello from Docker + Air + Go!")

	host := "postgres" // Container Service Name or container name
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	portStr := os.Getenv("DB_PORT")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic("Port is not number")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, dbname, port)

	db, dberr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dberr != nil {
		fmt.Println("Database connection error !!")
	} else {
		fmt.Println("Database connection established !!")
	}

	db.AutoMigrate(&models.User{})

	operatoins.InsertToTheDataBase(db)

}
