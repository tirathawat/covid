package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tirathawat/covid/config"
	"github.com/tirathawat/covid/internal/covid"
)

// Register register all routes.
func Register(cfg *config.App) *gin.Engine {
	e := gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "COVID API"})
	})

	e.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	hCovid := covid.NewHandler(covid.NewClient(cfg.CovidURL))
	e.GET("/covid/summary", hCovid.Summary)

	return e
}
