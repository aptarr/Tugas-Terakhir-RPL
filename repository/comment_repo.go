package repository

import (
	"Tugas-Pert4/entity"
	"context"

	"gorm.io/gorm"
)

// Comment
type commentRepository struct {
	db *gorm.DB
}

type CmdRepo interface {
	//Comment
	AddUserComment(ctx context.Context, req *gorm.DB, comment entity.Comment) (entity.Comment, error)
	GetUserComment(ctx context.Context, req *gorm.DB, comment []entity.Comment) ([]entity.Comment, error)
	DeleteUserCommentByID(ctx context.Context, req *gorm.DB, comment entity.Comment, id uint64) (entity.Comment, error)
	UpdateUserComment(ctx context.Context, req *gorm.DB, comment entity.Comment, id uint64) (entity.Comment, error)
}

func NewCommentRepository(db *gorm.DB) CmdRepo {
	return &commentRepository{
		db: db,
	}
}

// COMMENT
func (crepo *commentRepository) AddUserComment(ctx context.Context, req *gorm.DB, comment entity.Comment) (entity.Comment, error) {
	var err error
	if req == nil {
		req = crepo.db.WithContext(ctx).Debug().Create(&comment)
		err = req.Error
	} else {
		err = req.WithContext(ctx).Debug().Create(&comment).Error
	}

	if err != nil {
		return entity.Comment{}, err
	}

	return comment, nil
}

func (crepo *commentRepository) GetUserComment(ctx context.Context, req *gorm.DB, comment []entity.Comment) ([]entity.Comment, error) {
	var err error
	if req == nil {
		req = crepo.db.WithContext(ctx).Debug().Find(&comment)
		err = req.Error
	} else {
		err = req.WithContext(ctx).Debug().Find(&comment).Error
	}

	if err != nil {
		return []entity.Comment{}, err
	}

	return comment, nil
}

func (crepo *commentRepository) DeleteUserCommentByID(ctx context.Context, req *gorm.DB, comment entity.Comment, id uint64) (entity.Comment, error) {
	var err error

	if req == nil {
		req = crepo.db.WithContext(ctx).Debug().Where("id = ?", id).Delete(&comment)
		err = req.Error
	} else {
		err = req.WithContext(ctx).Debug().Where("id = ?", id).Delete(&comment).Error
	}

	if err != nil {
		return entity.Comment{}, err
	}

	return comment, nil
}

func (crepo *commentRepository) UpdateUserComment(ctx context.Context, req *gorm.DB, comment entity.Comment, id uint64) (entity.Comment, error) {
	var err error
	if req == nil {
		req = crepo.db.WithContext(ctx).Debug().Where("id = ?", id).Updates(&comment)
		err = req.Error
	} else {
		err = req.WithContext(ctx).Debug().Where("id = ?", id).Updates(&comment).Error
	}

	if err != nil {
		return entity.Comment{}, err
	}

	return comment, nil
}
