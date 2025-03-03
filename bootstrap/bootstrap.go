package bootstrap

import (
	"m-server-api/config"
	"m-server-api/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type Server struct {
	g *gin.Engine
}

// NewServer creates a new server
func NewServer() *Server {
	var s Server
	gin.SetMode(config.Get().Server.Mode)
	s.g = InitRouter()
	return &s
}

// Run starts the server
func (s *Server) Run() {
	s.g.Run(":" + cast.ToString(config.Get().Server.Port))
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.ErrorHandler)

	apiBase := r.Group(
		"/"+config.Get().Server.Prefix,
		middlewares.AuthMiddleware(),
		middlewares.ErrorHandler,
	)

	if len(GroupList) > 0 {
		for _, group := range GroupList {
			grp := apiBase.Group(group.RelativePath, group.Handlers...)
			for _, r := range group.Router {
				if r.Method == "ANY" {
					grp.Any(r.RelativePath, r.HandlerFunc...)
				} else {
					grp.Handle(r.Method, r.RelativePath, r.HandlerFunc...)
				}
			}
		}
	}

	return r
}
