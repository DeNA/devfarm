package awsdevicefarm

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"github.com/dena/devfarm/internal/pkg/assets"
	"io"
	"os"
	"sort"
	"time"
)

var packageNameInPackageJSON = "noop-tests"
var packageVersionInPackageJSON = "0.0.0"

type testPackageGen func(embedded testPackageEmbeddedData, writer io.Writer) error

// NOTE: Archives npm-bundled packages. This instruction is the official way.
// https://docs.aws.amazon.com/devicefarm/latest/developerguide/test-types-ios-appium-node.html
func newTestPackageGen() testPackageGen {
	return func(embedded testPackageEmbeddedData, writer io.Writer) error {
		zipWriter := zip.NewWriter(writer)
		defer zipWriter.Close()

		if err := writeTestPackageBundle(zipWriter.CreateHeader); err != nil {
			return err
		}

		entryIdx := 0
		embeddedEntries := make([]embeddedDataEntry, len(embedded))
		for filePath, file := range embedded {
			embeddedEntries[entryIdx] = embeddedDataEntry{filePath: filePath, file: file}
			entryIdx++
		}

		sort.Slice(embeddedEntries, func(i, j int) bool {
			return embeddedEntries[i].filePath < embeddedEntries[j].filePath
		})

		for _, entry := range embeddedEntries {
			if err := writeFile(string(entry.filePath), entry.file, zipWriter); err != nil {
				return err
			}
		}

		return nil
	}
}

func newTestPackageEmbeddedData() testPackageEmbeddedData {
	return map[embeddedDataFilePath]embeddedDataFile{
		"devfarmagent/devfarmagent.bash":         {executable: true, data: assets.Read(assets.DevfarmAgentBash)},
		"devfarmagent/linux-amd64/devfarmagent":  {executable: true, data: assets.Read(assets.DevfarmAgentLinuxAMD64)},
		"devfarmagent/darwin-amd64/devfarmagent": {executable: true, data: assets.Read(assets.DevfarmAgentDarwinAMD64)},

		"aws-device-farm/workflows/0-shared.bash":   {executable: true, data: assets.Read(assets.AWSDeviceFarmWorkflowShared)},
		"aws-device-farm/workflows/1-install.bash":  {executable: true, data: assets.Read(assets.AWSDeviceFarmWorkflowInstallStep)},
		"aws-device-farm/workflows/2-pretest.bash":  {executable: true, data: assets.Read(assets.AWSDeviceFarmWorkflowPreTestStep)},
		"aws-device-farm/workflows/3-test.bash":     {executable: true, data: assets.Read(assets.AWSDeviceFarmWorkflowTestStep)},
		"aws-device-farm/workflows/4-posttest.bash": {executable: true, data: assets.Read(assets.AWSDeviceFarmWorkflowPostTestStep)},

		"ios-deploy-agent/package.json":      {executable: false, data: assets.Read(assets.IOSDeployAgentPackageJSON)},
		"ios-deploy-agent/package-lock.json": {executable: false, data: assets.Read(assets.IOSDeployAgentPackageLockJSON)},
	}
}

type testPackageEmbeddedData map[embeddedDataFilePath]embeddedDataFile

type embeddedDataFilePath string

type embeddedDataFile struct {
	executable bool
	data       []byte
}

type embeddedDataEntry struct {
	filePath embeddedDataFilePath
	file     embeddedDataFile
}

// NOTE: Archives to mimic npm-bundle. Using npm-bundle is the official way.
// https://docs.aws.amazon.com/devicefarm/latest/developerguide/test-types-ios-appium-node.html
func writeTestPackageBundle(createHeader func(header *zip.FileHeader) (io.Writer, error)) error {
	basename := fmt.Sprintf("%s-%s", packageNameInPackageJSON, packageVersionInPackageJSON)
	filename := basename + ".tgz"

	header := &zip.FileHeader{
		Name:   filename,
		Method: zip.Deflate,
		// XXX: Avoid matching to zero value.
		Modified: time.Unix(1, 0),
	}
	header.SetMode(0644)

	fileWriter, fileErr := createHeader(header)
	if fileErr != nil {
		return fileErr
	}

	gzipWriter := gzip.NewWriter(fileWriter)
	defer gzipWriter.Close()

	tgzWriter := tar.NewWriter(gzipWriter)
	defer tgzWriter.Close()

	if err := writePackageJSON(tgzWriter); err != nil {
		return err
	}

	return nil
}

func writePackageJSON(tgzWriter *tar.Writer) error {
	// SEE: https://docs.npmjs.com/files/package.json
	packageJSON := []byte(fmt.Sprintf(
		`{ "name": %q, "version": %q }`,
		packageNameInPackageJSON,
		packageVersionInPackageJSON,
	))

	header, headerErr := tar.FileInfoHeader(testPackageFileInfo{
		filepath: "package/package.json",
		size:     int64(len(packageJSON)),
		mode:     0644,
		isDir:    false,
		// XXX: Avoid matching to zero value.
		modTime: time.Unix(1, 0),
	}, "")
	if headerErr != nil {
		return headerErr
	}
	// XXX: Avoid matching to zero value.
	header.AccessTime = time.Unix(1, 0)
	header.ChangeTime = time.Unix(1, 0)

	if err := tgzWriter.WriteHeader(header); err != nil {
		return err
	}

	if _, err := tgzWriter.Write(packageJSON); err != nil {
		return err
	}

	return nil
}

func writeFile(name string, file embeddedDataFile, zipWriter *zip.Writer) error {
	header := &zip.FileHeader{
		Name:   name,
		Method: zip.Deflate,
		// XXX: Avoid matching to zero value.
		Modified: time.Unix(1, 0),
	}
	if file.executable {
		header.SetMode(0755)
	} else {
		header.SetMode(0644)
	}

	fileWriter, fileErr := zipWriter.CreateHeader(header)
	if fileErr != nil {
		return fileErr
	}

	if _, err := fileWriter.Write(file.data); err != nil {
		return err
	}

	return nil
}

type testPackageFileInfo struct {
	filepath string
	size     int64
	mode     os.FileMode
	modTime  time.Time
	isDir    bool
}

func (t testPackageFileInfo) Name() string {
	return t.filepath
}

func (t testPackageFileInfo) Size() int64 {
	return t.size
}

func (t testPackageFileInfo) Mode() os.FileMode {
	return t.mode
}

func (t testPackageFileInfo) ModTime() time.Time {
	return t.modTime
}

func (t testPackageFileInfo) IsDir() bool {
	return t.isDir
}

func (t testPackageFileInfo) Sys() interface{} {
	return nil
}
