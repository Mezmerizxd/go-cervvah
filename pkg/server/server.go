package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	socialapp_v1 "github.com/mezmerizxd/zvyezda/api/social_app/v1"
	v1 "github.com/mezmerizxd/zvyezda/api/v1"
)

type Server struct {
	server *http.Server
}

/*
	The website for this project will run on its own server, so we don't need to serve static files.
	I will make a command that will spin up the website server and the API server at the same time 
	or I will make a command that will spin up the website server and the API server separately.

	Basic Gin Example:
		handler.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

	// Refuse all GET requests to / and /index.html and redirect them to the website.
	handler.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://zvyezda.me")
	})
*/

func New(addr string, cfg *socialapp_v1.Config) *Server {
	handler := gin.Default()

	// Middleware
	handler.Use(gin.Recovery())
	handler.Use(gin.Logger())
	handler.Use(gin.ErrorLogger())

	// Route GET / to the index.html file
	handler.StaticFile("/", "./renderer/index.html")

	// API Controllers
	v1.New(handler)
	socialapp_v1.New(handler, cfg)

	// Web Socket Controllers

	return &Server{
		server: &http.Server{
			Addr: addr,
			Handler: handler,
		},
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

	// router.Use(middleware.Logger)
	// router.Use(middleware.Recoverer)
	// router.Use(middleware.Timeout(time.Second * 15))
	// router.Use(middleware.RequestID)
	// router.Use(cors.Handler(cors.Options{
	// 	AllowedOrigins: []string{"https://*", "http://*", "*"},
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Origin"},
	// 	ExposedHeaders:   []string{"Link", "Access-Control-Allow-Origin", "Accept"},
	// 	AllowCredentials: false,
	// 	MaxAge:           300,
	// }))