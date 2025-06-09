package tests

import (
	"fmt"
	"rest-crud-go/internal/core/utils"
	"testing"
)

func TestVGenerateJWT(t *testing.T) {
	jwtToken, csrfToken, err := utils.GenerateJWT("kazyel")

	if err != nil {
		t.Errorf("error generating JWT: %v", err)
	}

	fmt.Println(jwtToken, csrfToken)
}

func TestVParsingJWT(t *testing.T) {
	jwtToken, csrfToken, err := utils.GenerateJWT("kazyel")

	fmt.Println(jwtToken, csrfToken)

	if err != nil {
		t.Errorf("error generating JWT: %v", err)
	}

	id, err := utils.ParseJWT(jwtToken)

	if err != nil {
		t.Errorf("error parsing JWT: %v", err)
	}

	fmt.Println(id)
}
