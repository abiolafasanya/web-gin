package models

type Book struct {
	
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Title string `gorm:"size:256" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	UserID	uint `gorm:"not null" json:"user_id"`
	User User `gorm:"foreignKey;UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}