package cli

import "fmt"

func NewErrorAndUsage(message string, usage string) *ErrorAndUsage {
	return &ErrorAndUsage{
		message: message,
		usage:   usage,
	}
}

type ErrorAndUsage struct {
	message string
	usage   string
}

func (u *ErrorAndUsage) SetMessage(message string) {
	u.message = message
}

func (u *ErrorAndUsage) Error() string {
	if len(u.message) > 0 {
		return fmt.Sprintf("%s\n%s", u.message, u.usage)
	}
	return u.usage
}
