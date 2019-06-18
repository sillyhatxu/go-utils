package header

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

func GetDeviceId(context *gin.Context) string {
	values := context.Request.Header["Device-Id"]
	if len(values) > 0 {
		return values[0]
	}
	return ""
}

func GetAcceptLanguage(context *gin.Context) language.Tag {
	local := language.English
	acceptLanguage := context.Request.Header["Accept-Language"]
	if acceptLanguage != nil && len(acceptLanguage) > 0 {
		switch acceptLanguage[0] {
		case "zh":
			local = language.Chinese
		case "id":
			local = language.Indonesian
		default:
			//"en-US"
			local = language.English
		}
	}
	return local
}

type ResponseMessage struct {
	Context *gin.Context
	MsgZH   string
	MsgEN   string
	MsgID   string
}

func (rm ResponseMessage) GetResponseMessage() string {
	tag := GetAcceptLanguage(rm.Context)
	switch tag {
	case language.Indonesian:
		return rm.MsgID
	case language.English:
		return rm.MsgEN
	case language.Chinese:
		return rm.MsgZH
	default:
		return rm.MsgEN
	}
}

func GetCurrentUserId(context *gin.Context, secretKey string, enable bool) (string, error) {
	if !enable {
		authorization := context.Request.Header["Authorization"]
		if len(authorization) > 0 {
			return authorization[0], nil
		}
		return "", errors.New("header data error.")
	}
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
