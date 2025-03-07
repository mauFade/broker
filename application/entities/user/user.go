package user

import "time"

type Address struct {
	ID            string `gorm:"type:uuid"`
	UserID        string `gorm:"type:uuid"`
	Country       string `gorm:"type:varchar"`
	ZipCode       string `gorm:"type:varchar"`
	Street        string `gorm:"type:varchar"`
	Number        string `gorm:"type:varchar"`
	Neighbourhood string `gorm:"type:varchar"`
	City          string `gorm:"type:varchar"`
	Estate        string `gorm:"type:varchar"`
}

type User struct {
	ID         string  `gorm:"type:uuid"`
	Name       string  `gorm:"type:varchar"`
	Email      string  `gorm:"type:varchar"`
	Phone      string  `gorm:"type:varchar"`
	Password   string  `gorm:"type:varchar"`
	Agency     string  `gorm:"type:varchar"`
	Bank       string  `gorm:"type:varchar"`
	Serial     string  `gorm:"type:varchar"`
	Profession string  `gorm:"type:varchar"`
	Address    Address `gorm:"foreignKey:UserID"`

	Deleted   bool       `gorm:"type:bool"`
	DeletedAt *time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time  `gorm:"type:timestamp"`
	CreatedAt time.Time  `gorm:"type:timestamp"`
}
