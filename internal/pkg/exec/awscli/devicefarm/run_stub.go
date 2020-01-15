package devicefarm

import "fmt"

var scheduleRunJSONExample = fmt.Sprintf(`{ "run": %s }`, pendingRunJSONExample)

var pendingRunJSONExample = `{
	"arn": "arn:aws:devicefarm:us-west-2:946725712716:run:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/13fdbdd5-f2f5-4f03-9a8c-91f54cd2b8b8",
	"name": "example.apk",
	"type": "APPIUM_NODE",
	"platform": "ANDROID_APP",
	"created": 1562254458.357,
	"status": "SCHEDULING",
	"result": "PENDING",
	"started": 1562254458.357,
	"counters": {
		"total": 0,
		"passed": 0,
		"failed": 0,
		"warned": 0,
		"errored": 0,
		"stopped": 0,
		"skipped": 0
	},
	"totalJobs": 1,
	"completedJobs": 0,
	"billingMethod": "METERED",
	"appUpload": "arn:aws:devicefarm:us-west-2:946725712716:upload:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/ccc7bf00-50e0-4728-81cd-90976901e469",
	"jobTimeoutMinutes": 150,
	"devicePoolArn": "arn:aws:devicefarm:us-west-2:946725712716:devicepool:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/761ceade-2939-4294-a7ca-731eb135cd75",
	"radios": {
		"wifi": true,
		"bluetooth": false,
		"nfc": true,
		"gps": true
	},
	"testSpecArn": "arn:aws:devicefarm:us-west-2:946725712716:upload:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/370752d5-34e3-4103-82f9-10ef65dbd81f"
}`

var pendingRunExample = NewRun(
	"arn:aws:devicefarm:us-west-2:946725712716:run:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/13fdbdd5-f2f5-4f03-9a8c-91f54cd2b8b8",
	RunStatusIsScheduling,
	RunResultIsPending,
	"arn:aws:devicefarm:us-west-2:946725712716:devicepool:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/761ceade-2939-4294-a7ca-731eb135cd75",
)

var stoppedRunJSONExample = ` {
	"arn": "arn:aws:devicefarm:us-west-2:946725712716:run:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/13fdbdd5-f2f5-4f03-9a8c-91f54cd2b8b8",
	"name": "example.apk",
	"type": "APPIUM_NODE",
	"platform": "ANDROID_APP",
	"created": 1562254458.357,
	"status": "COMPLETED",
	"result": "STOPPED",
	"started": 1562254458.357,
	"stopped": 1562254507.596,
	"counters": {
		"total": 0,
		"passed": 0,
		"failed": 0,
		"warned": 0,
		"errored": 0,
		"stopped": 0,
		"skipped": 0
	},
	"totalJobs": 1,
	"completedJobs": 1,
	"billingMethod": "METERED",
	"appUpload": "arn:aws:devicefarm:us-west-2:946725712716:upload:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/ccc7bf00-50e0-4728-81cd-90976901e469",
	"jobTimeoutMinutes": 150,
	"devicePoolArn": "arn:aws:devicefarm:us-west-2:946725712716:devicepool:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/761ceade-2939-4294-a7ca-731eb135cd75",
	"radios": {
		"wifi": true,
		"bluetooth": false,
		"nfc": true,
		"gps": true
	},
	"testSpecArn": "arn:aws:devicefarm:us-west-2:946725712716:upload:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/370752d5-34e3-4103-82f9-10ef65dbd81f"
}`

var stoppedRunExample = NewRun(
	"arn:aws:devicefarm:us-west-2:946725712716:run:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/13fdbdd5-f2f5-4f03-9a8c-91f54cd2b8b8",
	RunStatusIsCompleted,
	RunResultIsStopped,
	"arn:aws:devicefarm:us-west-2:946725712716:devicepool:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/761ceade-2939-4294-a7ca-731eb135cd75",
)

var passedRunJSONExample = `{
	"arn": "arn:aws:devicefarm:us-west-2:946725712716:run:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/1dedadce-7268-49ac-8d8c-9bdcefe7c050",
	"name": "example.apk",
	"type": "APPIUM_NODE",
	"platform": "ANDROID_APP",
	"created": 1562154530.992,
	"status": "COMPLETED",
	"result": "PASSED",
	"started": 1562154530.992,
	"stopped": 1562155005.991,
	"counters": {
		"total": 3,
		"passed": 3,
		"failed": 0,
		"warned": 0,
		"errored": 0,
		"stopped": 0,
		"skipped": 0
	},
	"totalJobs": 1,
	"completedJobs": 1,
	"billingMethod": "METERED",
	"deviceMinutes": {
		"total": 6.23,
		"metered": 4.21,
		"unmetered": 0.0
	},
	"appUpload": "arn:aws:devicefarm:us-west-2:946725712716:upload:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/ccc7bf00-50e0-4728-81cd-90976901e469",
	"jobTimeoutMinutes": 150,
	"devicePoolArn": "arn:aws:devicefarm:us-west-2:946725712716:devicepool:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/761ceade-2939-4294-a7ca-731eb135cd75",
	"radios": {
		"wifi": true,
		"bluetooth": false,
		"nfc": true,
		"gps": true
	},
	"testSpecArn": "arn:aws:devicefarm:us-west-2:946725712716:upload:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/370752d5-34e3-4103-82f9-10ef65dbd81f"
}`

var passedRunExample = NewRun(
	"arn:aws:devicefarm:us-west-2:946725712716:run:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/1dedadce-7268-49ac-8d8c-9bdcefe7c050",
	RunStatusIsCompleted,
	RunResultIsPassed,
	"arn:aws:devicefarm:us-west-2:946725712716:devicepool:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/761ceade-2939-4294-a7ca-731eb135cd75",
)

var listRunsJSONExample = fmt.Sprintf(`{ "runs": [ %s, %s ] }`, stoppedRunJSONExample, passedRunJSONExample)

var listRunsExample = []Run{stoppedRunExample, passedRunExample}

var getRunJSONExample = `{
    "run": {
        "arn": "arn:aws:devicefarm:us-west-2:946725712716:run:8c2406f2-c822-48a9-bd4f-8d55f8dc4a48/0da49866-239a-4c2a-bb49-e3e6418f093e",
        "name": "devfarm-Args-f3f42131.ipa",
        "type": "APPIUM_NODE",
        "platform": "IOS_APP",
        "created": 1568730216.271,
        "status": "COMPLETED",
        "result": "FAILED",
        "started": 1568730216.271,
        "stopped": 1568730657.168,
        "counters": {
            "total": 3,
            "passed": 2,
            "failed": 1,
            "warned": 0,
            "errored": 0,
            "stopped": 0,
            "skipped": 0
        },
        "totalJobs": 1,
        "completedJobs": 1,
        "billingMethod": "METERED",
        "deviceMinutes": {
            "total": 6.1,
            "metered": 5.34,
            "unmetered": 0.0
        },
        "appUpload": "arn:aws:devicefarm:us-west-2:946725712716:upload:8c2406f2-c822-48a9-bd4f-8d55f8dc4a48/a40cb5bf-c44d-4fd6-805e-b249ca724772",
        "jobTimeoutMinutes": 60,
        "devicePoolArn": "arn:aws:devicefarm:us-west-2:946725712716:devicepool:8c2406f2-c822-48a9-bd4f-8d55f8dc4a48/d8b4d906-bd9e-4e60-8d57-bc65431b8bbd",
        "radios": {
            "wifi": true,
            "bluetooth": false,
            "nfc": true,
            "gps": true
        },
        "skipAppResign": false,
        "testSpecArn": "arn:aws:devicefarm:us-west-2:946725712716:upload:8c2406f2-c822-48a9-bd4f-8d55f8dc4a48/d86ff54c-44dc-48ba-bf3a-c2198c376786"
    }
}`
var getRunExample = Run{
	ARN:           "arn:aws:devicefarm:us-west-2:946725712716:run:8c2406f2-c822-48a9-bd4f-8d55f8dc4a48/0da49866-239a-4c2a-bb49-e3e6418f093e",
	Status:        RunStatusIsCompleted,
	Result:        RunResultIsFailed,
	DevicePoolARN: "arn:aws:devicefarm:us-west-2:946725712716:devicepool:8c2406f2-c822-48a9-bd4f-8d55f8dc4a48/d8b4d906-bd9e-4e60-8d57-bc65431b8bbd",
}
