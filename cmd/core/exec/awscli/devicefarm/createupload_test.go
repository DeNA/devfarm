package devicefarm

import "testing"

func TestNewUploadCreator(t *testing.T) {
	execute := StubExecutor([]byte(createUploadJSONExample), []byte{}, nil)
	createUpload := NewUploadCreator(execute)

	got, err := createUpload(
		"arn:aws:devicefarm:ANY_UPLOAD",
		"dummy.ipa",
		UploadTypeIsIOSApp,
	)

	if err != nil {
		t.Errorf("got (_, %v), want (_, nil)", err)
		return
	}

	if got != initializedUploadExample {
		t.Errorf("got (%v, nil), want (%v, nil)", got, initializedUploadExample)
		return
	}
}
