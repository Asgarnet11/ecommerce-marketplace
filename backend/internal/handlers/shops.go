package handlers

import (
	"net/http"
	"strconv"

	"marketplace/backend/internal/store"

	"github.com/gin-gonic/gin"
)


type ShopHandler struct{ repo *store.ShopRepo }

func NewShopHandler(db *store.DB) *ShopHandler {
	return &ShopHandler{repo: store.NewShopRepo(db)}
}

func (h *ShopHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per,  _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))

	items, total, err := h.repo.List(c, store.ShopListParams{Page: page, Per: per})
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
	c.JSON(200, gin.H{"data": items, "page": page, "per_page": per, "total": total})
}

func (h *ShopHandler) Get(c *gin.Context) {
	slug := c.Param("slug")
	sh, err := h.repo.BySlug(c, slug)
	if err != nil { c.JSON(http.StatusNotFound, gin.H{"error":"Not Found"}); return }
	c.JSON(200, sh)
}

func (h *ShopHandler) Products(c *gin.Context) {
	slug := c.Param("slug")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per,  _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	sort := c.Query("sort")
	search := c.Query("search")

	repo := store.NewProductRepo(h.repo.DB) // pakai DB yang sama
	items, total, err := repo.ListByShopSlug(c, slug, store.ListParams{
		Search: search, Sort: sort, Page: page, Per: per,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": items, "page": page, "per_page": per, "total": total})
}

