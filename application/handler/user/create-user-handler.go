package user

import (
	"github.com/gofiber/fiber/v2"
	ur "github.com/mauFade/broker/application/repository/user"
	"github.com/mauFade/broker/application/usecase/user"
	"github.com/mauFade/broker/database"
)

func CreateUserHandler(c *fiber.Ctx) error {
	payload := struct {
		Name       string `json:"name"`
		Email      string `json:"email"`
		Phone      string `json:"phone"`
		Password   string `json:"password"`
		Profession string `json:"profession"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	uc := user.NewCreateUserUseCase(ur.NewUserRepository(database.Database.DataBase))

	uc.Execute(&user.CreateUserInput{
		Name:       payload.Name,
		Email:      payload.Email,
		Phone:      payload.Phone,
		Password:   payload.Password,
		Profession: payload.Profession,
	})

	return c.JSON(payload)
}
