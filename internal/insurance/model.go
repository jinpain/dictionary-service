package insurance

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	Type     int    `gorm:"not null"`
	Name     string `gorm:"type:varchar(255);not null"`
	FullName string `gorm:"type:varchar(255);not null"`
	Address  string `gorm:"type:text"`
	Phone    string `gorm:"type:char(11)"`
	Email    string `gorm:"type:varchar(32)"`
	Website  string `gorm:"type:varchar(255)"`
	OGRN     string `gorm:"type:char(13)"`
	OKPO     string `gorm:"type:varchar(10)"`
	OKATO    string `gorm:"type:varchar(11)"`
}

func (Model) TableName() string {
	return "insurance"
}
