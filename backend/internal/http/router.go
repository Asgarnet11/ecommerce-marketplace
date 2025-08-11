package http

import (
	"github.com/gin-gonic/gin"

	// handlers
	"marketplace/backend/internal/handlers"

	// middleware
	"marketplace/backend/internal/middleware"

	// store (untuk tipe *store.DB di NewRouter)
	"marketplace/backend/internal/store"
)

// NewRouter men-setup seluruh route API
func NewRouter(db *store.DB) *gin.Engine {
	r := gin.Default()

	// --- Middleware global ---
	r.Use(middleware.CORS())

	// --- Healthcheck ---
	r.GET("/health", handlers.Health)

	// --- Inisialisasi handler ---
	// public
	authH := handlers.NewAuthHandler(db)
	prodH := handlers.NewProductHandler(db)
	catH  := handlers.NewCategoryHandler(db)
	shopH := handlers.NewShopHandler(db)

	// --- Public routes (/v1) ---
	v1 := r.Group("/v1")
	{
		// Auth
		v1.POST("/auth/register", authH.Register)
		v1.POST("/auth/login",    authH.Login)

		// Products
		v1.GET("/products",       prodH.ListProducts)
		v1.GET("/products/:slug", prodH.GetProduct)

		// Categories
		v1.GET("/categories",       catH.List)  // flat
		v1.GET("/categories/tree",  catH.Tree)  // nested tree

		// Shops
		v1.GET("/shops",               shopH.List)
		v1.GET("/shops/:slug",         shopH.Get)
		v1.GET("/shops/:slug/products", shopH.Products)

		// Contoh endpoint protected sederhana (cek token)
		v1.GET("/me", middleware.AuthRequired(), func(c *gin.Context) {
			c.JSON(200, gin.H{
				"user_id": c.GetInt64("user_id"),
				"role":    c.GetString("role"),
			})
		})

		// --- Protected group (butuh Bearer token) ---
		auth := v1.Group("", middleware.AuthRequired())
		{
			// Cart
			cartH := handlers.NewCartHandler(db)
			auth.GET(   "/cart",        cartH.View)
			auth.POST(  "/cart/items",  cartH.Add)
			auth.PUT(   "/cart/items",  cartH.Update)
			auth.DELETE("/cart/items",  cartH.Remove)

			// Checkout
			checkoutH := handlers.NewCheckoutHandler(db)
			auth.POST("/checkout", checkoutH.Checkout)
		}
	}

	return r
}
