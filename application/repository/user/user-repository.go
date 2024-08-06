package user

import (
	"github.com/mauFade/broker/application/entities/user"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(d *gorm.DB) *UserRepository {
	return &UserRepository{
		db: d,
	}
}

func (r *UserRepository) Create(user *user.User) {
	r.db.Create(user)
}

func (r *UserRepository) Update(user *user.User) {
	r.db.Save(&user)
}

func (r *UserRepository) FindByEmail(email string) *user.User {
	var user *user.User

	result := r.db.Where("email = ?", email).First(&user)

	if result.RowsAffected == 0 {
		return nil
	}

	return user
}
