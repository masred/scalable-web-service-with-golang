package controller

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/helper"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/model/domain"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/model/http/request"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/model/http/response"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/service"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (userController *UserControllerImpl) Register(ctx *gin.Context) {
	var req request.UserRegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError := err.(validator.ValidationErrors)
		fieldErrorResponse := make(map[string]interface{})

		for _, v := range validationError {
			fieldErrorResponse[strings.ToLower(v.Field())] = helper.GetValidationErrorMsg(v)
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: fieldErrorResponse,
		})

		return
	}

	user := domain.User{
		Age:      req.Age,
		Email:    req.Email,
		Password: req.Password,
		Username: req.Username,
	}

	if err := userController.UserService.Register(&user); err != nil {
		fieldErrorResponse := make(map[string]interface{})

		if strings.Contains(err.Error(), "idx_users_email") {
			fieldErrorResponse["email"] = "Email is already used"
		}

		if strings.Contains(err.Error(), "idx_users_username") {
			fieldErrorResponse["username"] = "Username is already used"
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: fieldErrorResponse,
		})

		return
	}

	ctx.JSON(http.StatusCreated, response.SuccessResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data: response.UserRegisterResponse{
			Age:      user.Age,
			Email:    user.Email,
			ID:       user.ID,
			Username: user.Username,
		},
	})
}

func (userController *UserControllerImpl) Login(ctx *gin.Context) {
	var req request.UserLoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError := err.(validator.ValidationErrors)
		fieldErrorResponse := make(map[string]interface{})

		for _, v := range validationError {
			fieldErrorResponse[strings.ToLower(v.Field())] = helper.GetValidationErrorMsg(v)
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: fieldErrorResponse,
		})

		return
	}

	user := domain.User{
		Email:    req.Email,
		Password: req.Password,
	}

	if err := userController.UserService.Login(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Errors: err.Error(),
		})

		return
	}

	token := helper.GenerateToken(user.ID, user.Email)

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: response.UserLoginResponse{
			Token: token,
		},
	})
}

func (userController *UserControllerImpl) Update(ctx *gin.Context) {
	var req request.UserUpdateRequest

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	id := uint(userData["id"].(float64))

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError := err.(validator.ValidationErrors)
		fieldErrorResponse := make(map[string]interface{})

		for _, v := range validationError {
			fieldErrorResponse[strings.ToLower(v.Field())] = helper.GetValidationErrorMsg(v)
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: fieldErrorResponse,
		})

		return
	}

	user := domain.User{
		ID:       id,
		Username: req.Username,
		Email:    req.Email,
	}

	updatedUser, err := userController.UserService.Update(&user)
	if err != nil {
		fieldErrorResponse := make(map[string]interface{})

		if strings.Contains(err.Error(), "idx_users_email") {
			fieldErrorResponse["email"] = "Email is already used"
		}

		if strings.Contains(err.Error(), "idx_users_username") {
			fieldErrorResponse["username"] = "Username is already used"
		}

		if strings.Contains(err.Error(), "record not found") {
			fieldErrorResponse["id"] = "record not found"
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: fieldErrorResponse,
		})

		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: response.UserUpdateResponse{
			ID:        updatedUser.ID,
			Email:     updatedUser.Email,
			Username:  updatedUser.Username,
			Age:       updatedUser.Age,
			UpdatedAt: updatedUser.UpdatedAt,
		},
	})
}

func (userController *UserControllerImpl) Delete(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	id := uint(userData["id"].(float64))

	if err := userController.UserService.Delete(id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: response.UserDeleteResponse{
			Message: "Your account has been successfully deleted",
		},
	},
	)
}
