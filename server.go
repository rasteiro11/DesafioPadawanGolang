package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"server/controller"
	"server/entity"
	"server/repository"
	"server/service"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&entity.ExchangeResponse{})
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("FUCK SOMETHING WENT WRONG WITH DOTFILES")
		return
	}

	mysqlConnStr := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	//dsn := "admin:dora2012@tcp(127.0.0.1:3306)/golang_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(mysqlConnStr), &gorm.Config{})

	if err != nil {
		fmt.Print("FUCK SOMETHING WENT WRONG WITH DATABASE")
		return
	}

	r := gin.Default()

	// REPOSITORIES
	exchangeResponseRepo := repository.NewExchangeResponseRepositoryImpl(db)

	// SERVICES
	exchangeService := service.NewExchangeServiceImpl(exchangeResponseRepo)

	//CONTROLLERS
	exchangeController := controller.NewExchangeController(r, exchangeService)
	exchangeController.MountRoutes()

	Migrate(db)

	r.Run()

}
