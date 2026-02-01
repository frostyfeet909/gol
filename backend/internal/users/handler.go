package users

import (
	"api/types"
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

type createReq struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name"  validate:"required,min=2"`
}

// Create godoc
// @Summary      Create a new user
// @Description  Create a new user with email and name
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      createReq  true  "User info"
// @Success      201   {object}  users.User
// @Failure      400   {object}  types.ErrorResponse
// @Failure      409   {object}  types.ErrorResponse
// @Router       /v1/users [post]
func (h *Handler) create(c *echo.Context) error {
	var req createReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, &types.ErrorResponse{Message: err.Error()})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, &types.ErrorResponse{Message: err.Error()})
	}

	u, err := h.svc.create(c.Request().Context(), req.Email, req.Name)
	if err != nil {
		return c.JSON(mapErr(err), &types.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, u)
}

// Get godoc
// @Summary      Get a user
// @Description  Fetch user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  users.User
// @Failure      404  {object}  types.ErrorResponse
// @Router       /v1/users/{id} [get]
func (h *Handler) get(c *echo.Context) error {
	id := c.Param("id")

	u, err := h.svc.get(c.Request().Context(), id)
	if err != nil {
		return c.JSON(mapErr(err), &types.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, u)
}

// List godoc
// @Summary      List users
// @Description  Fetch all users
// @Tags         users
// @Produce      json
// @Success      200  {array}   users.User
// @Failure      404  {object}  types.ErrorResponse
// @Router       /v1/users [get]
func (h *Handler) list(c *echo.Context) error {
	us, err := h.svc.list(c.Request().Context())
	if err != nil {
		return c.JSON(mapErr(err), &types.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, us)

}

// Delete godoc
// @Summary      Delete a user
// @Description  Remove user by ID
// @Tags         users
// @Param        id   path      string  true  "User ID"
// @Success      204
// @Failure      404  {object}  types.ErrorResponse
// @Router       /v1/users/{id} [delete]
func (h *Handler) delete(c *echo.Context) error {
	id := c.Param("id")
	err := h.svc.remove(c.Request().Context(), id)
	if err != nil {
		return c.JSON(mapErr(err), &types.ErrorResponse{Message: err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
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
