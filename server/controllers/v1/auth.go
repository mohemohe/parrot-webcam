package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/mohemohe/parrot-webcam/server/models"
	"net/http"
)

type (
	PostAuthBody struct {
		ID string `json:"id"` // TODO:
		Password string `json:"password"`
	}
	PostAuthResult struct {
		User *models.User `json:"user"`
		Token *string `json:"token"`
	}
)

func AuthStatus(c echo.Context) error {
	user := c.Get("User").(*models.User)
	token := models.GenerateJwtClaims(user)
	return c.JSON(http.StatusOK, PostAuthResult{
		User: user,
		Token: token,
	})
}

func Auth(c echo.Context) error {
	body := new(PostAuthBody)
	if err := c.Bind(body); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	user, token := models.AuthroizeUser(body.ID, body.Password)
	if token == nil {
		return c.NoContent(http.StatusUnauthorized)
	}

	return c.JSON(http.StatusOK, PostAuthResult{
		User: user,
		Token: token,
	})
}