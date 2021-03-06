package router

import (
	"allaboutapps.dev/aw/go-starter/internal/api"
	"allaboutapps.dev/aw/go-starter/internal/api/handlers"
	"allaboutapps.dev/aw/go-starter/internal/api/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func Init(s *api.Server) {
	s.Echo = echo.New()

	s.Echo.Debug = s.Config.Echo.Debug
	s.Echo.HideBanner = true
	// TODO write proper wrapper for echo logger and zerolog/use library for proper logging support including levels and paylods
	s.Echo.Logger.SetOutput(log.With().Str("component", "echo").Str("level", "debug").Logger())

	s.Echo.HTTPErrorHandler = HTTPErrorHandler

	// ---
	// General middleware
	s.Echo.Pre(echoMiddleware.RemoveTrailingSlash())

	s.Echo.Use(echoMiddleware.Recover())
	s.Echo.Use(echoMiddleware.RequestID())
	s.Echo.Use(middleware.Logger())

	// TODO: fully remove this, no auth-required middleware should be applied by default!
	// We will fully let the groups / handlers decided which middlewares to use
	s.Echo.Use(middleware.AuthWithConfig(middleware.AuthConfig{S: s, Mode: middleware.AuthModeRequired, Skipper: func(c echo.Context) bool {
		switch c.Path() {
		case "/api/v1/auth/forgot-password",
			"/api/v1/auth/forgot-password/complete",
			"/api/v1/auth/login",
			"/api/v1/auth/refresh",
			"/api/v1/auth/register",
			"/testing-f679dbac-62bb-445d-b7e8-9f2c71ca382c":
			return true
		}

		return false
	}}))

	// ---
	// Initialize our general groups and set middleware to use above them
	s.Router = &api.Router{
		Root:      s.Echo.Group("/"),
		APIV1Auth: s.Echo.Group("/api/v1/auth"),
		APIV1Users: s.Echo.Group("/api/v1/users",
			middleware.AuthWithConfig(middleware.AuthConfig{S: s, Mode: middleware.AuthModeSecure, Scopes: []string{"cms"}})),
	}

	// ---
	// Finally attach our handlers
	handlers.AttachAllRoutes(s)
}
