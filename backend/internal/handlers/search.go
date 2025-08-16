package handlers

import (
	"net/http"
	"strconv"

	"marketplace/backend/internal/store"

	"github.com/gin-gonic/gin"
)

type SearchHandler struct{ repo *store.SearchRepo }
func NewSearchHandler(db *store.DB)*SearchHandler{ return &SearchHandler{repo: store.NewSearchRepo(db)} }

func (h *SearchHandler) Global(c *gin.Context) {
	q := c.Query("q")
	lp, _ := strconv.Atoi(c.DefaultQuery("limit_products","12"))
	ls, _ := strconv.Atoi(c.DefaultQuery("limit_shops","6"))

	if q == "" {
		c.JSON(200, gin.H{"products": []any{}, "shops": []any{}})
		return
	}
	res, err := h.repo.Search(c, q, lp, ls)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
	c.JSON(200, res)
}
