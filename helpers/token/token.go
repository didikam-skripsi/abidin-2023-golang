package token

import (
	"fmt"
	"gostarter-backend/models"
	"gostarter-backend/response"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GenerateToken(user models.User) (string, error) {

	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user"] = response.UserResponse{}.ResponseWithAccess(&user)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func TokenValid(c *fiber.Ctx) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *fiber.Ctx) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenUser(c *fiber.Ctx) (models.User, error) {
	user := models.User{}
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return user, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		authUser := claims["user"].(map[string]interface{})
		if uuidStr, ok := authUser["uuid"].(string); ok {
			userUUID, err := uuid.Parse(uuidStr)
			if err != nil {
				return user, err
			}
			user.UUID = userUUID
		}
		if username, ok := authUser["username"].(string); ok {
			user.Username = username
		}
		if role, ok := authUser["role"].(string); ok {
			user.Role = models.RoleType(role)
		}
	}
	return user, nil
}
