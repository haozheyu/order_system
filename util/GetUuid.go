package util

import (
	"fmt"
	"math/rand"

	"github.com/gofrs/uuid"
)

func GetUUID1() (string, error) {
	u2, err := uuid.NewV4()
	if err != nil {
		//fmt.Printf("Something went wrong: %s", err)
		return "", err
	}
	return u2.String(), nil
}

func GetUUID() (string, error)  {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
	b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return uuid,nil
}