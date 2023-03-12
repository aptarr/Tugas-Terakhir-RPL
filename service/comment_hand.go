package service

import (
	"Tugas-Pert4/dto"
	"Tugas-Pert4/entity"
	"Tugas-Pert4/repository"
	"context"

	"github.com/jinzhu/copier"
)

type commentHandler struct {
	commentRepo repository.CmdRepo
}

type CmdHndlr interface {
	//Comment
	AddCommentUser(ctx context.Context, userDTO dto.UserAddComment) (entity.Comment, error)
	GetCommentUser(ctx context.Context, user []entity.Comment) ([]entity.Comment, error)
	DeleteCommentUserByID(ctx context.Context, id uint64) (entity.Comment, error)
	UpdateCommentUser(ctx context.Context, userDTO dto.UserUpdateComment, id uint64) (entity.Comment, error)
}

func NewCommentHandler(cmd repository.CmdRepo) CmdHndlr {
	return &commentHandler{
		commentRepo: cmd,
	}
}

func (hndlr *commentHandler) AddCommentUser(ctx context.Context, commenDTO dto.UserAddComment) (entity.Comment, error) {
	var comment entity.Comment
	copier.Copy(&comment, &commenDTO)

	addComment, erraddComment := hndlr.commentRepo.AddUserComment(ctx, nil, comment)
	if erraddComment != nil {
		return entity.Comment{}, erraddComment
	}

	return addComment, nil
}

func (hndlr *commentHandler) GetCommentUser(ctx context.Context, comment []entity.Comment) ([]entity.Comment, error) {
	getComment, errgetComment := hndlr.commentRepo.GetUserComment(ctx, nil, comment)
	if errgetComment != nil {
		return []entity.Comment{}, errgetComment
	}

	return getComment, nil
}

func (hndlr *commentHandler) DeleteCommentUserByID(ctx context.Context, id uint64) (entity.Comment, error) {
	var comment entity.Comment
	delComment, errdelComment := hndlr.commentRepo.DeleteUserCommentByID(ctx, nil, comment, id)
	if errdelComment != nil {
		return entity.Comment{}, errdelComment
	}

	return delComment, nil
}

func (hndlr *commentHandler) UpdateCommentUser(ctx context.Context, commenDTO dto.UserUpdateComment, id uint64) (entity.Comment, error) {
	var comment entity.Comment
	copier.Copy(&comment, &commenDTO)

	updateComment, errupdateComment := hndlr.commentRepo.UpdateUserComment(ctx, nil, comment, id)
	if errupdateComment != nil {
		return entity.Comment{}, errupdateComment
	}

	return updateComment, nil
}
