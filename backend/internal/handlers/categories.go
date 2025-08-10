package handlers

import (
	"marketplace/backend/internal/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct{ repo *store.CategoryRepo }

func NewCategoryHandler(db *store.DB) *CategoryHandler {
	return &CategoryHandler{repo: store.NewCategoryRepo(db)}
}

func (h *CategoryHandler) List(c *gin.Context) {
	items, err := h.repo.List(c)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
	c.JSON(200, items)
}

func (h *CategoryHandler) Tree(c *gin.Context) {
	items, err := h.repo.ListTree(c)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
	c.JSON(200, items)
}
