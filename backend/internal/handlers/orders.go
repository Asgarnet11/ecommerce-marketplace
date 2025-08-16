package handlers

import (
	"strconv"

	"marketplace/backend/internal/store"

	"github.com/gin-gonic/gin"
)

type OrdersHandler struct{ repo *store.OrderRepo }
func NewOrdersHandler(db *store.DB)*OrdersHandler{ return &OrdersHandler{repo: store.NewOrderRepo(db)} }

func (h *OrdersHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page","1"))
	per,  _ := strconv.Atoi(c.DefaultQuery("per_page","10"))
	status := c.Query("status")
	items, total, err := h.repo.ListByUser(c, c.GetInt64("user_id"), store.OrderListParams{
		Status: status, Page: page, Per: per,
	})
	if err != nil { c.JSON(500, gin.H{"error": err.Error()}); return }
	c.JSON(200, gin.H{"data": items, "page": page, "per_page": per, "total": total})
}

func (h *OrdersHandler) Detail(c *gin.Context) {
	code := c.Param("code")
	d, err := h.repo.DetailByCode(c, c.GetInt64("user_id"), code)
	if err != nil { c.JSON(404, gin.H{"error":"Not Found"}); return }
	c.JSON(200, d)
}
