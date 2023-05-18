package auth

import (
	"nom/models"
	"nom/structs"

	"github.com/gofiber/fiber/v2"
)

var authRepo AuthRepository

func UserLogin(c *fiber.Ctx) error {
	var creds structs.Credentials

	var user models.User

	if err := c.BodyParser(&creds); err != nil {
		return c.Status(400).JSON(fiber.Map{})
	}
	return c.SendString("This is auth routes")

	// retrive user from the database
	err := authRepo.FindUserByEmail(creds.Username, &user)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "")
	}
	// check user data validity
	valid, err := user.PasswordMatches(user.Password)
	if err != nil || !valid {
		return c.Status(404).JSON(fiber.Map{"message": "invalid credentials"})
	}
	// generate the tokens and return
	u := structs.JwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  creds.Username,
		Email:     user.Email,
	}

	token, err := authRepo.GenerateTokenpar(&u)

	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSONP(token)

}
