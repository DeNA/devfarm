package devicefarm

import "testing"

func TestNewUploadDeleter(t *testing.T) {
	execute := AnySuccessfulExecutor()
	deleteUpload := NewUploadDeleter(execute)

	err := deleteUpload("arn:aws:devicefarm:ANY_UPLOAD")

	if err != nil {
		t.Errorf("got %v, want nil", err)
		return
	}
}
