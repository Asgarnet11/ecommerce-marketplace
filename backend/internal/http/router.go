package http

import (
	"marketplace/backend/internal/handlers"
	"marketplace/backend/internal/middleware"
	"marketplace/backend/internal/store"

	"github.com/gin-gonic/gin"
)

func NewRouter(db *store.DB) *gin.Engine {
  r := gin.Default()
  r.Use(middleware.CORS())

  r.GET("/health", handlers.Health)

  h := handlers.NewProductHandler(db)
  v1 := r.Group("/v1")
  {
    v1.GET("/products", h.ListProducts)
    v1.GET("/products/:slug", h.GetProduct)
  
	ch := handlers.NewCategoryHandler(db)
	v1.GET("/categories", ch.List)        
	v1.GET("/categories/tree", ch.Tree)  

	// shops
	sh := handlers.NewShopHandler(db)
	v1.GET("/shops", sh.List)
	v1.GET("/shops/:slug", sh.Get)

	}	
  return r
}
