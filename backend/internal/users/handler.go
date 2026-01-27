package users

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

type createReq struct {
	Email string `json:"email" binding:"required,email"`
	Name  string `json:"name"  binding:"required,min=2"`
}

// Create godoc
// @Summary      Create a new user
// @Description  Create a new user with email and name
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      createReq  true  "User info"
// @Success      201   {object}  users.User
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Router       /users [post]
func (h *Handler) Create(c *gin.Context) {
	var req createReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.svc.Create(c.Request.Context(), req.Email, req.Name)
	if err != nil {
		c.JSON(mapErr(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    u.ID,
		"email": u.Email,
		"name":  u.Name,
	})
}

// Get godoc
// @Summary      Get a user
// @Description  Fetch user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  users.User
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [get]
func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")

	u, err := h.svc.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(mapErr(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    u.ID,
		"email": u.Email,
		"name":  u.Name,
	})
}

func mapErr(err error) int {
	switch {
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, ErrEmailTaken):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
