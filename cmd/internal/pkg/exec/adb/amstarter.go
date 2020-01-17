package adb

type Intent string

type IntentExtras []string

type ActivityStarter func(serialNumber SerialNumber, intent Intent, args IntentExtras) error

func NewActivityStarter(adbCmd Executor) ActivityStarter {
	// > start [options] intent	Start an Activity specified by intent.
	// > See the Specification for intent arguments. Options are:
	// >
	// >   -D: Enable debugging.
	// >   -W: Wait for launch to complete.
	// >   --start-profiler file: Start profiler and send results to file.
	// >   -P file: Like --start-profiler, but profiling stops when the app goes idle.
	// >   -R count: Repeat the activity launch count times. Prior to each repeat, the top activity will be finished.
	// >   -S: Force stop the target app before starting the activity.
	// >   --opengl-trace: Enable tracing of OpenGL functions.
	// >   --user user_id | current: Specify which user to run as; if not specified, then run as the current user.
	//
	// SEE: https://developer.android.com/studio/command-line/adb#am
	return func(serialNumber SerialNumber, intent Intent, intentExtras IntentExtras) error {
		adbArgs := []string{
			"-s", string(serialNumber),
			"shell",
			"am",
			"start",
			"-W",
			"-S",
			"--activity-clear-top",
			"-n", string(intent),
		}
		adbArgs = append(adbArgs, intentExtras...)

		_, err := adbCmd(adbArgs...)
		if err != nil {
			return err
		}
		return nil
	}
}
