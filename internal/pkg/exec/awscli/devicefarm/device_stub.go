package devicefarm

import "fmt"

var listDevicesResponseJSONExample = fmt.Sprintf(`{ "devices": [%s, %s] }`,
	deviceIOSJSONExample,
	deviceAndroidJSONExample,
)

var deviceIOSJSONExample = `{
	"arn": "arn:aws:devicefarm:us-west-2::Device:A490B12A656C49678A80B5B0F7D33FA1",
	"name": "Apple iPhone XS",
	"manufacturer": "Apple",
	"model": "iPhone XS",
	"modelId": "{MTAG2,MTA22,MT8J2,MT942}",
	"formFactor": "PHONE",
	"platform": "IOS",
	"os": "12.0",
	"cpu": {
		"frequency": "Hz",
		"architecture": "arm64e",
		"clock": 0.0
	},
	"resolution": {
		"width": 1125,
		"height": 2436
	},
	"heapSize": 0,
	"memory": 64000000000,
	"image": "A490B12A656C49678A80B5B0F7D33FA1",
	"remoteAccessEnabled": true,
	"remoteDebugEnabled": false,
	"fleetType": "PUBLIC",
	"Availability": "HIGHLY_AVAILABLE"
}`

func DeviceIOSExample() Device {
	return NewAWSDeviceFarmDevice(
		"arn:aws:devicefarm:us-west-2::Device:A490B12A656C49678A80B5B0F7D33FA1",
		"Apple",
		"iPhone XS",
		PlatformIsIOS,
		"12.0",
		AvailabilityIsHighlyAvailable,
	)
}

var deviceAndroidJSONExample = `{
	"arn": "arn:aws:devicefarm:us-west-2::Device:4B2B87829E99484DBCD853D82A883BF5",
	"name": "Google Pixel 2",
	"manufacturer": "Pixel 2",
	"model": "Google Pixel 2",
	"modelId": "Google Pixel 2",
	"formFactor": "PHONE",
	"platform": "ANDROID",
	"os": "8.1.0",
	"cpu": {
		"frequency": "MHz",
		"architecture": "arm64-v8a",
		"clock": 1900.8
	},
	"resolution": {
		"width": 1080,
		"height": 1920
	},
	"heapSize": 512000000,
	"memory": 128000000000,
	"image": "4B2B87829E99484DBCD853D82A883BF5",
	"remoteAccessEnabled": false,
	"remoteDebugEnabled": false,
	"fleetType": "PUBLIC",
	"Availability": "HIGHLY_AVAILABLE"
}`

func DeviceAndroidExample() Device {
	return NewAWSDeviceFarmDevice(
		"arn:aws:devicefarm:us-west-2::Device:4B2B87829E99484DBCD853D82A883BF5",
		"Pixel 2",
		"Google Pixel 2",
		PlatformIsAndroid,
		"8.1.0",
		AvailabilityIsHighlyAvailable,
	)
}

func AnyDeviceFarmDevice() Device {
	return DeviceIOSExample()
}
