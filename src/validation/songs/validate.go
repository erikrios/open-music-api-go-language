package songs

import (
	"github.com/erikrios/open-music-api-go-language/src/api/songs/payloads"
	"github.com/erikrios/open-music-api-go-language/src/validation"
	"github.com/go-playground/validator/v10"
)

func Validate(payload payloads.Payload) []*validation.Error {
	var errors []*validation.Error
	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element validation.Error
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
