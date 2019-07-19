package webster

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func APIHealthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GoogleCaptchaHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
