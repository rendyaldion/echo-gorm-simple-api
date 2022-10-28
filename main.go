package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rendyaldion/echo-api/api"
	"github.com/rendyaldion/echo-api/database"
	"github.com/rendyaldion/echo-api/repositories"
)

func main() {
	db := database.InitDatabase()
	e := echo.New()
	repo := repositories.NewRepository(db.DB)
	api.NewServer(e, repo).Run()
}