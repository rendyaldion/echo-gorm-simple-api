package handler

import "github.com/rendyaldion/echo-api/repositories"

type Handler struct {
	repo repositories.Repositories
}

func InitHandler(repo repositories.Repositories) *Handler {
	return &Handler{
		repo: repo,
	}
}