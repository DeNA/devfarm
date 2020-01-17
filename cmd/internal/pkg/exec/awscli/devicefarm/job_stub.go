package devicefarm

import "fmt"

var listJobsJSONExample = fmt.Sprintf(`{ "jobs": [ %s ] }`, completedJobJSONExample)

var listJobsExample = []Job{completedJobExample}

var completedJobJSONExample = `{
	"arn": "arn:aws:devicefarm:us-west-2:946725712716:job:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/13fdbdd5-f2f5-4f03-9a8c-91f54cd2b8b8/00000",
	"name": "Google Pixel 3",
	"created": 1562254458.373,
	"status": "COMPLETED",
	"result": "STOPPED",
	"started": 1562254458.373,
	"stopped": 1562254506.682,
	"counters": {
		"total": 0,
		"passed": 0,
		"failed": 0,
		"warned": 0,
		"errored": 0,
		"stopped": 0,
		"skipped": 0
	},
	"device": {
		"arn": "arn:aws:devicefarm:us-west-2::device:CF6DC11E4C99430BA9A1BABAE5B45364",
		"name": "Google Pixel 3",
		"manufacturer": "Google",
		"model": "Google Pixel 3",
		"modelId": "Pixel 3",
		"formFactor": "PHONE",
		"platform": "ANDROID",
		"os": "9",
		"cpu": {
			"frequency": "MHz",
			"architecture": "arm64-v8a",
			"clock": 1766.4
		},
		"resolution": {
			"width": 1080,
			"height": 2160
		},
		"heapSize": 512000000,
		"memory": 64000000000,
		"image": "CF6DC11E4C99430BA9A1BABAE5B45364",
		"remoteAccessEnabled": true,
		"remoteDebugEnabled": false,
		"fleetType": "PUBLIC"
	}
}`

var completedJobExample = NewJob(
	"arn:aws:devicefarm:us-west-2:946725712716:job:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/13fdbdd5-f2f5-4f03-9a8c-91f54cd2b8b8/00000",
	JobStatusIsCompleted,
	NewAWSDeviceFarmDevice(
		"arn:aws:devicefarm:us-west-2::device:CF6DC11E4C99430BA9A1BABAE5B45364",
		"Google",
		"Google Pixel 3",
		PlatformIsAndroid,
		"9",
		AvailabilityIsUnknown,
	),
)
