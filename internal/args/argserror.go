package args

import "fmt"

type ArgsError struct {
	msg string
}

func (ae *ArgsError) Error() string {
	return fmt.Sprintf("Arguments error: %s", ae.msg)
}
