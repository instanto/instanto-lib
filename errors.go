package instanto_lib_db

type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func (err *ValidationError) Error() string {
	return err.Field + ": " + err.Reason
}
