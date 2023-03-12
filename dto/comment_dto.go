package dto

type UserAddComment struct {
	ID      uint64 `gorm:"primaryKey" json:"id"`
	Comment string `json:"comment" binding:"required"`
	BlogID  uint64 `gorm:"foreignKey" json:"blog_id" binding:"required"`
	UserID  uint64 `gorm:"foreignKey" json:"user_id" binding:"required"`
}

type UserUpdateComment struct {
	ID      uint64 `gorm:"primaryKey" json:"id"`
	Comment string `json:"comment"`
}
