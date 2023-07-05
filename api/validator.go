package api

import (
	"github.com/NopparootSuree/Learning-GO/utils"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.IsSupportCurrency(currency)
	}
	return false
}
