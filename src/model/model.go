package model

type Shortener struct {
	ID          uint   `json:"id" gorm:"primary_key; autoIncrement"`
	Slug        string `json:"slug" binding:"required" gorm:"type:varchar(255)"`
	OriginalUrl string `json:"original_url" binding:"required" gorm:"type:varchar(2048)"`
}
