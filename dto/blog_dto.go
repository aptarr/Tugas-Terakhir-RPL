package dto

type UserAddBlog struct {
	ID      uint64 `gorm:"primaryKey" json:"id"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Like    uint64 `json:"like" binding:"required"`
	UserID  uint64 `gorm:"foreignKey" json:"user_id" binding:"required"`
}

type UserUpdateBlog struct {
	ID      uint64 `gorm:"primaryKey" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Like    uint64 `json:"like"`
	UserID  uint64 `gorm:"foreignKey" json:"user_id"`
}
