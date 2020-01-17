package devicefarm

import (
	"encoding/json"
	"fmt"
)

type UploadARN string

func (a UploadARN) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(a))
}

func (a *UploadARN) UnmarshalJSON(bytes []byte) error {
	var s string

	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}

	*a = UploadARN(s)
	return nil
}

type UploadStatus string

func (u UploadStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(u))
}

func (u *UploadStatus) UnmarshalJSON(bytes []byte) error {
	var s string

	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}

	switch s {
	case "FAILED":
		*u = UploadStatusIsFailed
	case "INITIALIZED":
		*u = UploadStatusIsInitialized
	case "PROCESSING":
		*u = UploadStatusIsProcessing
	case "SUCCEEDED":
		*u = UploadStatusIsSucceeded
	default:
		return fmt.Errorf("unknown upload status: %q", s)
	}

	return nil
}

const (
	// FAILED: A failed status.
	// INITIALIZED: An initialized status.
	// PROCESSING: A processing status.
	// SUCCEEDED: A succeeded status.
	UploadStatusIsFailed      UploadStatus = "FAILED"
	UploadStatusIsInitialized UploadStatus = "INITIALIZED"
	UploadStatusIsProcessing  UploadStatus = "PROCESSING"
	UploadStatusIsSucceeded   UploadStatus = "SUCCEEDED"
)

type Upload struct {
	ARN      UploadARN      `json:"arn"`
	Name     UploadFileName `json:"name"`
	Type     UploadType     `json:"type"`
	Status   UploadStatus   `json:"status"`
	Metadata string         `json:"metadata"`
	URL      UploadURL      `json:"url"`
}

func NewUpload(uploadARN UploadARN, name UploadFileName, uploadType UploadType, status UploadStatus, metadata string, url UploadURL) Upload {
	return Upload{
		ARN:      uploadARN,
		Name:     name,
		Type:     uploadType,
		Status:   status,
		Metadata: metadata,
		URL:      url,
	}
}

type UploadFileName string

const (
	// NOTE: https://docs.aws.amazon.com/cli/latest/reference/devicefarm/create-upload.html
	// > ANDROID_APP: An Android upload.
	// > IOS_APP: An iOS upload.
	// > ...
	// > APPIUM_NODE_TEST_PACKAGE: An Appium Node.js test package upload.
	// > ...
	// > APPIUM_NODE_TEST_SPEC: An Appium Node.js test spec upload.
	UploadTypeIsAndroidApp UploadType = "ANDROID_APP"
	UploadTypeIsIOSApp     UploadType = "IOS_APP"

	UploadTypeIsAppiumNodeTestPackage UploadType = "APPIUM_NODE_TEST_PACKAGE"
	UploadTypeIsAppiumNodeTestSpec    UploadType = "APPIUM_NODE_TEST_SPEC"

	UploadTypeIsUnknown UploadType = "UNKNOWN"
)

type UploadType string

func (t *UploadType) UnmarshalJSON(bytes []byte) error {
	var s string

	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}

	switch s {
	case string(UploadTypeIsAndroidApp):
		*t = UploadTypeIsAndroidApp
	case string(UploadTypeIsIOSApp):
		*t = UploadTypeIsIOSApp
	case string(UploadTypeIsAppiumNodeTestPackage):
		*t = UploadTypeIsAppiumNodeTestPackage
	case string(UploadTypeIsAppiumNodeTestSpec):
		*t = UploadTypeIsAppiumNodeTestSpec
	default:
		// NOTE: Allow unknowns because only few types are defined and used by devfarm.
		*t = UploadTypeIsUnknown
	}

	return nil
}

func (t UploadType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t UploadType) MIMEType() string {
	// NOTE: https://www.iana.org/assignments/media-types/media-types.xhtml
	switch t {
	case UploadTypeIsAndroidApp, UploadTypeIsIOSApp:
		return "application/octet-stream"
	case UploadTypeIsAppiumNodeTestPackage:
		return "application/gzip"
	case UploadTypeIsAppiumNodeTestSpec:
		return "plain/text"
	default:
		return "application/octet-stream"
	}
}

type UploadURL string

func (t UploadURL) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *UploadURL) UnmarshalJSON(bytes []byte) error {
	var s string

	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}

	*t = UploadURL(s)
	return nil
}
