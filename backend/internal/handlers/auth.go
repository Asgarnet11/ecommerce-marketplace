package handlers

import (
	"net/http"
	"strings"

	"marketplace/backend/internal/config"
	"marketplace/backend/internal/store"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{
	users *store.UserRepo
	pw    store.Password
	jwt   *store.JWT
}

func NewAuthHandler(db *store.DB) *AuthHandler {
	return &AuthHandler{
		users: store.NewUserRepo(db),
		pw:    store.Password{},
		jwt:   store.NewJWT(config.LoadJWT()),
	}
}

type registerReq struct {
	Name     string `json:"name" binding:"required,min=2"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"omitempty,oneof=buyer seller"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var in registerReq
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
	}
	if in.Role == "" { in.Role = "buyer" }

	// cek existing
	if u, _ := h.users.ByEmail(c, strings.ToLower(in.Email)); u != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"email already registered"}); return
	}

	hash, err := h.pw.Hash(in.Password)
	if err != nil { c.JSON(500, gin.H{"error": "hash failed"}); return }

	u, err := h.users.Create(c, in.Name, strings.ToLower(in.Email), in.Role, hash)
	if err != nil { c.JSON(500, gin.H{"error": err.Error()}); return }

	token, exp, err := h.jwt.SignAccess(u.ID, u.Role)
	if err != nil { c.JSON(500, gin.H{"error":"token failed"}); return }

	c.JSON(201, gin.H{
		"user": gin.H{"id": u.ID, "name": u.Name, "email": u.Email, "role": u.Role, "status": u.Status},
		"access_token": token,
		"expires_at": exp.UTC(),
	})
}

type loginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var in loginReq
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
	}
	u, err := h.users.ByEmail(c, strings.ToLower(in.Email))
	if err != nil { c.JSON(http.StatusUnauthorized, gin.H{"error":"invalid credentials"}); return }

	if !h.pw.Verify(in.Password, u.PwHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"invalid credentials"}); return
	}

	token, exp, err := h.jwt.SignAccess(u.ID, u.Role)
	if err != nil { c.JSON(500, gin.H{"error":"token failed"}); return }

	c.JSON(200, gin.H{
		"user": gin.H{"id": u.ID, "name": u.Name, "email": u.Email, "role": u.Role, "status": u.Status},
		"access_token": token,
		"expires_at": exp.UTC(),
	})
}
