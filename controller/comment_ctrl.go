package controller

import (
	"Tugas-Pert4/dto"
	"Tugas-Pert4/entity"
	"Tugas-Pert4/service"
	"Tugas-Pert4/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type commentController struct {
	commentReq service.CmdHndlr
}

type CmdCtrl interface {
	//blog
	AddComment(ctx *gin.Context)
	GetComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
}

func NewCommentController(cmd service.CmdHndlr) CmdCtrl {
	return &commentController{
		commentReq: cmd,
	}
}

// AddComment
func (ctrl *commentController) AddComment(ctx *gin.Context) {
	var commenDTO dto.UserAddComment
	errcommenDTO := ctx.ShouldBind(&commenDTO)
	if errcommenDTO != nil {
		response := utils.ErrorResponse("Failed Request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	comment, errCommend := ctrl.commentReq.AddCommentUser(ctx, commenDTO)
	if errCommend != nil {
		response := utils.ErrorResponse("Failed Add Comment", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.SuccessResponse("Comment Added", http.StatusCreated, comment)
	ctx.AbortWithStatusJSON(http.StatusCreated, response)
}

// GetComment
func (ctrl *commentController) GetComment(ctx *gin.Context) {
	var comments []entity.Comment
	comment, errCommend := ctrl.commentReq.GetCommentUser(ctx, comments)
	if errCommend != nil {
		response := utils.ErrorResponse("Failed Get Comment", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.SuccessResponse("Comment Fetched Successfully", http.StatusCreated, comment)
	ctx.AbortWithStatusJSON(http.StatusCreated, response)
}

// DeleteComment
func (ctrl *commentController) DeleteComment(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := utils.ErrorResponse("ID must Integer", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	comment, errCommend := ctrl.commentReq.DeleteCommentUserByID(ctx, id)
	if errCommend != nil {
		response := utils.ErrorResponse("Failed Delete Comment", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.SuccessResponse("Comment Deleted", http.StatusCreated, comment)
	ctx.AbortWithStatusJSON(http.StatusCreated, response)
}

// UpdateComment
func (ctrl *commentController) UpdateComment(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := utils.ErrorResponse("ID must Integer", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var commenDTO dto.UserUpdateComment
	errUserDTO := ctx.ShouldBind(&commenDTO)
	if errUserDTO != nil {
		response := utils.ErrorResponse("Failed Request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	comment, errCommend := ctrl.commentReq.UpdateCommentUser(ctx, commenDTO, id)
	if errCommend != nil {
		response := utils.ErrorResponse("Failed Update Comment", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.SuccessResponse("Comment Updated", http.StatusCreated, comment)
	ctx.AbortWithStatusJSON(http.StatusCreated, response)
}
