package main

import (
	"fmt"
	"github.com/go-to/egp_backend/model"
	"github.com/go-to/egp_backend/repository"
	"github.com/go-to/egp_backend/router"
	"github.com/go-to/egp_backend/usecase"
	"github.com/go-to/egp_backend/util"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {
	// 環境変数読み込み
	loadEnv()

	// タイムゾーン
	locationName := os.Getenv("LOCATION_NAME")
	util.Init(locationName)

	// DB接続
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbProtocol := os.Getenv("DB_PROTOCOL")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", dbUser, dbPass, dbProtocol, dbHost, dbPort, dbName)
	dbConn, err := model.Init(dsn)
	if err != nil {
		log.Fatal(err)
	}

	// model設定
	configModel := model.NewConfigModel(dbConn)
	shopModel := model.NewShopModel(dbConn)

	// repository設定
	configRepository := repository.NewConfigRepository(*configModel)
	shopRepository := repository.NewShopRepository(*shopModel)

	// usecase設定
	shopUsecase := usecase.NewShopUseCase(*configRepository, *shopRepository)

	apiPortStr := os.Getenv("API_PORT")
	apiPort, err := strconv.Atoi(apiPortStr)
	if err != nil {
		log.Fatal(err)
	}

	u := router.Usecase{
		Shop: *shopUsecase,
	}
	router.New(apiPort, u)
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
