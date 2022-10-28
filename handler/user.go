package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rendyaldion/echo-api/models"
)

func (h *Handler) CreateUser(c echo.Context) error {

	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error": err.Error(),
		})
	}

	err := h.repo.CreateUser(context.Background(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error": err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK, map[string]bool{
		"success": true,
	})
}

func (h *Handler) GetUsers(c echo.Context) error {
	users, err := h.repo.GetUsers(context.Background(), []models.User{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data": users,
	})
}

func (h *Handler) GetUser(c echo.Context) error {
	user ,err := h.repo.GetUser(context.Background(), models.User{}, c.Param("name"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"user": user,
	})
}

func (h *Handler) UpdateUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error": err.Error(),
		})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	err := h.repo.UpdateUser(context.Background(), user, int64(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]bool{
		"success": true,
	})
}

func (h *Handler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.repo.DeleteUser(context.Background(), models.User{}, int64(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]bool{
		"success": true,
	})
}