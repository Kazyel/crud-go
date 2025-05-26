package tests

import (
	"fmt"
	"rest-crud-go/internal/core/utils"
	"testing"
)

func TestVGenerateJWT(t *testing.T) {
	id, err := utils.GenerateJWT("kazyel")

	if err != nil {
		t.Errorf("error generating JWT: %v", err)
	}

	fmt.Println(id)
}

func TestVParsingJWT(t *testing.T) {
	token, err := utils.GenerateJWT("kazyel")

	fmt.Println(token)

	if err != nil {
		t.Errorf("error generating JWT: %v", err)
	}

	id, err := utils.ParseJWT(token)

	if err != nil {
		t.Errorf("error parsing JWT: %v", err)
	}

	fmt.Println(id)
}
