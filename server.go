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
		fmt.Print("SOMETHING WENT WRONG WITH DOTFILES")
		return
	}

	mysqlConnStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_IP"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(mysqlConnStr), &gorm.Config{})

	if err != nil {
		fmt.Print("SOMETHING WENT WRONG WITH DATABASE")
		return
	}

	gin.SetMode(gin.ReleaseMode)
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
