package awsdevicefarm

import (
	"github.com/dena/devfarm/internal/pkg/logging"
	"github.com/dena/devfarm/internal/pkg/platforms"
	"testing"
)

func TestNewTestPackageUploader(t *testing.T) {
	spyLogger := logging.SpySeverityLogger()
	reserveAndUploaderCallArgs, spyReserveAndUploader := spyReserveUploaderIfNotExists(anySuccessfulReserveAndUploaderIfNotExists())

	uploadTestPackage := newTestPackageUploader(spyLogger, newTestPackageGen(), platforms.NewCRC32Hasher(), spyReserveAndUploader)

	_, err := uploadTestPackage("arn:aws:devicefarm:ANY_PROJECT")
	if err != nil {
		t.Errorf("got %v, want nil", err)
		t.Log(spyLogger.Logs.String())
		return
	}

	if len(*reserveAndUploaderCallArgs) != 1 {
		t.Errorf("number of reserveAndUpload calls are %v, want 1", len(*reserveAndUploaderCallArgs))
		t.Log(spyLogger.Logs.String())
		return
	}

	size := (*reserveAndUploaderCallArgs)[0].size
	if size < 1 {
		t.Errorf("got %d, want > 0", size)
		t.Log(spyLogger.Logs.String())
		return
	}
}
