package apis

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// func main() {
//     password := "secret"
//     hash, _ := HashPassword(password) // ignore error for the sake of simplicity

//     fmt.Println("Password:", password)
//     fmt.Println("Hash:    ", hash)

//     match := CheckPasswordHash(password, hash)
//     fmt.Println("Match:   ", match)
// }

// Create a JWT token
func CreateJWTToken(userID int64) (string, error) {

	expirationTime := time.Now().Add(time.Hour * 1)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    expirationTime.Unix(),
	})

	secretKey := []byte("my-secret-key")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Verify a JWT token
func VerifyJWTToken(tokenString string) (int64, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}

		secretKey := []byte("my-secret-key")
		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token claims")
	}

	userID := int64(claims["userID"].(int64))
	return userID, nil
}
