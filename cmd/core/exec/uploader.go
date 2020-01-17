package exec

import (
	"fmt"
	"github.com/dena/devfarm/cmd/core/logging"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type Uploader func(url string, modifyRequest func(*http.Request), reader io.Reader) error

func NewUploader(logger logging.SeverityLogger, dryRun bool) Uploader {
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	if dryRun {
		return func(url string, modifyRequest func(*http.Request), reader io.Reader) error {
			return dryUpload(logger, url, modifyRequest, reader)
		}
	}

	return func(url string, modifyRequest func(*http.Request), reader io.Reader) error {
		return upload(logger, &client, url, modifyRequest, reader)
	}
}

func dryUpload(logger logging.SeverityLogger, url string, modifyRequest func(r *http.Request), reader io.Reader) error {
	request, reqErr := http.NewRequest("", url, reader)
	if reqErr != nil {
		return reqErr
	}
	modifyRequest(request)

	bytes, err := ioutil.ReadAll(reader)

	if err != nil {
		logger.Error(fmt.Sprintf("upload: %s (read error: %s)", url, err.Error()))
		return err
	}

	logger.Debug(fmt.Sprintf("upload: %s (%d bytes)", url, len(bytes)))
	return nil
}

func upload(logger logging.SeverityLogger, client *http.Client, url string, modifyRequest func(r *http.Request), reader io.Reader) error {
	// NOTE: Out to temp file to be able to retry using console by users.
	backupFile, backupFilepath, backupErr := createUploadBackupFile()
	if backupErr != nil {
		return backupErr
	}
	defer backupFile.Close()
	teeReader := io.TeeReader(reader, backupFile)

	request, reqErr := http.NewRequest("", url, teeReader)
	if reqErr != nil {
		return reqErr
	}
	modifyRequest(request)

	response, err := client.Do(request)
	if err != nil {
		// NOTE: Dumps unsent payload.
		devNull, ok := ioutil.Discard.(io.ReaderFrom)
		if ok {
			_, _ = devNull.ReadFrom(teeReader)
		} else {
			_, _ = ioutil.ReadAll(teeReader)
		}

		logger.Debug(fmt.Sprintf("upload network error: %s", err.Error()))
		logger.Debug(fmt.Sprintf("you can retry: $ curl -T '%s' '%s'", backupFilepath, url))
		return err
	}
	defer response.Body.Close()

	isSuccess := response.StatusCode > 100 && response.StatusCode < 400
	if !isSuccess {
		// NOTE: Dumps unsent payload.
		devNull, ok := ioutil.Discard.(io.ReaderFrom)
		if ok {
			_, _ = devNull.ReadFrom(teeReader)
		} else {
			_, _ = ioutil.ReadAll(teeReader)
		}

		logger.Debug(fmt.Sprintf("upload error: %s", response.Status))
		logger.Debug(fmt.Sprintf("you can retry: $ curl -T '%s' '%s'", backupFilepath, url))
		body, _ := ioutil.ReadAll(response.Body)
		return fmt.Errorf("unexpected HTTP status: %s\n%s", response.Status, string(body))
	}

	logger.Debug(fmt.Sprintf("upload complete: %s", response.Status))
	return nil
}

func createUploadBackupFile() (io.WriteCloser, string, error) {
	tempDir, tempDirErr := ioutil.TempDir(os.TempDir(), "devfarm-uploaded")
	if tempDirErr != nil {
		return nil, "", tempDirErr
	}
	mirrorFilepath := filepath.Join(tempDir, "payload")
	mirrorFile, openErr := os.OpenFile(mirrorFilepath, os.O_RDWR|os.O_CREATE, 0644)
	if openErr != nil {
		return nil, "", openErr
	}

	return mirrorFile, mirrorFilepath, nil
}
