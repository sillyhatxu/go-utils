package header

import (
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSrc(t *testing.T) {
	dec, err := base64.URLEncoding.DecodeString("")
	assert.Nil(t, err)
	enc := base64.URLEncoding.EncodeToString(dec)
	fmt.Println(enc)
}

func TestSrc1(t *testing.T) {
	dec, err := base64.URLEncoding.DecodeString(".")
	assert.Nil(t, err)
	enc := base64.URLEncoding.EncodeToString(dec)
	fmt.Println(enc)
}

func TestDecode(t *testing.T) {
	dec, err := base64.URLEncoding.DecodeString("MWE2NjBkNTQxZTc4MWRiNzE4NDNkZTg1YjcwNWQzN2IK")
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err)
	assert.NotNil(t, dec)
	fmt.Println(dec)

	dec, err = base64.URLEncoding.DecodeString("QURWQU5DRUFJREVWRUxPUEFQUA==")
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err)
	assert.NotNil(t, dec)
	fmt.Println(dec)

	dec, err = base64.URLEncoding.DecodeString("QURWQU5DRUFJREVWRUxPUEJPU1M=")
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err)
	assert.NotNil(t, dec)
	fmt.Println(dec)

	dec, err = base64.URLEncoding.DecodeString("QURWQU5DRUFJU1RBR0lOR0FQUA==")
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err)
	assert.NotNil(t, dec)
	fmt.Println(dec)

	dec, err = base64.URLEncoding.DecodeString("QURWQU5DRUFJU1RBR0lOR0VCT1NT")
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err)
	assert.NotNil(t, dec)
	fmt.Println(dec)
}

func TestParseToken(t *testing.T) {
	secretKey := "MWE2NjBkNTQxZTc4MWRiNzE4NDNkZTg1YjcwNWQzN2IK"
	tokenString := `eyJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJTSE9QSU5UQVIiLCJzdWIiOiJ7XCJpZFwiOlwiVTVDRjg2QzU3QkJGNTBDMDAwMUYxMjEwQ1wiLFwibW9iaWxlTnVtYmVyXCI6XCIrNjI4MTIzMDk5MDQzNFwiLFwibmtwVXNlcklkXCI6XCJVNUNGODZDNTdCQkY1MEMwMDAxRjEyMTBDXCIsXCJjaGFubmVsXCI6XCJTSE9QSU5UQVJcIn0iLCJhdWQiOiJXRUIiLCJpYXQiOjE1NTk4MTYyNjMsImV4cCI6MTU2MDQyMTA2M30.3s0D4bz9kyjc6Zp4tyGt7QK0uHjLo5xwoaYsJBMWvtbwFbeFDPk6SDOOzABiogVMfZzH2cbUZ6z2I6faE0jLaQ`
	userId, err := ParseToken(tokenString, secretKey)
	assert.Nil(t, err)
	assert.NotNil(t, userId)
	assert.EqualValues(t, userId, "U5CF86C57BBF50C0001F1210C")
}

func TestGetCurrent(t *testing.T) {
	tokenString := `eyJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJTSE9QSU5UQVIiLCJzdWIiOiJ7XCJpZFwiOlwiVTVDRjg2QzU3QkJGNTBDMDAwMUYxMjEwQ1wiLFwibW9iaWxlTnVtYmVyXCI6XCIrNjI4MTIzMDk5MDQzNFwiLFwibmtwVXNlcklkXCI6XCJVNUNGODZDNTdCQkY1MEMwMDAxRjEyMTBDXCIsXCJjaGFubmVsXCI6XCJTSE9QSU5UQVJcIn0iLCJhdWQiOiJXRUIiLCJpYXQiOjE1NTk4MTYyNjMsImV4cCI6MTU2MDQyMTA2M30.3s0D4bz9kyjc6Zp4tyGt7QK0uHjLo5xwoaYsJBMWvtbwFbeFDPk6SDOOzABiogVMfZzH2cbUZ6z2I6faE0jLaQ`
	standardClaims := jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &standardClaims, func(token *jwt.Token) (interface{}, error) {
		dec, err := base64.URLEncoding.DecodeString("MWE2NjBkNTQxZTc4MWRiNzE4NDNkZTg1YjcwNWQzN2IK")
		if err != nil {
			return nil, err
		}
		//return []byte("MWE2NjBkNTQxZTc4MWRiNzE4NDNkZTg1YjcwNWQzN2IK"), nil
		return dec, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestGetC(t *testing.T) {
	tokenString := `eyJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJTSE9QSU5UQVIiLCJzdWIiOiJ7XCJpZFwiOlwiVTVDNzg5MzQ1MTlERDI5MDAwMTM5MTA5OVwiLFwibW9iaWxlTnVtYmVyXCI6XCIrNjI4NzA2MDkxMjAxXCIsXCJua3BVc2VySWRcIjpcIlU1Qzc4OTM0NTE5REQyOTAwMDEzOTEwOTlcIn0iLCJhdWQiOiJXRUIiLCJpYXQiOjE1NTk4MTcyNTAsImV4cCI6MTU2MDQyMjA1MH0.TrXnYqUW91fqlQTH54bGrthurlAW0pBd7QK-oYsiQON6J8gnWyqS8SVUoGoKSkZm1DTno3wApLQfC6SNb60ZMg`
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		dec, err := base64.URLEncoding.DecodeString("MWE2NjBkNTQxZTc4MWRiNzE4NDNkZTg1YjcwNWQzN2IK")
		//dec, err := base64.URLEncoding.DecodeString("asdfasdf")
		//dec, err := base64.URLEncoding.DecodeString("asdfasdf")
		if err != nil {
			return nil, err
		}
		//return []byte("MWE2NjBkNTQxZTc4MWRiNzE4NDNkZTg1YjcwNWQzN2IK"), nil
		return dec, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestGetCurr(t *testing.T) {
	//tokenString := "eyJhbGciOiJIUzUxMiJ9.eyJ1c2VyTmFtZSI6InNoaWt1YW4ueHVAYWR2YW5jZS5haSIsImlzcyI6InNob3BpbnRhci1ib3NzLWFwaSIsImlhdCI6MTU1OTgxMzMxNiwiZXhwIjoxNTYyNDA1MzE2fQ.zjsVRkcZbNl6ZYwvGaEikCXB-aW8XyPZ1iwAfXYhC9Q9aFJ6s3r_WfiXpI0eN2-xcSvm4PbV66vffWq7gUeCIA"
	tokenString := "eyJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJua3AtYXBwIiwic3ViIjoie1wiaWRcIjpcIlU1Q0Y2MjZCRDk0OUVFOTAwMDFEMzhFRkFcIixcIm5rcFVzZXJJZFwiOlwiVTVCQkVGN0MxNTU0QjhGNTY1MEUyMDFBMVwiLFwiY2hhbm5lbFwiOlwiU0hPUElOVEFSXCJ9IiwiYXVkIjoiV0VCIiwiaWF0IjoxNTU5ODAxNDYyLCJleHAiOjE1NjA0MDYyNjJ9.Xg4TD3BIcKnL7ylxkZDQ8k75yWfnkdFlYJUoUfjH70bDS-9rZIVr1Md5OgfKQ1aGXnMxvvgZq9yprU3le2uRzg"
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		dec, err := base64.URLEncoding.DecodeString("")
		if err != nil {
			return nil, err
		}
		return dec, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestGetCurrentUserId(t *testing.T) {
	//tokenString := "eyJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJua3AtYXBwIiwic3ViIjoie1wiaWRcIjpcIlU1QjZEMzA5QkZCNzhDRjVEMDY2MDk4RDNcIixcIm5rcFVzZXJJZFwiOlwiVTVCNkQzMDlCRkI3OENGNUQwNjYwOThEM1wiLFwiY2hhbm5lbFwiOlwiU0hPUElOVEFSXCJ9IiwiYXVkIjoiV0VCIiwiaWF0IjoxNTU5ODAzOTMzLCJleHAiOjE1NjA0MDg3MzN9.EGDsaidZMa2wFl4ojpENeZ3ayfppEeslhmFMePo1OgnbtZyCtE8XIGDIEaWetaN-EtpStSf8NWN671FsEps8Pg"
	tokenString := "eyJhbGciOiJIUzUxMiJ9.eyJ1c2VyTmFtZSI6InNoaWt1YW4ueHVAYWR2YW5jZS5haSIsImlzcyI6InNob3BpbnRhci1ib3NzLWFwaSIsImlhdCI6MTU1OTgxMzMxNiwiZXhwIjoxNTYyNDA1MzE2fQ.zjsVRkcZbNl6ZYwvGaEikCXB-aW8XyPZ1iwAfXYhC9Q9aFJ6s3r_WfiXpI0eN2-xcSvm4PbV66vffWq7gUeCIA"
	//tokenString := "eyJhbGciOiJIUzUxMiJ9.eyJ1c2VyTmFtZSI6InNoaWt1YW4ueHVAYWR2YW5jZS5haSIsImlzcyI6InNob3BpbnRhci1ib3NzLWFwaSIsImlhdCI6MTU1OTgwOTM1NywiZXhwIjoxNTYyNDAxMzU3fQ.rJ2YuLbtwzMaYcnAZcmTDHnmCggdeXoRBpWsft6XWjaynwrOzNXRx66zoVyYpWw32pfhGCp9AplkTzITVShyLg"
	//tokenString := "eyJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJua3AtYXBwIiwic3ViIjoie1wiaWRcIjpcIlU1QjZEMzA5QkZCNzhDRjVEMDY2MDk4RDNcIixcIm5rcFVzZXJJZFwiOlwiVTVCNkQzMDlCRkI3OENGNUQwNjYwOThEM1wiLFwiY2hhbm5lbFwiOlwiU0hPUElOVEFSXCJ9IiwiYXVkIjoiV0VCIiwiaWF0IjoxNTU5ODAzOTMzLCJleHAiOjE1NjA0MDg3MzN9.EGDsaidZMa2wFl4ojpENeZ3ayfppEeslhmFMePo1OgnbtZyCtE8XIGDIEaWetaN-EtpStSf8NWN671FsEps8Pg"
	//secretKey := "."
	//claim := jwt.StandardClaims
	//type MyCustomClaims struct {
	//	Foo string `json:"foo"`
	//	jwt.StandardClaims
	//}
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		//token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		//sDec, _ := base64.StdEncoding.DecodeString(".")
		//dec := base64.URLEncoding.EncodeToString([]byte("."))
		dec, err := base64.URLEncoding.DecodeString("...")
		//dec, err := base64.URLEncoding.DecodeString(base64.URLEncoding.EncodeToString([]byte("...")))
		if err != nil {
			return nil, err
		}
		//return base64.URLEncoding.EncodeToString([]byte(".")), nil
		//return base64.URLEncoding.EncodeToString([]byte("")), nil
		//return []byte("..."), nil
		return dec, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err)
	assert.NotNil(t, token)
}
