package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/bhouse1273/justpaid-gorm/internal/config"
	"github.com/bhouse1273/justpaid-gorm/internal/db"
	"github.com/bhouse1273/justpaid-gorm/internal/handlers"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	_ = godotenv.Load()
	cfg := config.New()
	dbconn := db.MustOpen()

	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		c.Logger().Error(err)
		_ = c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	e.Use(middleware.Recover(), middleware.Logger())

	// CORS
	if len(cfg.CORSOrigins) == 0 {
		e.Use(middleware.CORS())
	} else {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     cfg.CORSOrigins,
			AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodOptions},
			AllowHeaders:     []string{"Content-Type", "Authorization"},
			AllowCredentials: true,
		}))
	}

	h := handlers.New(dbconn)
	e.GET("/healthz", h.Healthz)

	api := e.Group("/api")
	// Optional JWT auth
	if secret := cfg.JWTSecret; strings.TrimSpace(secret) != "" {
		registerAPI(api, secret)
	}

	api.POST("/campaigns/:id/materialize", h.MaterializeCampaign)
	api.POST("/actions/claim", h.ClaimDueActions)
	api.POST("/actions/:id/complete", h.CompleteAction)

	log.Printf("listening on :%s", cfg.Port)
	if err := e.Start(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}

func registerAPI(g *echo.Group, secret string) {
	g.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(secret),
		// Use NewClaimsFunc (not Claims) with jwt/v5:
		NewClaimsFunc: func(c echo.Context) jwt.Claims { return new(jwt.RegisteredClaims) },
		// You can look up tokens in multiple places (comma-separated):
		TokenLookup: "header:Authorization,cookie:access_token,query:access_token",
	}))
}
