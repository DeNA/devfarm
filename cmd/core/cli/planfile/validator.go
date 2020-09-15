package planfile

import (
	"errors"
	"fmt"
	"github.com/dena/devfarm/cmd/core/exec"
	"github.com/dena/devfarm/cmd/core/platforms"
	"os"
	"path/filepath"
	"time"
)

type ValidateFunc func(unsafePlanfile UnsafePlanFile) (Planfile, error)

func NewValidateFunc(stat exec.StatFunc) ValidateFunc {
	validateAsIOS := NewIOSValidateFunc(stat)
	validateAsAndroid := NewAndroidValidateFunc(stat)
	return func(unsafePlanfile UnsafePlanFile) (Planfile, error) {
		errs := make([]error, 0)
		planFile := Planfile{}

		for rawGroupName, unsafePlans := range unsafePlanfile.InstanceGroups {
			groupName := platforms.InstanceGroupName(rawGroupName)

			for unsafePlanIdx, unsafePlan := range unsafePlans {
				var plan platforms.EitherPlan
				locationHint := platforms.LocationHintForPlanfileIndex(unsafePlanIdx, groupName)

				osName := unsafePlan.DetectOS()
				switch osName {
				case platforms.OSIsIOS:
					iosPlan, iosErr := validateAsIOS(unsafePlanfile.Path, unsafePlan, groupName, locationHint)
					if iosErr != nil {
						errs = append(errs, iosErr)
						continue
					}
					plan = iosPlan.Either()

				case platforms.OSIsAndroid:
					androidPlan, androidErr := validateAsAndroid(unsafePlanfile.Path, unsafePlan, groupName, locationHint)
					if androidErr != nil {
						errs = append(errs, androidErr)
						continue
					}
					plan = androidPlan.Either()

				default:
					err := SemanticError{
						errors:       []error{fmt.Errorf("not supported OS: %q", osName)},
						locationHint: platforms.LocationHintForPlanfileIndex(unsafePlanIdx, groupName),
					}
					errs = append(errs, err)
					continue
				}

				planFile.AddPlan(plan)
			}
		}

		if len(errs) > 0 {
			return Planfile{}, SemanticError{errors: errs}
		}

		return planFile, nil
	}
}

type IOSValidateFunc func(unsafePlanPath FilePath, unsafePlan UnsafePlan, groupName platforms.InstanceGroupName, locationHint string) (platforms.IOSPlan, error)

func NewIOSValidateFunc(stat exec.StatFunc) IOSValidateFunc {
	return func(unsafePlanPath FilePath, unsafePlan UnsafePlan, groupName platforms.InstanceGroupName, locationHint string) (platforms.IOSPlan, error) {
		errs := make([]error, 0)

		var platform platforms.ID
		if len(unsafePlan.Platform) > 0 {
			platform = platforms.ID(unsafePlan.Platform)
		} else {
			errs = append(errs, errors.New("platform: must not be empty"))
		}

		var osVersion platforms.IOSVersion
		if len(unsafePlan.IOS) > 0 {
			osVersion = platforms.IOSVersion(unsafePlan.IOS)
		} else {
			errs = append(errs, errors.New("ios: must not be empty"))
		}

		var deviceName platforms.IOSDeviceName
		if len(unsafePlan.Device) > 0 {
			deviceName = platforms.IOSDeviceName(unsafePlan.Device)
		} else {
			errs = append(errs, errors.New("device: must not be empty"))
		}

		var ipa platforms.IPAPathOnLocal
		if len(unsafePlan.IPA) > 0 {
			unsafeIpaPathMaybeRel := os.ExpandEnv(unsafePlan.IPA)
			var unsafeIpaPathAbs string
			if filepath.IsAbs(unsafeIpaPathMaybeRel) {
				unsafeIpaPathAbs = unsafeIpaPathMaybeRel
			} else {
				unsafeIpaPathAbs = filepath.Join(filepath.Dir(string(unsafePlanPath)), unsafeIpaPathMaybeRel)
			}
			if _, err := stat(unsafeIpaPathAbs); err != nil {
				errs = append(errs, err)
			} else {
				ipa = platforms.IPAPathOnLocal(unsafeIpaPathAbs)
			}
		} else {
			errs = append(errs, errors.New("ipa: must not be empty"))
		}

		var args platforms.IOSArgs
		if unsafePlan.IOSArgs == nil {
			args = platforms.IOSArgs{}
		} else {
			args = unsafePlan.IOSArgs
		}

		var lifetime time.Duration
		if unsafePlan.LifetimeSec > 0 {
			lifetime = time.Duration(unsafePlan.LifetimeSec) * time.Second
		} else {
			errs = append(errs, errors.New("lifetime_sec: must be greater than 0"))
		}

		if len(errs) > 0 {
			return platforms.IOSPlan{}, SemanticError{errors: errs}
		}

		return platforms.NewIOSPlan(
			platform,
			groupName,
			platforms.NewIOSDevice(deviceName, osVersion),
			ipa,
			args,
			lifetime,
			locationHint,
		), nil
	}
}

type AndroidValidateFunc func(unsafePlanfilePath FilePath, unsafePlan UnsafePlan, groupName platforms.InstanceGroupName, locationHint string) (platforms.AndroidPlan, error)

func NewAndroidValidateFunc(stat exec.StatFunc) AndroidValidateFunc {
	return func(unsafePlanfilePath FilePath, unsafePlan UnsafePlan, groupName platforms.InstanceGroupName, locationHint string) (platforms.AndroidPlan, error) {
		errs := make([]error, 0)

		var platform platforms.ID
		if len(unsafePlan.Platform) > 0 {
			platform = platforms.ID(unsafePlan.Platform)
		} else {
			errs = append(errs, errors.New("platform: must not be empty"))
		}

		var osVersion platforms.AndroidVersion
		if len(unsafePlan.Android) > 0 {
			osVersion = platforms.AndroidVersion(unsafePlan.Android)
		} else {
			errs = append(errs, errors.New("android: must not be empty"))
		}

		var deviceName platforms.AndroidDeviceName
		if len(unsafePlan.Device) > 0 {
			deviceName = platforms.AndroidDeviceName(unsafePlan.Device)
		} else {
			errs = append(errs, errors.New("device: must not be empty"))
		}

		var apk platforms.APKPathOnLocal
		if len(unsafePlan.APK) > 0 {
			unsafeApkPathMaybeRel := os.ExpandEnv(unsafePlan.APK)
			var unsafeApkPathAbs string
			if filepath.IsAbs(unsafeApkPathMaybeRel) {
				unsafeApkPathAbs = unsafeApkPathMaybeRel
			} else {
				unsafeApkPathAbs = filepath.Join(filepath.Dir(string(unsafePlanfilePath)), unsafeApkPathMaybeRel)
			}

			if _, err := stat(unsafeApkPathAbs); err != nil {
				errs = append(errs, err)
			} else {
				apk = platforms.APKPathOnLocal(unsafeApkPathAbs)
			}
		} else {
			errs = append(errs, errors.New("apk: must not be empty"))
		}

		var appID platforms.AndroidAppID
		if len(unsafePlan.AndroidAppID) > 0 {
			appID = platforms.AndroidAppID(unsafePlan.AndroidAppID)
		} else {
			errs = append(errs, errors.New("app_id: must not be empty"))
		}

		var intentExtras platforms.AndroidIntentExtras
		if unsafePlan.IntentExtras == nil {
			intentExtras = platforms.AndroidIntentExtras{}
		} else {
			intentExtras = unsafePlan.IntentExtras
		}

		var lifetime time.Duration
		if unsafePlan.LifetimeSec > 0 {
			lifetime = time.Duration(unsafePlan.LifetimeSec) * time.Second
		} else {
			errs = append(errs, errors.New("lifetime_sec: must be greater than 0"))
		}

		if len(errs) > 0 {
			return platforms.AndroidPlan{}, SemanticError{errors: errs}
		}

		return platforms.NewAndroidPlan(
			platform,
			groupName,
			platforms.NewAndroidDevice(deviceName, osVersion),
			apk,
			appID,
			intentExtras,
			lifetime,
			locationHint,
		), nil
	}
}
