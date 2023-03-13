package jwt

import (
	"errors"
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const SECRET_KEY = "someSecret_addThisToConf"

func GenerateToken(email string, role string) (string, error) {
	tokenLifespan := 1 // in hrs

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(SECRET_KEY))
}

func TokenValid(ctx *gin.Context, role string) error {
	tokenString := extractToken(ctx)
	token, err := parseToken(tokenString)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid && claims["role"] == role {
		return nil
	}

	return errors.New("token or role not valid")
}

func ExtractTokenEmail(ctx *gin.Context) (string, error) {
	tokenString := extractToken(ctx)
	token, err := parseToken(tokenString)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return fmt.Sprintf("%s", claims["email"]), nil
	}
	return "", nil
}

func extractToken(ctx *gin.Context) string {
	/*token := ctx.Query("token")
	if token != "" {
		return token
	}*/
	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token not signed properly")
		}
		return []byte(SECRET_KEY), nil
	})
	return token, err
}
