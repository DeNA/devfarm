package foreverios

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
	device            platforms.IOSDevice
	ipaPath           platforms.IPAPathOnLocal
	iosArgs           platforms.IOSArgs
	lifetime          time.Duration
}

func takeOptions(args []string) (options, *cli.ErrorAndUsage) {
	flags, usageBuf := cli.NewFlagSet([]string{})

	verbose := cli.DefineVerboseOpts(flags)
	dryRun := cli.DefineDryRunOpts(flags)
	unsafeInstanceGroupName := cli.DefineInstanceGroupNameOpts(flags)
	unsafePlatformName := flags.String("platform", "", "platform name listed by 'devfarm list-devices' (required)")
	unsafeDeviceName := flags.String("device", "", "device name listed by 'devfarm list-devices' (required)")
	unsafeIOSVersion := flags.String("os-version", "", "iOS version listed by 'devfarm list-devices' (required)")
	unsafeIPAPath := flags.String("ipa", "", "ipa file to launch (required)")
	unsafeAppArgs := flags.String("args-json", "[]", "arguments that will be passed to the iOS app (via ProcessInfo#arguments) after decoding to plain arguments")
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

	iosVersion, osVersionErr := validateIOSVersion(*unsafeIOSVersion)
	if osVersionErr != nil {
		flags.Usage()
		return options{}, cli.NewErrorAndUsage(
			fmt.Sprintf("--os-version: %s", osVersionErr.Error()),
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

	ipaPath, ipaPathErr := validateIPAPath(*unsafeIPAPath)
	if ipaPathErr != nil {
		flags.Usage()
		return options{}, cli.NewErrorAndUsage(
			fmt.Sprintf("--ipa: %s", ipaPathErr.Error()),
			usageBuf.String(),
		)
	}

	iosArgs, iosArgsErr := validateIOSArgs(*unsafeAppArgs)
	if iosArgsErr != nil {
		flags.Usage()
		return options{}, cli.NewErrorAndUsage(
			fmt.Sprintf("--args-json: %s", iosArgsErr.Error()),
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
		device:            platforms.NewIOSDevice(deviceName, iosVersion),
		ipaPath:           ipaPath,
		iosArgs:           iosArgs,
		lifetime:          lifetime,
	}, nil
}

func validateIOSVersion(unsafeIOSVersion string) (platforms.IOSVersion, error) {
	if len(unsafeIOSVersion) < 1 {
		return "", errors.New("must not be empty")
	}

	return platforms.IOSVersion(unsafeIOSVersion), nil
}

func validateIPAPath(unsafeIPAPath string) (platforms.IPAPathOnLocal, error) {
	if len(unsafeIPAPath) < 1 {
		return "", errors.New("must not be empty")
	}

	file, openErr := os.Open(unsafeIPAPath)
	if openErr != nil {
		return "", openErr
	}
	_ = file.Close() // XXX: devfarm processes do not live long time. So it should not problem.

	return platforms.IPAPathOnLocal(unsafeIPAPath), nil
}

func validateIOSArgs(unsafeIOSArgsJSON string) (platforms.IOSArgs, error) {
	var iosArgs platforms.IOSArgs

	if err := json.Unmarshal([]byte(unsafeIOSArgsJSON), &iosArgs); err != nil {
		return nil, err
	}

	return iosArgs, nil
}

func validateDeviceName(unsafeDeviceName string) (platforms.IOSDeviceName, error) {
	if len(unsafeDeviceName) < 1 {
		return "", errors.New("must not be empty")
	}

	return platforms.IOSDeviceName(unsafeDeviceName), nil
}

func validateLifetime(unsafeLifetimeSec int) (time.Duration, error) {
	if unsafeLifetimeSec < 1 {
		return 0, errors.New("must be greater than 0")
	}
	return time.Duration(unsafeLifetimeSec) * time.Second, nil
}
