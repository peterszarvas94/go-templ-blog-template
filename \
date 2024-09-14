package errors

import "fmt"

type ProtectedRouteError struct {
	Filename string
	Route    string
}

func (e *ProtectedRouteError) Error() string {
	return fmt.Sprintf(
		"You can not have '%s' as filename because '%s' is a custom route",
		e.Filename,
		e.Route,
	)
}
