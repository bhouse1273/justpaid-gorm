package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/bhouse1273/justpaid-gorm/internal/config"
	"github.com/bhouse1273/justpaid-gorm/internal/db"
	"github.com/bhouse1273/justpaid-gorm/internal/handlers"
	"go.uber.org/zap"

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
	// Echo's logger isn't a *zap.Logger; provide a safe fallback zap logger for global config
	config.JPLogger = zap.NewNop()

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

	// Resource CRUD
	api.GET("/organizations", h.ListOrganizations)
	api.POST("/organizations", h.CreateOrganization)
	api.GET("/organizations/:id", h.GetOrganization)
	api.PUT("/organizations/:id", h.UpdateOrganization)
	api.DELETE("/organizations/:id", h.DeleteOrganization)

	api.GET("/properties", h.ListProperties)
	api.POST("/properties", h.CreateProperty)
	api.GET("/properties/:id", h.GetProperty)
	api.PUT("/properties/:id", h.UpdateProperty)
	api.DELETE("/properties/:id", h.DeleteProperty)

	api.GET("/offers", h.ListOffers)
	api.POST("/offers", h.CreateOffer)
	api.GET("/offers/:id", h.GetOffer)
	api.PUT("/offers/:id", h.UpdateOffer)
	api.DELETE("/offers/:id", h.DeleteOffer)

	api.GET("/portfolios", h.ListPortfolios)
	api.POST("/portfolios", h.CreatePortfolio)
	api.GET("/portfolios/:id", h.GetPortfolio)
	api.PUT("/portfolios/:id", h.UpdatePortfolio)
	api.DELETE("/portfolios/:id", h.DeletePortfolio)

	api.GET("/port-templates", h.ListPortTemplates)
	api.POST("/port-templates", h.CreatePortTemplate)
	api.GET("/port-templates/:id", h.GetPortTemplate)
	api.PUT("/port-templates/:id", h.UpdatePortTemplate)
	api.DELETE("/port-templates/:id", h.DeletePortTemplate)

	api.GET("/campaigns", h.ListCampaigns)
	api.POST("/campaigns", h.CreateCampaign)
	api.GET("/campaigns/:id", h.GetCampaign)
	api.PUT("/campaigns/:id", h.UpdateCampaign)
	api.DELETE("/campaigns/:id", h.DeleteCampaign)

	api.GET("/accounts", h.ListAccounts)
	api.POST("/accounts", h.CreateAccount)
	api.GET("/accounts/:id", h.GetAccount)
	api.PUT("/accounts/:id", h.UpdateAccount)
	api.DELETE("/accounts/:id", h.DeleteAccount)

	api.GET("/payments", h.ListPayments)
	api.POST("/payments", h.CreatePayment)
	api.GET("/payments/:id", h.GetPayment)
	api.PUT("/payments/:id", h.UpdatePayment)
	api.DELETE("/payments/:id", h.DeletePayment)

	api.GET("/users", h.ListUsers)
	api.POST("/users", h.CreateUser)
	api.GET("/users/:id", h.GetUser)
	api.PUT("/users/:id", h.UpdateUser)
	api.DELETE("/users/:id", h.DeleteUser)

	// Serve OpenAPI specification
	e.GET("/openapi.yaml", func(c echo.Context) error {
		return c.File("./openapi.yaml")
	})

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
