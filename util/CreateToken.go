package util

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)
var (
	MySigningKey = []byte("$sdf11fsdaf23t6342")
	)
type jwtCustomClaims  struct {
	Phone string `json:"phone"`
	Admin bool `json:"admin"`
	jwt.StandardClaims
}
func CreateToken(SecreKey []byte, issuer string, Phone string,isAdmin bool)(tokenString  string,err error){
	claims := &jwtCustomClaims{
		Phone,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()), //过期时间
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(SecreKey)
	return
}

func ParseToken(tokenSrt string, SecretKey []byte) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	claims = token.Claims

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println(claims["phone"], claims["nbf"])
	} else {
		log.Println(err)
	}
	return

}




