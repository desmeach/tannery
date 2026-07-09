package auth

import (
	"context"
	"errors"
	"net/http"
	"tanneryApp/internal/application/auth/command"
	"tanneryApp/internal/application/auth/query"
	"tanneryApp/internal/delivery/http/v1/auth/request"
	"tanneryApp/internal/delivery/http/v1/auth/response"
	errs "tanneryApp/internal/domain/errors"

	"github.com/gin-gonic/gin"
)

type Service interface {
	SignUp(ctx context.Context, cmd command.SignUpCommand) error
	SignIn(ctx context.Context, q query.SignInQuery) (*query.SignInQueryResult, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// SignUp register user
// @Summary Register new user account
// @Tags auth
// @Description create user account
// @Accept json
// @Param request body dto.UserSignUpRequest true "Sign up request"
// @Success 201 "ok"
// @Router /auth/sign-up [post]
func (h *Handler) SignUp(c *gin.Context) {
	var req request.SignUpRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid input body",
		})

		return
	}

	cmd := req.ToSignUpCommand()

	if err := h.service.SignUp(c.Request.Context(), cmd); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

// SignIn user login
// @Summary Login user
// @Tags auth
// @Description authenticate user and get token
// @Accept json
// @Produce json
// @Param request body dto.UserSignInRequest true "Sign in request"
// @Success 200 {object} dto.UserSignInResponse
// @Router /auth/sign-in [post]
func (h *Handler) SignIn(c *gin.Context) {
	var req request.SignInRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid input body",
		})
		return
	}

	q := req.ToSignInQuery()

	result, err := h.service.SignIn(c.Request.Context(), *q)
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": errs.ErrUserNotFound.Error(),
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	res := response.ToSignInResponse(*result)

	c.JSON(http.StatusOK, res)
}
