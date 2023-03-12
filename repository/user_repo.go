package repository

import (
	"Tugas-Pert4/entity"
	"context"

	"gorm.io/gorm"
)

// user
type userRepository struct {
	db *gorm.DB
}

type UsrRepo interface {
	//user
	AddUserAccount(ctx context.Context, req *gorm.DB, user entity.User) (entity.User, error)
	FindUserEmail(ctx context.Context, req *gorm.DB, email string) (entity.User, error)
	GetUserAccount(ctx context.Context, req *gorm.DB, user []entity.User) ([]entity.User, error)
	GetUserAccountByID(ctx context.Context, req *gorm.DB, user entity.User, id uint64) (entity.User, error)
	UpdateUserAccount(ctx context.Context, req *gorm.DB, user entity.User, id uint64) (entity.User, error)
	DeleteUserAccountByID(ctx context.Context, req *gorm.DB, user entity.User, id uint64) (entity.User, error)
}

func NewUserRepository(db *gorm.DB) UsrRepo {
	return &userRepository{
		db: db,
	}
}

func (urepo *userRepository) AddUserAccount(ctx context.Context, req *gorm.DB, user entity.User) (entity.User, error) {
	var err error
	if req == nil {
		req = urepo.db.WithContext(ctx).Debug().Create(&user)
		err = req.Error
	} else {
		err = req.WithContext(ctx).Debug().Create(&user).Error
	}

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (urepo *userRepository) GetUserAccount(ctx context.Context, req *gorm.DB, user []entity.User) ([]entity.User, error) {
	var err error
	if req == nil {
		req = urepo.db.WithContext(ctx).Debug().Preload("Blog").Find(&user)
		err = req.Error
	} else {
		err = req.WithContext(ctx).Debug().Preload("Blog").Find(&user).Error
	}

	if err != nil {
		return []entity.User{}, err
	}

	return user, nil
}

func (urepo *userRepository) GetUserAccountByID(ctx context.Context, req *gorm.DB, user entity.User, id uint64) (entity.User, error) {
	var err error
	if req == nil {
		req = urepo.db.WithContext(ctx).Debug().Where("id = ?", id).Find(&user)
		err = req.Error
	} else {
		err = req.WithContext(ctx).Debug().Where("id = ?", id).Find(&user).Error
	}

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (urepo *userRepository) FindUserEmail(ctx context.Context, req *gorm.DB, email string) (entity.User, error) {
	var err error
	var user entity.User
	if req == nil {
		req = urepo.db.WithContext(ctx).Debug().Where("email = ?", email).Find(&user)
		err = req.Error
	} else {
		err = req.WithContext(ctx).Debug().Where("email = ?", email).Find(&user).Error
	}

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (urepo *userRepository) UpdateUserAccount(ctx context.Context, req *gorm.DB, user entity.User, id uint64) (entity.User, error) {
	var err error
	if req == nil {
		req = urepo.db.WithContext(ctx).Debug().Where("id = ?", id).Updates(&user)
		err = req.Error
	} else {
		err = req.WithContext(ctx).Debug().Where("id = ?", id).Updates(&user).Error
	}

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (urepo *userRepository) DeleteUserAccountByID(ctx context.Context, req *gorm.DB, user entity.User, id uint64) (entity.User, error) {
	var err error

	var blog entity.Blog
	var comment entity.Comment

	if req == nil {
		urepo.db.WithContext(ctx).Debug().Where("user_id = ?", id).Delete(&comment)
		urepo.db.WithContext(ctx).Debug().Where("user_id = ?", id).Delete(&blog)
		req = urepo.db.WithContext(ctx).Debug().Where("id = ?", id).Delete(&user)
		err = req.Error
	} else {
		err = req.WithContext(ctx).Debug().Where("id = ?", id).Delete(&user).Error
	}

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
