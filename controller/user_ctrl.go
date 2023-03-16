package controller

import (
	"Tugas-Pert4/dto"
	"Tugas-Pert4/entity"
	"Tugas-Pert4/midleware"
	"Tugas-Pert4/service"
	"Tugas-Pert4/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userReq service.UserHndlr
}

type UserCtrl interface {
	//user
	AddAccount(ctx *gin.Context)
	LoginAccount(ctx *gin.Context)
	ValidateAccount(ctx *gin.Context)

	GetAccount(ctx *gin.Context)
	GetAccountByID(ctx *gin.Context)
	UpdateAccount(ctx *gin.Context)
	DeleteAccount(ctx *gin.Context)
}

func NewUserController(usr service.UserHndlr) UserCtrl {
	return &userController{
		userReq: usr,
	}
}

// AddAccount
func (ctrl *userController) AddAccount(ctx *gin.Context) {
	var userDTO dto.UserAddAccount
	errUserDTO := ctx.ShouldBind(&userDTO)
	if errUserDTO != nil {
		response := utils.ErrorResponse("Failed Request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, errUser := ctrl.userReq.AddAccountUser(ctx, userDTO)
	if errUser != nil {
		response := utils.ErrorResponse("Failed Add Account", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.SuccessResponse("Account Added", http.StatusCreated, user)
	ctx.AbortWithStatusJSON(http.StatusCreated, response)
}

// LoginAccount
func (ctrl *userController) LoginAccount(ctx *gin.Context) {
	var userDTO dto.UserAddAccount
	errUserDTO := ctx.ShouldBind(&userDTO)
	if errUserDTO != nil {
		response := utils.ErrorResponse("Failed Request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tokens, errUser := ctrl.userReq.GetEmailUser(ctx, userDTO)
	if errUser != nil {
		response := utils.ErrorResponse("Failed Login Account", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	ctx.Header("Authorization", tokens)

	response := utils.SuccessResponse("Login Sucsefull", http.StatusCreated, tokens)
	ctx.AbortWithStatusJSON(http.StatusCreated, response)
}

// ValidateAccount
func (ctrl *userController) ValidateAccount(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := utils.ErrorResponse("ID must Integer", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, errUser := ctrl.userReq.GetAccountUserByID(ctx, id)
	if errUser != nil {
		response := utils.ErrorResponse("Failed Get Account", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	jwt_new := service.NewJWTService()
	tokens, err := jwt_new.ValidateToken(ctx.GetHeader("Authorization"))

	if err != nil {
		response := utils.ErrorResponse("Failed Validate Account", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	role, err := jwt_new.GetRoleByToken(tokens.Raw)

	if err != nil {
		response := utils.ErrorResponse("Failed Get Role Account", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	midleware.Authenticate(jwt_new, role)

	response := utils.SuccessResponse("I'm logged in", http.StatusCreated, user)
	ctx.AbortWithStatusJSON(http.StatusCreated, response)
}

// GetAccount
func (ctrl *userController) GetAccount(ctx *gin.Context) {
	var userDTO []entity.User

	user, errUser := ctrl.userReq.GetAccountUser(ctx, userDTO)
	if errUser != nil {
		response := utils.ErrorResponse("Failed Get Account", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.SuccessResponse("Users Fetched Successfully", http.StatusCreated, user)
	ctx.AbortWithStatusJSON(http.StatusCreated, response)
}

// GetAccountByID
func (ctrl *userController) GetAccountByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := utils.ErrorResponse("ID must Integer", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, errUser := ctrl.userReq.GetAccountUserByID(ctx, id)
	if errUser != nil {
		response := utils.ErrorResponse("Failed Get Account", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.SuccessResponse("Users Fetched Successfully", http.StatusCreated, user)
	ctx.AbortWithStatusJSON(http.StatusCreated, response)
}

// UpdateAccount
func (ctrl *userController) UpdateAccount(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := utils.ErrorResponse("ID must Integer", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var userDTO dto.UserUpdateAccount
	errUserDTO := ctx.ShouldBind(&userDTO)
	if errUserDTO != nil {
		response := utils.ErrorResponse("Failed Request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, errUser := ctrl.userReq.UpdateAccountUser(ctx, userDTO, id)
	if errUser != nil {
		response := utils.ErrorResponse("Failed Update Account", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.SuccessResponse("Account Updated", http.StatusCreated, user)
	ctx.AbortWithStatusJSON(http.StatusCreated, response)
}

// DeleteAccount
func (ctrl *userController) DeleteAccount(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := utils.ErrorResponse("ID must Integer", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, errUser := ctrl.userReq.DeleteAccountUserByID(ctx, id)
	if errUser != nil {
		response := utils.ErrorResponse("Failed Delete Account", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.SuccessResponse("Account Deleted", http.StatusCreated, user)
	ctx.AbortWithStatusJSON(http.StatusCreated, response)
}
