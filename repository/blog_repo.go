package repository

import (
	"Tugas-Pert4/entity"
	"context"

	"gorm.io/gorm"
)

// Blog
type blogRepository struct {
	db *gorm.DB
}

type BlgRepo interface {
	//blog
	AddUserBlog(ctx context.Context, req *gorm.DB, blog entity.Blog) (entity.Blog, error)
	GetUserBlog(ctx context.Context, req *gorm.DB, blog []entity.Blog) ([]entity.Blog, error)
	GetUserBlogDetail(ctx context.Context, req *gorm.DB, blog []entity.Blog) ([]entity.Blog, error)
	DeleteUserBlogByID(ctx context.Context, req *gorm.DB, blog entity.Blog, id uint64) (entity.Blog, error)
	DeleteUserBlogByIDUser(ctx context.Context, req *gorm.DB, blog entity.Blog, id uint64) (entity.Blog, error)
	UpdateUserBlog(ctx context.Context, req *gorm.DB, blog entity.Blog, id uint64) (entity.Blog, error)
}

func NewBlogRepository(db *gorm.DB) BlgRepo {
	return &blogRepository{
		db: db,
	}
}

// BLOG
func (brepo *blogRepository) AddUserBlog(ctx context.Context, req *gorm.DB, blog entity.Blog) (entity.Blog, error) {
	var err error
	if req == nil {
		req = brepo.db.WithContext(ctx).Debug().Create(&blog)
		err = req.Error
	} else {
		err = req.WithContext(ctx).Debug().Create(&blog).Error
	}

	if err != nil {
		return entity.Blog{}, err
	}

	return blog, nil
}

func (brepo *blogRepository) GetUserBlog(ctx context.Context, req *gorm.DB, blog []entity.Blog) ([]entity.Blog, error) {
	var err error
	if req == nil {
		req = brepo.db.WithContext(ctx).Debug().Find(&blog)
		err = req.Error
	} else {
		err = req.WithContext(ctx).Debug().Find(&blog).Error
	}

	if err != nil {
		return []entity.Blog{}, err
	}

	return blog, nil
}

func (brepo *blogRepository) GetUserBlogDetail(ctx context.Context, req *gorm.DB, blog []entity.Blog) ([]entity.Blog, error) {
	var err error
	if req == nil {
		req = brepo.db.WithContext(ctx).Debug().Preload("Comment").Find(&blog)
		err = req.Error
	} else {
		err = req.WithContext(ctx).Debug().Preload("Comment").Find(&blog).Error
	}

	if err != nil {
		return []entity.Blog{}, err
	}

	return blog, nil
}

func (brepo *blogRepository) DeleteUserBlogByID(ctx context.Context, req *gorm.DB, blog entity.Blog, id uint64) (entity.Blog, error) {
	var err error
	var comment entity.Comment
	if req == nil {
		brepo.db.WithContext(ctx).Debug().Where("user_id = ?", id).Delete(&comment)
		req = brepo.db.WithContext(ctx).Debug().Where("user_id = ?", id).Delete(&blog)
		err = req.Error
	} else {
		err = req.WithContext(ctx).Debug().Where("user_id = ?", id).Delete(&blog).Error
	}

	if err != nil {
		return entity.Blog{}, err
	}

	return blog, nil
}

func (brepo *blogRepository) DeleteUserBlogByIDUser(ctx context.Context, req *gorm.DB, blog entity.Blog, id uint64) (entity.Blog, error) {
	var err error
	var comment entity.Comment
	if req == nil {
		brepo.db.WithContext(ctx).Debug().Where("blog_id = ?", id).Delete(&comment)
		req = brepo.db.WithContext(ctx).Debug().Where("id = ?", id).Delete(&blog)
		err = req.Error
	} else {
		err = req.WithContext(ctx).Debug().Where("id = ?", id).Delete(&blog).Error
	}

	if err != nil {
		return entity.Blog{}, err
	}

	return blog, nil
}

func (brepo *blogRepository) UpdateUserBlog(ctx context.Context, req *gorm.DB, blog entity.Blog, id uint64) (entity.Blog, error) {
	var err error
	if req == nil {
		req = brepo.db.WithContext(ctx).Debug().Where("id = ?", id).Updates(&blog)
		err = req.Error
	} else {
		err = req.WithContext(ctx).Debug().Where("id = ?", id).Updates(&blog).Error
	}

	if err != nil {
		return entity.Blog{}, err
	}

	return blog, nil
}
