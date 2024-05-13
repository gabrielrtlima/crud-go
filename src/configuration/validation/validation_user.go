package validation

import (
	"errors"

	resterr "github.com/gabrielrtlima/crud-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/goccy/go-json"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {

		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en_US")
		en_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(
	validation_err error,
) *resterr.RestErr {
	var jsonError *json.UnmarshalTypeError
	var validationErrors validator.ValidationErrors

	if errors.As(validation_err, &jsonError) {
		return resterr.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &validationErrors) {
		causes := make([]resterr.Causes, 0)
		for _, err := range validationErrors {
			causes = append(causes, resterr.Causes{
				Field:   err.Field(),
				Message: err.Translate(transl),
			})
		}
		return resterr.NewBadRequestValidationError("Invalid field", causes)
	}
	return resterr.NewBadRequestError("Error trying to validate fields")
}
