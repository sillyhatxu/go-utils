package header

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetDeviceId(context *gin.Context) string {
	values := context.Request.Header["Device-Id"]
	if len(values) > 0 {
		return values[0]
	}
	return ""
}

func GetCurrentUserId(context *gin.Context, secretKey string) (string, error) {
	cookie, err := context.Request.Cookie("X-ADV-TOKEN")
	if err != nil {
		return "", err
	}
	return ParseToken(cookie.Value, secretKey)
}

func ParseToken(tokenSrc, secretKey string) (string, error) {
	standardClaims := jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenSrc, &standardClaims, func(token *jwt.Token) (interface{}, error) {
		dec, err := base64.URLEncoding.DecodeString(secretKey)
		if err != nil {
			return nil, err
		}
		return dec, nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("Token valid error.")
	}
	var user User
	if err := json.Unmarshal([]byte(standardClaims.Subject), &user); err != nil {
		return "", err
	}
	return user.Id, nil
}

type User struct {
	Id           string `json:"id"`
	MobileNumber string `json:"mobileNumber"`
	NkpUserId    string `json:"nkpUserId"`
	Channel      string `json:"channel"`
}
