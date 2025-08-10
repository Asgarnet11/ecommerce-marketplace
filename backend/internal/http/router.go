package http

import (
	"marketplace/backend/internal/handlers"
	"marketplace/backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.GET("/health", handlers.Health)

	v1 := r.Group("/v1")
	{
		v1.GET("/products", handlers.ListProducts)      // mock
	v1.GET("/products/:slug", handlers.GetProduct)   // mock
	}

	return r
}
