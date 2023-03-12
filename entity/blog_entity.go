package entity

type Blog struct {
	ID      uint64    `gorm:"primaryKey" json:"id"`
	Title   string    `json:"title" binding:"required"`
	Content string    `json:"content" binding:"required"`
	Like    uint64    `json:"like" binding:"required"`
	UserID  uint64    `gorm:"foreignKey" json:"user_id" binding:"required"`
	User    *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"transaksi,omitempty"`
	Comment []Comment `json:"comments,omitempty"`
}
