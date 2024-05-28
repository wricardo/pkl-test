package configservice

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"bitbucket.org/zetaactions/pklpoc/config"
	"github.com/gin-gonic/gin"
)

// ConfigService represents the configuration service
type ConfigService struct {
}

// NewConfigService creates a new ConfigService
func NewConfigService() *ConfigService {
	return &ConfigService{}
}

// formationHandler handles requests to the /formation/:formation endpoint
func (s *ConfigService) formationHandler(c *gin.Context) {
	environment := c.Param("environment")
	formation := c.Param("formation")

	// Load the formation file
	filePath := filepath.Join("config", "environment", environment, formation, "main.pkl")
	f, err := config.LoadFromPath(context.Background(), filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error loading formation: %s", err)})
		return
	}

	// Encode the formation data to JSON
	encoded, err := json.Marshal(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error encoding JSON: %s", err)})
		return
	}

	// Set response headers and write the response
	c.Header("Content-Type", "application/json")
	c.Writer.Write(encoded)
}

// BuildRouter creates and configures the Gin router
func (s *ConfigService) BuildRouter() *gin.Engine {
	router := gin.Default()

	// Endpoint to handle formation requests
	router.GET("/formation/:environment/:formation", s.formationHandler)

	// Serve static files from the environment directory
	router.Static("/static", "config")

	return router
}
