package internal

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/haquenafeem/shrinkie/api"
	"github.com/haquenafeem/shrinkie/db"
	"github.com/haquenafeem/shrinkie/internal/consts"
	"github.com/haquenafeem/shrinkie/repository"
	"github.com/joho/godotenv"
)

type AppRunner struct{}

func NewAppRunner() *AppRunner {
	return &AppRunner{}
}

func (appRunner *AppRunner) RunMust() {
	err := appRunner.Run()
	if err != nil {
		panic(err)
	}
}

func (appRunner *AppRunner) loadEnv() error {
	return godotenv.Load()
}

func (appRunner *AppRunner) getDbPath() string {
	dbPath := os.Getenv("DB_PATH")
	if dbPath != "" {
		return dbPath
	}

	return consts.DEFAULT_DB_PATH
}

func (appRunner *AppRunner) getPort() (int, error) {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		return consts.DEFAULT_PORT, nil
	}

	return strconv.Atoi(portStr)
}

func (appRunner *AppRunner) Run() error {
	err := appRunner.loadEnv()
	if err != nil {
		return err
	}

	db, err := db.DB(appRunner.getDbPath())
	if err != nil {
		return err
	}

	repo, err := repository.New(db)
	if err != nil {
		return err
	}

	engine := gin.New()
	api := api.New(repo, engine)

	port, err := appRunner.getPort()
	if err != nil {
		return err
	}

	log.Println("server starting at ", port)
	api.Run(port)

	return nil
}
