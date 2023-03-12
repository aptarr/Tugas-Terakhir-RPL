package entity

type Comment struct {
	ID      uint64 `gorm:"primaryKey" json:"id"`
	Comment string `json:"comment" binding:"required"`
	BlogID  uint64 `gorm:"foreignKey" json:"blog_id" binding:"required"`
	Blog    *Blog  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"transaksi1,omitempty"`
	UserID  uint64 `gorm:"foreignKey" json:"user_id" binding:"required"`
	User    *User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"transaksi2,omitempty"`
}
