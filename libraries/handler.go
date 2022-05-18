package libraries

import (
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
	Engine *echo.Echo
}

func NewHandler(logger Logger) Handler {
	newHandler := echo.New()
	newHandler.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: logger.NewEchoLogger()}))

	return Handler{Engine: newHandler}
}
