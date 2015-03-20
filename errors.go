package instantolib

type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func (err *ValidationError) Error() string {
	return err.Field + ": " + err.Reason
}
