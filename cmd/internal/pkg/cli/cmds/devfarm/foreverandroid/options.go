package foreverandroid

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dena/devfarm/cmd/internal/pkg/cli"
	"github.com/dena/devfarm/cmd/internal/pkg/platforms"
	"github.com/dena/devfarm/cmd/internal/pkg/platforms/all"
	"os"
	"time"
)

type options struct {
	verbose           bool
	dryRun            bool
	instanceGroupName platforms.InstanceGroupName
	platformID        platforms.ID
	device            platforms.AndroidDevice
	apkPath           platforms.APKPathOnLocal
	appID             platforms.AndroidAppID
	intentExtras      platforms.AndroidIntentExtras
	lifetime          time.Duration
}

func takeOptions(args []string) (options, *cli.ErrorAndUsage) {
	flags, usageBuf := cli.NewFlagSet([]string{})

	verbose := cli.DefineVerboseOpts(flags)
	dryRun := cli.DefineDryRunOpts(flags)
	unsafeInstanceGroupName := cli.DefineInstanceGroupNameOpts(flags)
	unsafePlatformName := flags.String("platform", "", "platform name listed by 'devfarm list-devices' (required)")
	unsafeDeviceName := flags.String("device", "", "device name listed by 'devfarm list-devices' (required)")
	unsafeAndroidVersion := flags.String("os-version", "", "Android version listed by 'devfarm list-devices' (required)")
	unsafeAPKPath := flags.String("apk", "", "apk file to launch (required)")
	unsafeAppID := flags.String("app-id", "", "application ID (it often called as 'package name') to the app (required)")
	unsafeIntentExtras := flags.String("intent-extras-json", "[]", "arguments that will be passed to the Android app (via Intent#getExtras) after decoding to plain arguments")
	unsafeLifetimeSec := flags.Int("lifetime-sec", 900, "duration during app launching")

	if err := cli.Parse(args, flags, usageBuf); err != nil {
		return options{}, err
	}

	instanceGroupName, groupErr := platforms.NewInstanceGroupName(*unsafeInstanceGroupName)
	if groupErr != nil {
		flags.Usage()
		return options{}, cli.NewErrorAndUsage(
			fmt.Sprintf("--instance-group: %s", groupErr.Error()),
			usageBuf.String(),
		)
	}

	platformID, platformErr := all.ValidatePlatformID(*unsafePlatformName)
	if platformErr != nil {
		flags.Usage()
		return options{}, cli.NewErrorAndUsage(
			fmt.Sprintf("--platform: %s", platformErr.Error()),
			usageBuf.String(),
		)
	}

	androidVersion, androidVersionErr := validateAndroidVersion(*unsafeAndroidVersion)
	if androidVersionErr != nil {
		flags.Usage()
		return options{}, cli.NewErrorAndUsage(
			fmt.Sprintf("--os-version: %s", androidVersionErr.Error()),
			usageBuf.String(),
		)
	}

	deviceName, deviceNameErr := validateDeviceName(*unsafeDeviceName)
	if deviceNameErr != nil {
		flags.Usage()
		return options{}, cli.NewErrorAndUsage(
			fmt.Sprintf("--device: %s", deviceNameErr.Error()),
			usageBuf.String(),
		)
	}

	apkPath, apkPathErr := validateAPKPath(*unsafeAPKPath)
	if apkPathErr != nil {
		flags.Usage()
		return options{}, cli.NewErrorAndUsage(
			fmt.Sprintf("--apk: %s", apkPathErr.Error()),
			usageBuf.String(),
		)
	}

	appID, appIDErr := validateAppID(*unsafeAppID)
	if appIDErr != nil {
		flags.Usage()
		return options{}, cli.NewErrorAndUsage(
			fmt.Sprintf("--app-id: %s", appIDErr.Error()),
			usageBuf.String(),
		)
	}

	intentExtras, intentExtrasErr := validateIntentExtras(*unsafeIntentExtras)
	if intentExtrasErr != nil {
		flags.Usage()
		return options{}, cli.NewErrorAndUsage(
			fmt.Sprintf("--intent-extras-json: %s", intentExtrasErr.Error()),
			usageBuf.String(),
		)
	}

	lifetime, lifetimeErr := validateLifetime(*unsafeLifetimeSec)
	if lifetimeErr != nil {
		flags.Usage()
		return options{}, cli.NewErrorAndUsage(
			fmt.Sprintf("--lifetime-sec: %s", lifetimeErr.Error()),
			usageBuf.String(),
		)
	}

	return options{
		verbose:           *verbose,
		dryRun:            *dryRun,
		instanceGroupName: instanceGroupName,
		platformID:        platformID,
		device:            platforms.NewAndroidDevice(deviceName, androidVersion),
		apkPath:           apkPath,
		appID:             appID,
		intentExtras:      intentExtras,
		lifetime:          lifetime,
	}, nil
}

func validateAndroidVersion(unsafeAndroidVersion string) (platforms.AndroidVersion, error) {
	if len(unsafeAndroidVersion) < 1 {
		return "", errors.New("must not be empty")
	}

	return platforms.AndroidVersion(unsafeAndroidVersion), nil
}

func validateAPKPath(unsafeAPKPath string) (platforms.APKPathOnLocal, error) {
	if len(unsafeAPKPath) < 1 {
		return "", errors.New("must not be empty")
	}

	file, openErr := os.Open(unsafeAPKPath)
	if openErr != nil {
		return "", openErr
	}
	_ = file.Close() // XXX: devfarm processes do not live long time. So it should not problem.

	return platforms.APKPathOnLocal(unsafeAPKPath), nil
}

func validateAppID(unsafeAppID string) (platforms.AndroidAppID, error) {
	if len(unsafeAppID) < 1 {
		return "", errors.New("must not be empty")
	}

	return platforms.AndroidAppID(unsafeAppID), nil
}

func validateIntentExtras(unsafeIntentExtrasJSON string) (platforms.AndroidIntentExtras, error) {
	var intentExtras platforms.AndroidIntentExtras

	if err := json.Unmarshal([]byte(unsafeIntentExtrasJSON), &intentExtras); err != nil {
		return nil, err
	}

	return intentExtras, nil
}

func validateDeviceName(unsafeDeviceName string) (platforms.AndroidDeviceName, error) {
	if len(unsafeDeviceName) < 1 {
		return "", errors.New("must not be empty")
	}

	return platforms.AndroidDeviceName(unsafeDeviceName), nil
}

func validateLifetime(unsafeLifetimeSec int) (time.Duration, error) {
	if unsafeLifetimeSec < 1 {
		return 0, errors.New("must be greater than 0")
	}
	return time.Duration(unsafeLifetimeSec) * time.Second, nil
}
