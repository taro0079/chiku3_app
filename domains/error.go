package domains

import "time"

type UserNameError struct {
	When time.Time
	What string
}

func (e UserNameError) Error() string {
	return "error"
}
