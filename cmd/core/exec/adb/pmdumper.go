package adb

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

type PackageName string

type MainIntentFinder func(serialNumber SerialNumber, packageName PackageName) (Intent, error)

func NewMainIntentFinder(adbCmd Executor) MainIntentFinder {
	return func(serialNumber SerialNumber, packageName PackageName) (Intent, error) {
		stdout, err := adbCmd("-s", string(serialNumber), "shell", "pm", "dump", string(packageName))
		if err != nil {
			return "", err
		}

		// EXAMPLE:
		// > $ adb shell pm dump com.example.package
		// > DUMP OF SERVICE package:
		// >  Activity Resolver Table:
		// >    Non-Data Actions:
		// >        android.intent.action.MAIN:
		// >          aba9035 com.example.package/ExampleActivity filter 8416d53
		scanner := bufio.NewScanner(bytes.NewReader(stdout))
		isInActivityLine := false

		for scanner.Scan() {
			line := scanner.Text()
			trimmedLine := strings.TrimSpace(line)

			if isInActivityLine {
				components := strings.Split(trimmedLine, " ")

				if len(components) < 2 {
					return "", fmt.Errorf("unrecognizable intent description: %q", trimmedLine)
				}

				unsafeMainIntent := components[1]

				seemsValid := strings.HasPrefix(unsafeMainIntent, string(packageName))
				if !seemsValid {
					return "", fmt.Errorf("invalid intent: %q (in %q)", unsafeMainIntent, trimmedLine)
				}

				return Intent(unsafeMainIntent), nil
			}

			isAndroidIntentActionMain := strings.HasPrefix(trimmedLine, `android.intent.action.MAIN:`)
			if !isAndroidIntentActionMain {
				continue
			} else {
				isInActivityLine = true
				continue
			}
		}

		return "", fmt.Errorf("android.intent.action.MAIN not found:\n%s", stdout)
	}
}
