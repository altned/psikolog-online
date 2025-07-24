package routes

import (
	"github.com/example/psikolog-online/internal/handlers"
	"github.com/example/psikolog-online/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
			auth.POST("/forgot-password", handlers.ForgotPassword)
		}

		api.Use(middlewares.JWTMiddleware())

		api.POST("/chat/cs", handlers.ChatCS)
		api.POST("/chat/psychologist", handlers.ChatPsychologist)

		api.GET("/tests", handlers.ListTests)
		api.POST("/tests/:id/start", handlers.StartTest)
		api.POST("/tests/:id/submit", handlers.SubmitTest)
		api.GET("/tests/:id/result", handlers.GetTestResult)

		api.GET("/articles", handlers.ListArticles)
		api.GET("/articles/:id", handlers.GetArticle)

		api.GET("/history", handlers.GetHistory)
	}

	return r
}
