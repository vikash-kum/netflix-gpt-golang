package controllers

import (
	"net/http"
	"netflix-gpt-backend/config"
	models "netflix-gpt-backend/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var JwtSecret = "your_secret_key_here"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
	var input models.User

	// Bind JSON input to the User model
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Hash the user's password
	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Store the hashed password instead of the plain-text one
	input.Password = hashedPassword

	// Save the user with the hashed password in the database
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	// Send success response
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// LoginUser handles user login and issues a JWT token
func Login(c *gin.Context) {

	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User
	// Query the user by email
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials email"})
		return
	}

	// Compare the stored hashed password with the input password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials password"})
		return
	}

	// Generate JWT token (assuming you have a function for this)
	token, err := GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateToken generates a JWT token for the user

func GenerateToken(user models.User) (string, error) {
	// Define the token expiration time
	expirationTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours

	// Create the claims
	claims := &Claims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "your-app-name", // Set the issuer (optional)
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with your secret key
	signedToken, err := token.SignedString([]byte(JwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
