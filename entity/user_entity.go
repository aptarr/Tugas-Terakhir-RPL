package entity

import (
	"Tugas-Pert4/utils"

	"gorm.io/gorm"
)

type User struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Blog     []Blog `json:"blogs,omitempty"`
}

func (usr *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	usr.Password, err = utils.HasAndSalt(usr.Password)
	if err != nil {
		return err
	}
	return nil
}

func (usr *User) BeforeUpdate(tx *gorm.DB) error {
	var err error
	if usr.Password != "" {
		usr.Password, err = utils.HasAndSalt(usr.Password)
	}
	if err != nil {
		return err
	}
	return nil
}
