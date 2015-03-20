package instantolib

import "fmt"

func validateDegree(field, value string) *ValidationError {
	if value != "none" && value != "dr" && value != "dra" {
		return &ValidationError{field, "value must be none, dr or dra"}
	}
	return nil
}

func validateScope(field, value string) *ValidationError {
	if value != "regional" && value != "national" && value != "international" {
		return &ValidationError{field, "value must be regional, national or international"}
	}
	return nil
}
func validateNotEmpty(field, value string) *ValidationError {
	if len(field) == 0 {
		return &ValidationError{field, "cannot be empty"}
	}
	return nil
}
func validateIsNumber(field string, value int64) *ValidationError {
	if value < 0 {
		return &ValidationError{field, fmt.Sprintf("must be greater than %d", value)}
	}
	return nil
}
func validateLength(field, value string, length int) *ValidationError {
	if len(value) > length {
		return &ValidationError{field, fmt.Sprintf("length cannot be greater than %d", length)}
	}
	return nil
}
