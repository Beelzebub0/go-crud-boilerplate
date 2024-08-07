package restserver

import (
	"fmt"
	"sync"

	"github.com/Beelzebub0/go-crud-boilerplate/src/business/usecase"
	config "github.com/Beelzebub0/go-crud-boilerplate/src/conf"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var once = &sync.Once{}

type REST interface{}

type rest struct {
	uc   usecase.Usecase
	conf config.ServerConfig
}

func Init(uc usecase.Usecase, conf config.ServerConfig) REST {
	var e *rest
	once.Do(func() {
		e = &rest{
			uc:   uc,
			conf: conf,
		}
		e.Serve()
	})
	return e
}

func (e *rest) Serve() {
	// Set Mode for Gin : debug, test, release
	gin.SetMode(e.conf.Gin.Mode)

	// Build Gin Engine
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// Use CORS if by config
	if e.conf.Gin.CORS.Enabled && e.conf.Gin.Mode == "debug" {
		fmt.Print("Running GIN with CORS Enabled.\n")
		r.Use(cors.New(cors.Config{
			AllowOrigins:     e.conf.Gin.CORS.AllowedOrigins,
			AllowMethods:     e.conf.Gin.CORS.AllowedMethods,
			AllowHeaders:     e.conf.Gin.CORS.AllowHeaders,
			ExposeHeaders:    e.conf.Gin.CORS.ExposedHeaders,
			AllowCredentials: e.conf.Gin.CORS.AllowCredentials,
		}))
	}

	// Ping Endpoint
	r.GET("/api/ping", func(c *gin.Context) { c.String(200, "PONG!") })

	// API V1
	v1 := r.Group("/api/v1")

	//Area
	v1.GET("/area", e.GetArea)
	v1.GET("/area/:id", e.GetAreaByID)
	v1.POST("/area", e.CreateArea)
	v1.PUT("/area/:id", e.UpdateArea)
	v1.DELETE("/area/:id", e.DeleteArea)

	r.Run(":" + e.conf.Port)
}
