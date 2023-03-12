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

type blogController struct {
	blogReq service.BlgHndlr
}

type BlogCtrl interface {
	//blog
	AddBlog(ctx *gin.Context)
	GetBlog(ctx *gin.Context)
	GetBlogDetail(ctx *gin.Context)
	DeleteBlogUser(ctx *gin.Context)
	DeleteBlog(ctx *gin.Context)
	UpdateBlog(ctx *gin.Context)
}

func NewBlogController(usr service.BlgHndlr) BlogCtrl {
	return &blogController{
		blogReq: usr,
	}
}

// AddBlog
func (ctrl *blogController) AddBlog(ctx *gin.Context) {
	var blogDTO dto.UserAddBlog
	errBlogDTO := ctx.ShouldBind(&blogDTO)
	if errBlogDTO != nil {
		response := utils.ErrorResponse("Failed Request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	blog, errBlog := ctrl.blogReq.AddBlogUser(ctx, blogDTO)
	if errBlog != nil {
		response := utils.ErrorResponse("Failed Add Blog", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.SuccessResponse("Blog Added", http.StatusCreated, blog)
	ctx.AbortWithStatusJSON(http.StatusCreated, response)
}

// GetBlog
func (ctrl *blogController) GetBlog(ctx *gin.Context) {
	var blogs []entity.Blog
	blog, errBlog := ctrl.blogReq.GetBlogUser(ctx, blogs)
	if errBlog != nil {
		response := utils.ErrorResponse("Failed Get Account", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.SuccessResponse("Blog Fetched Successfully", http.StatusCreated, blog)
	ctx.AbortWithStatusJSON(http.StatusCreated, response)
}

// GetBlogDetail
func (ctrl *blogController) GetBlogDetail(ctx *gin.Context) {
	var blog []entity.Blog
	blog, errUser := ctrl.blogReq.GetBlogUserDetail(ctx, blog)
	if errUser != nil {
		response := utils.ErrorResponse("Failed Get Account", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.SuccessResponse("Blog Fetched Successfully", http.StatusCreated, blog)
	ctx.AbortWithStatusJSON(http.StatusCreated, response)
}

// DeleteBlog
func (ctrl *blogController) DeleteBlogUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := utils.ErrorResponse("ID must Integer", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	blog, errBlog := ctrl.blogReq.DeleteBlogUserByID(ctx, id)
	if errBlog != nil {
		response := utils.ErrorResponse("Failed Delete Blog", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.SuccessResponse("Blog Deleted", http.StatusCreated, blog)
	ctx.AbortWithStatusJSON(http.StatusCreated, response)
}

func (ctrl *blogController) DeleteBlog(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := utils.ErrorResponse("ID must Integer", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	blog, errBlog := ctrl.blogReq.DeleteBlogUserByIDUser(ctx, id)
	if errBlog != nil {
		response := utils.ErrorResponse("Failed Delete Blog", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.SuccessResponse("Blog Deleted", http.StatusCreated, blog)
	ctx.AbortWithStatusJSON(http.StatusCreated, response)
}

// UpdateAccount
func (ctrl *blogController) UpdateBlog(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := utils.ErrorResponse("ID must Integer", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var blogDTO dto.UserUpdateBlog
	errUserDTO := ctx.ShouldBind(&blogDTO)
	if errUserDTO != nil {
		response := utils.ErrorResponse("Failed Request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	blog, errBlog := ctrl.blogReq.UpdateBlogUser(ctx, blogDTO, id)
	if errBlog != nil {
		response := utils.ErrorResponse("Failed Update Blog", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.SuccessResponse("Blog Updated", http.StatusCreated, blog)
	ctx.AbortWithStatusJSON(http.StatusCreated, response)
}
