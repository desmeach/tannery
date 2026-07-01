package auth

import (
	"context"
	"net/http"
	"tanneryApp/internal/application/auth"
	"tanneryApp/internal/delivery/http/v1/auth/dto"

	"github.com/gin-gonic/gin"
)

type Service interface {
	SignUp(ctx context.Context, login, password string) error
	SignIn(ctx context.Context, login, password string) (auth.SignInResult, error)
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
// @Produce json
// @Param request body dto.UserSignUpRequest true "Sign up request"
// @Success 201 {object} map[string]interface{} "ok"
// @Router /auth/sign-up [post]
func (h *Handler) SignUp(c *gin.Context) {
	var input dto.UserSignUpRequest
	if err := c.BindJSON(&input); err != nil {
		//newResponse(c, http.StatusBadRequest, "invalid input body")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid input body",
		})

		return
	}

	if err := h.service.SignUp(c.Request.Context(), input.Login, input.Password); err != nil {
		//if errors.Is(err, domain.ErrUserAlreadyExists) {
		//	newResponse(c, http.StatusBadRequest, err.Error())
		//
		//	return
		//}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		//newResponse(c, http.StatusInternalServerError, err.Error())

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
	var input dto.UserSignInRequest
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid input body",
		})
		return
	}

	response, err := h.service.SignIn(c.Request.Context(), input.Login, input.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.UserSignInResponse{
		UserId:      response.UserId,
		AccessToken: response.AccessToken,
	})
}
