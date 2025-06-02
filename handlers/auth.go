package handlers

import (
	"gofiber-auth/database"
	"gofiber-auth/models"
	"gofiber-auth/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Register(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	input.Password = string(hashedPassword)

	if err := database.DB.Create(&input).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "User already exists"})
	}

	return c.JSON(fiber.Map{"message": "Registered successfully"})
}

func Login(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "User not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Incorrect password"})
	}

	token, _ := utils.GenerateJWT(user.Email, time.Hour*24)
	return c.JSON(fiber.Map{"token": token})
}

func ForgotPassword(c *fiber.Ctx) error {
	var input struct {
		Email string `json:"email"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid email"})
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	token, _ := utils.GenerateJWT(user.Email, time.Minute*30)
	user.ResetToken = token
	database.DB.Save(&user)

	utils.SendResetEmail(user.Email, token)
	return c.JSON(fiber.Map{"message": "Reset email sent"})
}

func GetMessage(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Test is Completed",
	})
}

func ResetPassword(c *fiber.Ctx) error {
	var input struct {
		Token       string `json:"token"`
		NewPassword string `json:"newPassword"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	email, err := utils.ParseJWT(input.Token)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid token"})
	}

	var user models.User
	database.DB.Where("email = ?", email).First(&user)

	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.NewPassword), 14)
	user.Password = string(hashed)
	user.ResetToken = ""
	database.DB.Save(&user)

	return c.JSON(fiber.Map{"message": "Password updated"})
}
