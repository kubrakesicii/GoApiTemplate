package entity

type Test struct {
	Id       uint   `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Email    string `gorm:"not null"`
	Password []byte
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
}
