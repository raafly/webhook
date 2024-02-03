package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type PasswordResetToken struct {
	UserID      string
	Token       string
	Expiration  time.Time
}

var tokenStore map[string]PasswordResetToken

func init() {
	tokenStore = make(map[string]PasswordResetToken)
}

// Fungsi ini akan menghasilkan token untuk reset password dan menyimpannya di server
func generateAndStoreResetPasswordToken(userID string) string {
	token := generateRandomToken(16) // Ganti dengan fungsi pembuat token sesuai kebutuhan
	expirationTime := time.Now().Add(1 * time.minute) // Ganti dengan waktu kadaluarsa yang diinginkan

	// Simpan token bersamaan dengan informasi pengguna di server
	tokenStore[token] = PasswordResetToken{
		UserID:     userID,
		Token:      token,
		Expiration: expirationTime,
	}

	return token
}

// Fungsi ini akan memeriksa validitas token reset password saat pengguna mengakses tautan
func validateResetPasswordToken(token string) (bool, string) {
	resetToken, exists := tokenStore[token]
	if !exists || time.Now().After(resetToken.Expiration) {
		// Token tidak valid atau sudah kadaluarsa
		return false, ""
	}

	return true, resetToken.UserID
}

// Fungsi ini akan menghasilkan token acak (random string)
func generateRandomToken(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	token := make([]byte, length)
	for i := range token {
		token[i] = charset[randomInt(len(charset))]
	}
	return string(token)
}

// Fungsi ini akan menghasilkan nilai acak dalam rentang [0, max)
func randomInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func main() {
	app := fiber.New()

	// Endpoint untuk mengirim link reset password dengan token acak
	app.Get("/send-reset-link", func(c *fiber.Ctx) error {
		// Contoh: Mendapatkan UserID dari pengguna yang meminta reset password
		userID := c.Query("user_id")
		if userID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "User ID is required.",
			})
		}

		// Contoh: Membuat dan menyimpan token reset password
		token := generateAndStoreResetPasswordToken(userID)

		// Mengirim link reset password beserta token dalam response
		resetLink := fmt.Sprintf("/reset-password?token=%s", token)
		return c.JSON(fiber.Map{
			"reset_link": resetLink,
		})
	})

	// Endpoint untuk reset password
	app.Get("/reset-password", func(c *fiber.Ctx) error {
		// Mengambil token dari query parameters
		token := c.Query("token")

		// Mengecek apakah token valid
		isValid, userID := validateResetPasswordToken(token)
		if !isValid {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid or expired token.",
			})
		}

		// Jika token valid, dapatkan informasi pengguna dari token
		return c.JSON(fiber.Map{
			"message":   "Token is valid.",
			"user_id":   userID,
			"token":     token,
			"valid_for": tokenStore[token].Expiration.Sub(time.Now()).String(),
		})
	})

	// Jalankan server di port 3000
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
