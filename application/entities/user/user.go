package user

import "time"

type Address struct {
	ID     string `gorm:"type:uuid"`
	UserID string `gorm:"type:uuid"`
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
