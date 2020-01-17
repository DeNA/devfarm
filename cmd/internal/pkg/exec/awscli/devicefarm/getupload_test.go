package devicefarm

import "testing"

func TestNewUploadGetter(t *testing.T) {
	execute := StubExecutor([]byte(getUploadJSONExample), []byte{}, nil)
	getUpload := NewUploadGetter(execute)

	got, err := getUpload("arn:aws:devicefarm:ANY_UPLOAD_ARN")

	if err != nil {
		t.Errorf("got (_, %v), want (_, nil)", err)
		return
	}

	if got != getUploadExample {
		t.Errorf("got (%v, nil), want (%v, nil)", got, getUploadExample)
		return
	}
}
