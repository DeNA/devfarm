package exec

import (
	"fmt"
	"github.com/dena/devfarm/cmd/internal/pkg/logging"
	"io/ioutil"
	"net/http"
)

type HTTPSender func(request *http.Request) ([]byte, *http.Response, error)

func NewHTTPSender(logger logging.SeverityLogger, dryRun bool) HTTPSender {
	httpClient := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	if dryRun {
		return func(request *http.Request) ([]byte, *http.Response, error) {
			return drySendHTTP(logger, request)
		}
	}

	return func(request *http.Request) ([]byte, *http.Response, error) {
		return sendHTTP(logger, &httpClient, request)
	}
}

func drySendHTTP(logger logging.SeverityLogger, request *http.Request) ([]byte, *http.Response, error) {
	bodyBytes, readErr := ioutil.ReadAll(request.Body)
	if readErr != nil {
		return nil, nil, readErr
	}

	if closeErr := request.Body.Close(); closeErr != nil {
		return nil, nil, closeErr
	}

	logger.Debug(fmt.Sprintf("http (dry run): send the request to %s\n%s", request.URL, string(bodyBytes)))
	responseBody := make([]byte, 0)

	return responseBody, &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body:       nil,
	}, nil
}

// FIXME: Log response using TeeReader.
func sendHTTP(logger logging.SeverityLogger, httpClient *http.Client, request *http.Request) ([]byte, *http.Response, error) {
	logger.Debug(fmt.Sprintf("http: send the request to %s", request.URL))

	// XXX: Prevent to use keep-alive, because it might get io.EOF...
	request.Close = true

	res, err := httpClient.Do(request)
	if err != nil {
		logger.Debug(fmt.Sprintf("http: failed to send: %s returned:\n%s", request.URL, err.Error()))
		return nil, nil, err
	}
	defer res.Body.Close()

	bodyBytes, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		logger.Debug(fmt.Sprintf("http: cannot read response.Body: %s", bodyErr.Error()))
		return nil, nil, bodyErr
	}
	logger.Debug(fmt.Sprintf("http: received:\n%s", string(bodyBytes)))

	// XXX: Force to use the []byte that is a first value returned.
	res.Body = nil

	return bodyBytes, res, nil
}
