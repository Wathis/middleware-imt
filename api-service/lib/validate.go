package lib

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

func ValidateQueryString(r *http.Request, rules govalidator.MapData, messages govalidator.MapData) map[string]interface{} {
	opts := govalidator.Options{
		Request:         r,
		Rules:           rules,
		Messages:        messages,
		RequiredDefault: true,
	}
	v := govalidator.New(opts)
	validations := v.Validate()
	if len(validations) > 0 {
		return map[string]interface{}{"validationError": validations}
	}
	return nil
}
