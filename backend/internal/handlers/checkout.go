package handlers

import (
	"net/http"

	"marketplace/backend/internal/store"

	"github.com/gin-gonic/gin"
)

type CheckoutHandler struct{ db *store.DB }
func NewCheckoutHandler(db *store.DB)*CheckoutHandler{ return &CheckoutHandler{db: db} }

type checkoutReq struct {
  Address struct {
    Label string `json:"label"`
    Recipient string `json:"recipient"`
    Phone string `json:"phone"`
    Line1 string `json:"line1"`
    City string `json:"city"`
    Province string `json:"province"`
    PostalCode string `json:"postal_code"`
    Notes string `json:"notes"`
  } `json:"address"`
  Courier string  `json:"courier"` // "Dummy" untuk MVP
  Service string  `json:"service"` // "REG"
  Note    *string `json:"note"`
}

func (h *CheckoutHandler) Checkout(c *gin.Context) {
  var in checkoutReq
  if err := c.ShouldBindJSON(&in); err!=nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
  }
  // convert address ke map
  addr := map[string]any{
    "label": in.Address.Label, "recipient": in.Address.Recipient, "phone": in.Address.Phone,
    "line1": in.Address.Line1, "city": in.Address.City, "province": in.Address.Province,
    "postal_code": in.Address.PostalCode, "notes": in.Address.Notes,
    "courier": in.Courier, "service": in.Service,
  }
  res, err := h.db.Checkout(c, store.CheckoutInput{
    UserID:  c.GetInt64("user_id"),
    Address: addr, Courier: in.Courier, Service: in.Service, Note: in.Note,
  })
  if err!=nil { c.JSON(400, gin.H{"error": err.Error()}); return }
  c.JSON(201, res)
}
