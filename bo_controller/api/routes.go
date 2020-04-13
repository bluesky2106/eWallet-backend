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

		// auth API group
		auth := api.Group("/auth")
		auth.POST("/register", s.Register)
		auth.POST("/login", authMw.LoginHandler)
		auth.Use(authMw.MiddlewareFunc())
		{
			auth.GET("/user-profile", s.UserProfile)
			auth.PUT("/user-profile", s.UpdateUserProfile)
			auth.POST("/user-change-pwd", s.ChangePwd)
		}
	}
}
