package handlers

import (
	"marketplace/backend/internal/store"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct{ repo *store.ProductRepo }

func NewProductHandler(db *store.DB) *ProductHandler {
  return &ProductHandler{repo: store.NewProductRepo(db)}
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
  page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
  per, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
  items, total, err := h.repo.List(c, store.ListParams{
    Search: c.Query("search"),
    Sort:   c.Query("sort"),
    Page:   page,
    Per:    per,
  })
  if err != nil { c.JSON(500, gin.H{"error": err.Error()}); return }
  c.JSON(200, gin.H{"data": items, "page": page, "per_page": per, "total": total})
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
  it, err := h.repo.BySlug(c, c.Param("slug"))
  if err != nil { c.JSON(404, gin.H{"error":"Not Found"}); return }
  c.JSON(200, it)
}
