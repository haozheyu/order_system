package util

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
)


func TestCreateToken(t *testing.T) {
	token, _ := CreateToken([]byte(MySigningKey), "YDQ", "22222222", true)
	t.Log(token)

	claims, err := ParseToken(token, []byte(MySigningKey))
	if nil != err {
		t.Log(" err :", err)
	}
	t.Log("claims:", claims)
	t.Log("claims phone:", claims.(jwt.MapClaims)["phone"]) //接口类型 需要反射获取
}

func BenchmarkCreateToken(b *testing.B) {
	token, _ := CreateToken([]byte(MySigningKey), "YDQ", "22222222", true)
	b.Log(token)

	claims, err := ParseToken(token, []byte(MySigningKey))
	if nil != err {
		b.Log(" err :", err)
	}
	b.Log("claims:", claims)
	b.Log("claims phone:", claims.(jwt.MapClaims)["phone"])
}

