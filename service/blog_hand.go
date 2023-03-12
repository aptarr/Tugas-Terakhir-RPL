package service

import (
	"Tugas-Pert4/dto"
	"Tugas-Pert4/entity"
	"Tugas-Pert4/repository"
	"context"

	"github.com/jinzhu/copier"
)

type blogHandler struct {
	blogRepo repository.BlgRepo
}

type BlgHndlr interface {
	//blog
	AddBlogUser(ctx context.Context, userDTO dto.UserAddBlog) (entity.Blog, error)
	GetBlogUser(ctx context.Context, user []entity.Blog) ([]entity.Blog, error)
	GetBlogUserDetail(ctx context.Context, user []entity.Blog) ([]entity.Blog, error)
	DeleteBlogUserByID(ctx context.Context, id uint64) (entity.Blog, error)
	DeleteBlogUserByIDUser(ctx context.Context, id uint64) (entity.Blog, error)
	UpdateBlogUser(ctx context.Context, userDTO dto.UserUpdateBlog, id uint64) (entity.Blog, error)
}

func NewBlogHandler(blg repository.BlgRepo) BlgHndlr {
	return &blogHandler{
		blogRepo: blg,
	}
}

func (hndlr *blogHandler) AddBlogUser(ctx context.Context, blogDTO dto.UserAddBlog) (entity.Blog, error) {
	var blog entity.Blog
	copier.Copy(&blog, &blogDTO)

	addBlog, errAddBlog := hndlr.blogRepo.AddUserBlog(ctx, nil, blog)
	if errAddBlog != nil {
		return entity.Blog{}, errAddBlog
	}

	return addBlog, nil
}

func (hndlr *blogHandler) GetBlogUser(ctx context.Context, blog []entity.Blog) ([]entity.Blog, error) {
	getBlog, errgetBlog := hndlr.blogRepo.GetUserBlog(ctx, nil, blog)
	if errgetBlog != nil {
		return []entity.Blog{}, errgetBlog
	}

	return getBlog, nil
}

func (hndlr *blogHandler) GetBlogUserDetail(ctx context.Context, blog []entity.Blog) ([]entity.Blog, error) {
	getBlog, errgetBlog := hndlr.blogRepo.GetUserBlogDetail(ctx, nil, blog)
	if errgetBlog != nil {
		return []entity.Blog{}, errgetBlog
	}

	return getBlog, nil
}

func (hndlr *blogHandler) DeleteBlogUserByID(ctx context.Context, id uint64) (entity.Blog, error) {
	var blog entity.Blog
	delBlog, errdelBlog := hndlr.blogRepo.DeleteUserBlogByID(ctx, nil, blog, id)
	if errdelBlog != nil {
		return entity.Blog{}, errdelBlog
	}

	return delBlog, nil
}

func (hndlr *blogHandler) DeleteBlogUserByIDUser(ctx context.Context, id uint64) (entity.Blog, error) {
	var blog entity.Blog
	delBlog, errdelBlog := hndlr.blogRepo.DeleteUserBlogByIDUser(ctx, nil, blog, id)
	if errdelBlog != nil {
		return entity.Blog{}, errdelBlog
	}

	return delBlog, nil
}

func (hndlr *blogHandler) UpdateBlogUser(ctx context.Context, blogDTO dto.UserUpdateBlog, id uint64) (entity.Blog, error) {
	var blog entity.Blog
	copier.Copy(&blog, &blogDTO)

	updateBlog, errupdateBlog := hndlr.blogRepo.UpdateUserBlog(ctx, nil, blog, id)
	if errupdateBlog != nil {
		return entity.Blog{}, errupdateBlog
	}

	return updateBlog, nil
}
