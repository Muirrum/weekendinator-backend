package auth

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"

	"github.com/Muirrum/weekendinator-backend/db"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SetupRoutes(r *gin.RouterGroup) {
	r.POST("/login", loginUser)
	r.POST("/register", registerUser)
}

func loginUser(c *gin.Context) {
	c.JSON(200, gin.H{"status": "unimplemented"})
}

func registerUser(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	username := c.PostForm("username")

	password_hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := db.User{
		Username:     username,
		PasswordHash: string(password_hash),
		Email:        email,
		FirstName:    firstName,
		LastName:     lastName,
	}
	result := db.DB.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"status": "error", "message": result.Error})
	}

	rawToken, err := generateToken()
	if err != nil {
		c.String(500, err.Error())
		return
	}
	token := &db.Token{
		UserID:    user.ID,
		Token:     rawToken,
		ExpiresAt: time.Now().Add(2592000 * time.Second),
	}
	result = db.DB.Create(&token)
	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"status": "error", "message": result.Error})
		return
	}

	c.SetCookie("token", token.Token, 2592000, "/", "", true, true)

	go sweepOldTokens()
}

func generateToken() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}

func sweepOldTokens() {
	log.Println("Cleaning old tokens...")
	result := db.DB.Unscoped().Delete(db.Token{}, "expires_at <= ?", time.Now())
	if result.Error != nil {
		log.Println(result.Error)
	}
}
