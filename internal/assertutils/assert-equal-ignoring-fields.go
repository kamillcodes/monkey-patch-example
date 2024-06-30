package assertutils

import (
	"reflect"
	"testing"
)

func AssertEqualIgnoringFields(t *testing.T, expected, actual interface{}, ignoreFields ...string) {
	expVal := reflect.ValueOf(expected)
	actVal := reflect.ValueOf(actual)

	for i := 0; i < expVal.Type().NumField(); i++ {
		field := expVal.Type().Field(i).Name
		if contains(ignoreFields, field) {
			continue
		}

		if !reflect.DeepEqual(expVal.Field(i).Interface(), actVal.Field(i).Interface()) {
			t.Errorf("Field %s does not match. Expected %v, got %v", field, expVal.Field(i), actVal.Field(i))
		}
	}
}