package awsdevicefarm

import (
	"github.com/dena/devfarm/cmd/internal/pkg/exec"
	"github.com/dena/devfarm/cmd/internal/pkg/exec/awscli/devicefarm"
	"github.com/dena/devfarm/cmd/internal/pkg/logging"
	"github.com/dena/devfarm/cmd/internal/pkg/platforms"
	"testing"
)

func TestNewAppUploader(t *testing.T) {
	spyLogger := logging.SpySeverityLogger()
	uploaderCallArgs, spyUploader := spyReserveUploaderIfNotExists(anySuccessfulReserveAndUploaderIfNotExists())

	uploadApp := newAppUploader(
		spyLogger,
		exec.FakeFileOpener([]byte("hello world")),
		platforms.NewCRC32Hasher(),
		spyUploader,
	)

	_, err := uploadApp("path/to/example-app.ipa", platforms.OSIsIOS, "arn:aws:devicefarm:ANY_PROJECT")
	if err != nil {
		t.Errorf("got (_, %v), want (_, nil)", err)
		t.Log(spyLogger.Logs.String())
		return
	}

	if len(*uploaderCallArgs) != 1 {
		t.Errorf("number of reserveUpload calls are %d, want 1", len(*uploaderCallArgs))
		t.Log(spyLogger.Logs.String())
		return
	}

	capturedUploadFileName := (*uploaderCallArgs)[0].uploadProp.fileName

	var expectedUploadFileName devicefarm.UploadFileName = "devfarm-example-app-0d4a1185.ipa"
	if capturedUploadFileName != expectedUploadFileName {
		t.Errorf("got %v, want %v", capturedUploadFileName, expectedUploadFileName)
		t.Log(spyLogger.Logs.String())
		return
	}
}
