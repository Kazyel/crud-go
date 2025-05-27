package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Must be at least %s characters", fe.Param())
	case "max":
		return fmt.Sprintf("Must be at most %s characters", fe.Param())
	case "alphanum":
		return "Should contain only alphanumeric characters"
	}
	return fe.Error()
}

func HandleBindingError(c *gin.Context, err error) {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))

		for i, fe := range ve {
			out[i] = ApiError{fe.Field(), getErrorMessage(fe)}
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
	} else if errors.Is(err, io.EOF) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Request body is empty"})
	} else if errors.Is(err, &json.UnmarshalTypeError{}) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid type provided"})
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
