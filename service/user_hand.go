package service

import (
	"Tugas-Pert4/dto"
	"Tugas-Pert4/entity"
	"Tugas-Pert4/repository"
	"Tugas-Pert4/utils"
	"context"
	"errors"

	"github.com/jinzhu/copier"
)

type userHandler struct {
	userRepo repository.UsrRepo
}

type UserHndlr interface {
	//user
	AddAccountUser(ctx context.Context, userDTO dto.UserAddAccount) (entity.User, error)
	GetEmailUser(ctx context.Context, userDTO dto.UserAddAccount) (string, error)
	GetAccountUser(ctx context.Context, userDTO []entity.User) ([]entity.User, error)
	GetAccountUserByID(ctx context.Context, id uint64) (entity.User, error)
	UpdateAccountUser(ctx context.Context, userDTO dto.UserUpdateAccount, id uint64) (entity.User, error)
	DeleteAccountUserByID(ctx context.Context, id uint64) (entity.User, error)
}

func NewUserHandler(usr repository.UsrRepo) UserHndlr {
	return &userHandler{
		userRepo: usr,
	}
}

func (hndlr *userHandler) AddAccountUser(ctx context.Context, userDTO dto.UserAddAccount) (entity.User, error) {
	var user entity.User
	copier.Copy(&user, &userDTO)

	//check duplicate email
	emailUser, err := hndlr.userRepo.FindUserEmail(ctx, nil, user.Email)
	if err != nil {
		return entity.User{}, err
	}

	if emailUser.Email == userDTO.Email {
		return entity.User{}, errors.New("email has been used")
	}

	addAccount, errAddAccount := hndlr.userRepo.AddUserAccount(ctx, nil, user)
	if errAddAccount != nil {
		return entity.User{}, errAddAccount
	}

	return addAccount, nil
}

func (hndlr *userHandler) GetEmailUser(ctx context.Context, userDTO dto.UserAddAccount) (string, error) {
	var user entity.User
	copier.Copy(&user, &userDTO)

	//check email
	emailUser, err := hndlr.userRepo.FindUserEmail(ctx, nil, user.Email)
	if err != nil {
		return "", errors.New("invalid email")
	}

	if utils.PasswordCompare(emailUser.Password, []byte(userDTO.Password)) != nil {
		return "", errors.New("invalid pasword")
	}

	jwt_new := NewJWTService()
	tokens := jwt_new.GenerateToken(emailUser.ID, "admin")

	return tokens, nil
}

func (hndlr *userHandler) GetAccountUser(ctx context.Context, userDTO []entity.User) ([]entity.User, error) {
	getAccount, errAddAccount := hndlr.userRepo.GetUserAccount(ctx, nil, userDTO)
	if errAddAccount != nil {
		return []entity.User{}, errAddAccount
	}

	return getAccount, nil
}

func (hndlr *userHandler) GetAccountUserByID(ctx context.Context, id uint64) (entity.User, error) {
	var user entity.User
	getAccount, errAddAccount := hndlr.userRepo.GetUserAccountByID(ctx, nil, user, id)
	if errAddAccount != nil {
		return entity.User{}, errAddAccount
	}

	return getAccount, nil
}

func (hndlr *userHandler) UpdateAccountUser(ctx context.Context, userDTO dto.UserUpdateAccount, id uint64) (entity.User, error) {
	var user entity.User
	copier.Copy(&user, &userDTO)

	addAccount, errAddAccount := hndlr.userRepo.UpdateUserAccount(ctx, nil, user, id)
	if errAddAccount != nil {
		return entity.User{}, errAddAccount
	}

	return addAccount, nil
}

func (hndlr *userHandler) DeleteAccountUserByID(ctx context.Context, id uint64) (entity.User, error) {
	var user entity.User
	getAccount, errAddAccount := hndlr.userRepo.DeleteUserAccountByID(ctx, nil, user, id)
	if errAddAccount != nil {
		return entity.User{}, errAddAccount
	}

	return getAccount, nil
}
