package adb

type ApkInstaller func(serialNumber SerialNumber, appPath string) error

func NewApkInstaller(adbCmd Executor) ApkInstaller {
	return func(serialNumber SerialNumber, appPath string) error {
		// > Push packages to the device and install them. Possible options are the following:
		// >
		// >   --abi abi-identifier : Force install an app for a specific ABI.
		// >   -l: Forward lock app.
		// >   -r: Replace the existing app.
		// >   -t: Allow test packages. If the APK is built using a developer preview SDK (if the targetSdkVersion is a letter instead of a number), you must include the -t option with the install command if you are installing a test APK. For more information see -t option.
		// >   -s: Install the app on the SD card.
		// >   -d: Allow version code downgrade (debugging packages only).
		// >   -g: Grant all runtime permissions.
		//
		// SEE: https://developer.android.com/studio/command-line/adb#am
		_, err := adbCmd("-s", string(serialNumber), "install", "-r", "-g", appPath)
		if err != nil {
			return err
		}
		return nil
	}
}
