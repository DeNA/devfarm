package executor

import (
	"net/http"
)

func AnySuccessfulHTTPSender() HTTPSender {
	responseBody := make([]byte, 0)

	res := http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body:       nil,
	}

	return StubHTTPSender(responseBody, &res, nil)
}

func StubHTTPSender(bodyBytes []byte, response *http.Response, err error) HTTPSender {
	return func(*http.Request) ([]byte, *http.Response, error) {
		return bodyBytes, response, err
	}
}
