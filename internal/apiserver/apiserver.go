package apiserver

import (
	"github.com/aabbuukkaarr8/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wb-go/wbf/zlog"
)

type APIServer struct {
	config *Config
	router *gin.Engine
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: gin.Default(),
	}

}
func (s *APIServer) Run() error {
	zlog.Logger.Info().Msg("Starting API server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}
func (s *APIServer) configLogger() error { return nil }

func (s *APIServer) ConfigureRouter(handler *handler.Handler) {
	s.router.POST("/shorten", handler.Create)
	s.router.GET("/s/:short_url", handler.Redirect)
	s.router.GET("/analytics/:short_url", handler.GetAnalytics)
}
