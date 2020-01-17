package awsdevicefarm

import (
	"github.com/dena/devfarm/cmd/internal/pkg/exec/awscli/devicefarm"
)

var devfarmUploadNamePrefix = "devfarm-"

type uploadProperty struct {
	fileName   devicefarm.UploadFileName
	uploadType devicefarm.UploadType
}

func newUploadProperty(fileName devicefarm.UploadFileName, uploadType devicefarm.UploadType) uploadProperty {
	return uploadProperty{
		fileName:   fileName,
		uploadType: uploadType,
	}
}