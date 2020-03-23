package api

import (
	jwt "github.com/appleboy/gin-jwt/v2"
)

// Routes : ...
func (s *Server) Routes(authMw *jwt.GinJWTMiddleware) {
	s.g.GET("/", s.DefaultWelcome)
	api := s.g.Group("/api")
	{
		api.GET("/", s.Welcome)

		// deviceToken API group
		productGroup := api.Group("/product-group")
		// productGroup.Use(authMw.MiddlewareFunc())
		{
			productGroup.POST("/add", s.AddProductGroup)
		}
	}
}
