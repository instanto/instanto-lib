package instanto_lib_db

import (
	"strings"
)

func HelperValidateOrderBy(orderBy string, allowedColumns []string) (err *ValidationError) {
	orderByColumns := strings.Split(orderBy, ",")
	for _, order := range orderByColumns {
		selector := strings.Split(order, " ")
		if len(selector) != 2 {
			err = &ValidationError{"order", "bad order query"}
			return
		}
		column := selector[0]
		direction := selector[1]
		if direction != "asc" && direction != "desc" {
			err = &ValidationError{"order", "bad order query"}
			return
		}
		founded := HelperIsContained(column, allowedColumns)
		if founded == false {
			err = &ValidationError{"order", "bad order query"}
			return
		}
	}
	return
}
func HelperIsContained(needle string, haystack []string) (found bool) {
	for _, value := range haystack {
		if value == needle {
			found = true
		}
	}
	return
}
