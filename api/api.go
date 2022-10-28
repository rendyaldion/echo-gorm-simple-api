package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rendyaldion/echo-api/handler"
	"github.com/rendyaldion/echo-api/repositories"
)

type Server interface {
	Run()
}

type server struct {
	echo *echo.Echo
	repo repositories.Repositories
}

func NewServer(echo *echo.Echo, repo repositories.Repositories) Server {
	return &server{
		echo: echo,
		repo: repo,
	}
}

func (s *server) router(h *handler.Handler) {
	s.echo.POST("/", h.CreateUser)
	s.echo.GET("/", h.GetUsers)
	s.echo.GET("/:name", h.GetUser)
	s.echo.PUT("/:id", h.UpdateUser)
	s.echo.DELETE("/:id", h.DeleteUser)
}

func (s *server) Run() {
	s.router(handler.InitHandler(s.repo))
	
	if err := s.echo.Start(":9000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}