package handlers

import "github.com/gin-gonic/gin"

var sampleProducts = []gin.H{
	{"id": 1, "shop_id": 1, "title": "Beras Premium 5kg", "slug": "beras-premium-5kg", "price": 75000, "stock": 120},
	{"id": 2, "shop_id": 1, "title": "Minyak Goreng 2L", "slug": "minyak-goreng-2l", "price": 38000, "stock": 50},
}

func ListProducts(c *gin.Context) {
	c.JSON(200, gin.H{"data": sampleProducts, "page": 1, "per_page": 20, "total": len(sampleProducts)})
}

func GetProduct(c *gin.Context) {
	slug := c.Param("slug")
	for _, p := range sampleProducts {
		if p["slug"] == slug { c.JSON(200, p); return }
	}
	c.JSON(404, gin.H{"error": "Not Found"})
}
