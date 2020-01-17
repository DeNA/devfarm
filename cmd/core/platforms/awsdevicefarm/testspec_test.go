package awsdevicefarm

import (
	"bufio"
	"github.com/dena/devfarm/cmd/core/logging"
	"github.com/dena/devfarm/cmd/core/platforms"
	"strings"
	"testing"
)

func TestNewTestSpecUploader(t *testing.T) {
	spyLogger := logging.SpySeverityLogger()
	reserveAndUploaderCallArgs, spyReserveAndUploader := spyReserveUploaderIfNotExists(anySuccessfulReserveAndUploaderIfNotExists())

	uploadTestSpec := newTestSpecUploader(spyLogger, platforms.NewCRC32Hasher(), spyReserveAndUploader)

	_, err := uploadTestSpec("arn:aws:devicefarm:ANY_PROJECT", anyTestSpec())
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

func TestGenerateCustomTestEnvSpec(t *testing.T) {
	yaml, err := generateCustomTestEnvSpec(anyTestSpecEmbeddedData{})
	if err != nil {
		t.Errorf("got %v, want nil", err)
		return
	}

	scanner := bufio.NewScanner(strings.NewReader(string(yaml)))

	for scanner.Scan() {
		line := scanner.Text()

		isComment := strings.HasPrefix(strings.TrimSpace(line), "#")

		if !isComment && strings.Contains(line, `'`) {
			t.Errorf("do no use single-quote because AWS Device Farm will embed the string in `echo '...'`")
			t.Logf("this line: %q", line)
		}
	}
}
