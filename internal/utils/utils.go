package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"oliva-back/internal/models"
	"regexp"
	"time"
)

func ValidateEmail(mail string) bool {
	regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return regex.MatchString(mail)
}
func ValidatePhoneNumber(phone_number string) bool {
	regex := regexp.MustCompile("^(\\+7|7|8)?[\\s\\-]?\\(?[489][0-9]{2}\\)?[\\s\\-]?[0-9]{3}[\\s\\-]?[0-9]{2}[\\s\\-]?[0-9]{2}$")
	return regex.MatchString(phone_number)
}
func CheckEmpty(data interface{}) bool {
	if data == nil {
		return true
	}

	switch v := data.(type) {
	case string:
		return v == ""
	case int:
		return v == 0
	case float64:
		return v == 0.0
	case bool:
		return !v
	case []interface{}:
		return len(v) == 0
	default:
		return false
	}
}

func NewToken(user models.Users, duration time.Duration, secret []byte) (string, error) {
	claims := jwt.MapClaims{
		"username": user.Surname,
		"mail":     user.Email,
		"exp":      time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}
