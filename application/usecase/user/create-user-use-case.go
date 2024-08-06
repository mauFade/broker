package user

import (
	"fmt"

	u "github.com/mauFade/broker/application/entities/user"
	"github.com/mauFade/broker/application/repository/user"
)

type createUserUseCase struct {
	userRepository *user.UserRepository
}

type CreateUserInput struct {
	Name       string
	Email      string
	Phone      string
	Password   string
	Profession string
}

func NewCreateUserUseCase(r *user.UserRepository) *createUserUseCase {
	return &createUserUseCase{
		userRepository: r,
	}
}

func (s *createUserUseCase) Execute(data *CreateUserInput) {
	s.verifyEmail(data.Email)

	s.userRepository.Create(&u.User{
		ID: "1",
	})
}

func (s *createUserUseCase) verifyEmail(email string) {
	userExist := s.userRepository.FindByEmail(email)

	fmt.Println(userExist)

	if userExist != nil {
		fmt.Println("EXISTE CARAI")
	} else {
		fmt.Println("NÃO EXISTE CARAI")
	}

}
