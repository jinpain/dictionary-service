package insurance

type Model struct {
	ID       string `gorm:"primaryKey" json:"id"`
	Type     int    `gorm:"not null" json:"type"`
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	FullName string `gorm:"type:varchar(255);not null" json:"full_name"`
	Address  string `gorm:"type:text" json:"address"`
	Phone    string `gorm:"type:char(11)" json:"phone"`
	Email    string `gorm:"type:varchar(32)" json:"email"`
	Website  string `gorm:"type:varchar(255)" json:"website"`
	OGRN     string `gorm:"type:char(13)" json:"ogrn"`
	OKPO     string `gorm:"type:varchar(10)" json:"okpo"`
	OKATO    string `gorm:"type:varchar(11)" json:"okato"`
}

func (Model) TableName() string {
	return "insurance"
}
