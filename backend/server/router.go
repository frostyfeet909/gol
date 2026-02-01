package server

import (
	"api/config"
	"api/internal/users"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Handlers struct {
	Config      config.Config
	UserHandler *users.Handler
}

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

func createRouters(h Handlers) *echo.Echo {
	router := echo.New()
	router.Validator = &customValidator{validator: validator.New()}

	router.HTTPErrorHandler = echo.DefaultHTTPErrorHandler(
		true,
	)

	router.Use(middleware.BodyLimit(1048576))
	router.Use(middleware.ContextTimeout(60 * time.Second))
	// router.Use(middleware.CSRF())
	router.Use(middleware.Decompress())
	router.Use(middleware.Gzip())
	router.Use(middleware.RequestLogger())
	// router.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20.0)))
	router.Use(middleware.Recover())
	router.Use(middleware.Secure())

	api := router.Group("/api")

	api.GET("/health", health)
	api.GET("/version", version(h.Config))

	registerSwagger(router)

	v1 := api.Group("/v1")

	users.RegisterRouter(v1, h.UserHandler)

	return router
}
