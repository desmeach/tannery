package user

import (
	"context"
	"errors"
	"net/http"
	"tanneryApp/internal/application/user/command"
	"tanneryApp/internal/application/user/query"
	"tanneryApp/internal/delivery/http/v1/user/request"
	"tanneryApp/internal/delivery/http/v1/user/response"
	errs "tanneryApp/internal/domain/errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProfileService interface {
	Create(ctx context.Context, cmd command.CreateProfileCommand) (*command.CreateProfileCommandResult, error)
	FindById(ctx context.Context, q query.FindByIdQuery) (*query.FindByIdQueryResult, error)
	FindByUserId(ctx context.Context, q query.FindByUserIdQuery) (*query.FindByUserIdQueryResult, error)
	UpdateNickname(ctx context.Context, cmd command.UpdateProfileCommand) error
}

type Handler struct {
	profileService ProfileService
}

func NewHandler(profileService ProfileService) *Handler {
	return &Handler{profileService: profileService}
}

// GetCurrentProfile return current user profile
// @Summary Show current user profile
// @Tags Profile
// @Description return current user profile
// @Produce json
// @Success 200 {object} response.GetCurrentProfileResponse
// @Router /user/profile [get]
func (h *Handler) GetCurrentProfile(c *gin.Context) {
	userIdVal, exists := c.Get("userId")
	if !exists {
		c.Status(http.StatusUnauthorized)
		return
	}

	userIdStr, ok := userIdVal.(string)
	if !ok {
		c.Status(http.StatusUnauthorized)
		return
	}

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	q := query.FindByUserIdQuery{UserId: userId}
	profile, err := h.profileService.FindByUserId(c, q)
	if err != nil {
		if errors.Is(err, errs.ErrProfileNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "profile not found",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	res := response.ToGetCurrentProfileResponse(*profile)

	c.JSON(http.StatusOK, res)
}

// CreateProfile create user profile
// @Summary Create user profile
// @Tags Profile
// @Description creates user profile if not exits
// @Accept json
// @Produce json
// @Param request body request.CreateProfileRequest true "User body"
// @Success 201 {object} response.CreateProfileResponse
// @Router /user/profile [post]
func (h *Handler) CreateProfile(c *gin.Context) {
	userIdVal, exists := c.Get("userId")
	if !exists {
		c.Status(http.StatusUnauthorized)
		return
	}

	userIdStr, ok := userIdVal.(string)
	if !ok {
		c.Status(http.StatusUnauthorized)
		return
	}

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	var req request.CreateProfileRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	cmd := req.ToCreateProfileCommand(userId)

	result, err := h.profileService.Create(c.Request.Context(), *cmd)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	res := response.ToCreateProfileResponse(*result)

	c.JSON(http.StatusOK, res)
}
