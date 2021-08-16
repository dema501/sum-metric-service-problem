package rest

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

// Server is a rest with store
type Server struct {
	LimitReqSec   int
	BasicAuthUser string
	BasicAuthPWD  string
	ServerPort    string
}

// Run the lister and request's router, activate rest server
func (s *Server) Run(options ...func(*gin.RouterGroup)) {
	log.Printf("[INFO] activate rest server")

	router := gin.New()

	router.Use(gin.Recovery())

	router.Use(s.loggerMiddleware())

	router.GET("/ping", s.pingCtrl)

	v1 := router.Group("/v1")

	// Cors headers
	configCors := cors.DefaultConfig()
	configCors.AllowAllOrigins = true
	configCors.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	v1.Use(cors.New(configCors))

	// Set Authorization if we have ENV settings
	if len(s.BasicAuthUser) > 0 {
		v1.Use(gin.BasicAuth(gin.Accounts{
			s.BasicAuthUser: s.BasicAuthPWD,
		}))
	}

	for _, op := range options {
		if op != nil {
			op(v1)
		}
	}

	log.Fatal(router.Run(":" + s.ServerPort))
}

// PingHandler handle url /v1/ping
func (s *Server) pingCtrl(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (s *Server) loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()

		log.Printf("[INFO] %s %s %s %v %d",
			c.Request.Method, c.Request.URL.Path,
			c.ClientIP(), time.Since(t), c.Writer.Status())
	}
}
