package security

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/matejelenc/rest-api/data"
)

//my key used for signing tokens
var secretKey = []byte(os.Getenv("JWT_KEY"))

//claims used when creating a jwt token
type CustomClaims struct {
	UserID string
	jwt.StandardClaims
}

//GenerateJWT creates a jwt token for a logged in user
func GenerateJWT(user *data.Person) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)

	claims := &CustomClaims{
		UserID: strconv.Itoa(int(user.ID)),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

//ValidateToken validates a jwt token
func ValidateToken(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("Token")
	if err != nil {
		return "", err
	}
	claims := CustomClaims{}
	tokenStr := cookie.Value
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("Token not valid")
	}

	return claims.UserID, nil
}
