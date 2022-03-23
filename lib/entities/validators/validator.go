package validators

import (
	"reflect"
	"regexp"
)

type Validator struct{}

func (validator *Validator) IsEmail(param string) bool {
	regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return regex.MatchString(param)
}

func (validator *Validator) IsString(param string) bool {
	condition := reflect.TypeOf(param)
	return condition.String() == "string"
}

func (validator *Validator) IsDecimal(param float64) bool {
	condition := reflect.TypeOf(param)
	return condition.String() == "float64"
}

func (validator *Validator) IsInt(param int64) bool {
	condition := reflect.TypeOf(param)
	return condition.String() == "int64"
}
