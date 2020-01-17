package testutil

var AnyError error = &anyError{}

type anyError struct{}

func (anyError) Error() string {
	return "ANY_ERROR_FOR_TESTING"
}
