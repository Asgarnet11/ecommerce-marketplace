package handlers

import (
	"strconv"

	"marketplace/backend/internal/store"

	"github.com/gin-gonic/gin"
)

type CartHandler struct{ repo *store.CartRepo }
func NewCartHandler(db *store.DB)*CartHandler{ return &CartHandler{repo: store.NewCartRepo(db)} }

func userID(c *gin.Context) int64 { return c.GetInt64("user_id") }

func (h *CartHandler) View(c *gin.Context) {
  items, subtotal, err := h.repo.View(c, userID(c))
  if err!=nil { c.JSON(500, gin.H{"error": err.Error()}); return }
  c.JSON(200, gin.H{"items": items, "subtotal": subtotal})
}

type cartReq struct { ProductID int64 `json:"product_id"`; Qty int32 `json:"qty"` }

func (h *CartHandler) Add(c *gin.Context) {
  var in cartReq
  if err := c.ShouldBindJSON(&in); err!=nil || in.ProductID==0 || in.Qty<=0 {
    c.JSON(400, gin.H{"error":"invalid payload"}); return
  }
  if err := h.repo.UpsertItem(c, userID(c), in.ProductID, in.Qty); err!=nil {
    c.JSON(500, gin.H{"error": err.Error()}); return
  }
  h.View(c)
}

func (h *CartHandler) Update(c *gin.Context) {
  var in cartReq
  if err := c.ShouldBindJSON(&in); err!=nil || in.ProductID==0 {
    c.JSON(400, gin.H{"error":"invalid payload"}); return
  }
  if err := h.repo.UpsertItem(c, userID(c), in.ProductID, in.Qty); err!=nil {
    c.JSON(500, gin.H{"error": err.Error()}); return
  }
  h.View(c)
}

func (h *CartHandler) Remove(c *gin.Context) {
  pid, _ := strconv.ParseInt(c.Query("product_id"),10,64)
  if pid==0 { c.JSON(400, gin.H{"error":"product_id required"}); return }
  if err := h.repo.Remove(c, userID(c), pid); err!=nil {
    c.JSON(500, gin.H{"error": err.Error()}); return
  }
  h.View(c)
}
