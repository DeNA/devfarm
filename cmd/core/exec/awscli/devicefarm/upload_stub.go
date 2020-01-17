package devicefarm

import (
	"fmt"
)

func AnyUpload() Upload {
	return NewUpload(
		"arn:aws:aatp:ANY_UPLOAD",
		"dummy.ipa",
		UploadTypeIsIOSApp,
		UploadStatusIsFailed,
		`{"message":"ANY_METADATA"}`,
		"https://example.com/any-upload",
	)
}

var initializedUploadJSONExample = `{
	"arn": "arn:aws:devicefarm:us-west-2:946725712716:upload:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/15d0d483-8d14-4595-b6dd-dd77418e321c",
	"name": "dummy.ipa",
	"created": 1562080358.228,
	"type": "IOS_APP",
	"status": "INITIALIZED",
	"url": "https://prod-us-west-2-uploads.s3-us-west-2.amazonaws.com/arn%3Aaws%3Adevicefarm%3Aus-west-2%3A946725712716%3Aproject%3Aa935698c-9ab7-4b02-aadf-365bf6dcdbd7/uploads/arn%3Aaws%3Adevicefarm%3Aus-west-2%3A946725712716%3Aupload%3Aa935698c-9ab7-4b02-aadf-365bf6dcdbd7/15d0d483-8d14-4595-b6dd-dd77418e321c/dummy.ipa?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20190702T151238Z&X-Amz-SignedHeaders=host&X-Amz-Expires=86400&X-Amz-Credential=AKIAJSORV74ENYFBITRQ%2F20190702%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Signature=0000000000000000000000000000000000000000000000000000000000000000",
	"category": "PRIVATE"
}`

var initializedUploadExample = NewUpload(
	"arn:aws:devicefarm:us-west-2:946725712716:upload:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/15d0d483-8d14-4595-b6dd-dd77418e321c",
	"dummy.ipa",
	UploadTypeIsIOSApp,
	UploadStatusIsInitialized,
	"",
	"https://prod-us-west-2-uploads.s3-us-west-2.amazonaws.com/arn%3Aaws%3Adevicefarm%3Aus-west-2%3A946725712716%3Aproject%3Aa935698c-9ab7-4b02-aadf-365bf6dcdbd7/uploads/arn%3Aaws%3Adevicefarm%3Aus-west-2%3A946725712716%3Aupload%3Aa935698c-9ab7-4b02-aadf-365bf6dcdbd7/15d0d483-8d14-4595-b6dd-dd77418e321c/dummy.ipa?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20190702T151238Z&X-Amz-SignedHeaders=host&X-Amz-Expires=86400&X-Amz-Credential=AKIAJSORV74ENYFBITRQ%2F20190702%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Signature=0000000000000000000000000000000000000000000000000000000000000000",
)

var createUploadJSONExample = fmt.Sprintf(`{"upload": %s}`, initializedUploadJSONExample)

var succeededUploadJSONExample = `{
	"arn": "arn:aws:aatp:us-west-2:946725712716:upload:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/ccc7bf00-50e0-4728-81cd-90976901e469",
	"name": "example.apk",
	"created": 1562083587.411,
	"type": "ANDROID_APP",
	"status": "SUCCEEDED",
	"url": "https://prod-us-west-2-uploads.s3-us-west-2.amazonaws.com/arn%3Aaws%3Adevicefarm%3Aus-west-2%3A946725712716%3Aproject%3Aa935698c-9ab7-4b02-aadf-365bf6dcdbd7/uploads/arn%3Aaws%3Adevicefarm%3Aus-west-2%3A946725712716%3Aupload%3Aa935698c-9ab7-4b02-aadf-365bf6dcdbd7/ccc7bf00-50e0-4728-81cd-90976901e469/example.apk?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20190702T170211Z&X-Amz-SignedHeaders=host&X-Amz-Expires=86400&X-Amz-Credential=AKIAJSORV74ENYFBITRQ%2F20190702%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Signature=0000000000000000000000000000000000000000000000000000000000000000",
	"metadata": "{\"device_admin\":false,\"activity_name\":\"com.unity3d.player.UnityPlayerActivity\",\"version_name\":\"1.0\",\"screens\":[\"small\",\"normal\",\"large\",\"xlarge\"],\"error_type\":null,\"sdk_version\":\"23\",\"package_name\":\"com.DeNA.GODLIKE\",\"version_code\":\"1\",\"native_code\":[\"armeabi-v7a\",\"x86\"],\"target_sdk_version\":\"29\"}",
	"category": "PRIVATE"
}`

var succeededUploadExample = NewUpload(
	"arn:aws:aatp:us-west-2:946725712716:upload:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/ccc7bf00-50e0-4728-81cd-90976901e469",
	"example.apk",
	UploadTypeIsAndroidApp,
	UploadStatusIsSucceeded,
	`{"device_admin":false,"activity_name":"com.unity3d.player.UnityPlayerActivity","version_name":"1.0","screens":["small","normal","large","xlarge"],"error_type":null,"sdk_version":"23","package_name":"com.DeNA.GODLIKE","version_code":"1","native_code":["armeabi-v7a","x86"],"target_sdk_version":"29"}`,
	"https://prod-us-west-2-uploads.s3-us-west-2.amazonaws.com/arn%3Aaws%3Adevicefarm%3Aus-west-2%3A946725712716%3Aproject%3Aa935698c-9ab7-4b02-aadf-365bf6dcdbd7/uploads/arn%3Aaws%3Adevicefarm%3Aus-west-2%3A946725712716%3Aupload%3Aa935698c-9ab7-4b02-aadf-365bf6dcdbd7/ccc7bf00-50e0-4728-81cd-90976901e469/example.apk?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20190702T170211Z&X-Amz-SignedHeaders=host&X-Amz-Expires=86400&X-Amz-Credential=AKIAJSORV74ENYFBITRQ%2F20190702%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Signature=0000000000000000000000000000000000000000000000000000000000000000",
)

var getUploadJSONExample = fmt.Sprintf(`{"upload": %s}`, succeededUploadJSONExample)
var getUploadExample = succeededUploadExample

var listUploadsJSONExample = fmt.Sprintf(`{"uploads": [%s]}`, succeededUploadJSONExample)

var listUploadsExample = []Upload{
	succeededUploadExample,
}
