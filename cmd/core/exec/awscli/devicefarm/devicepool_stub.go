package devicefarm

import "fmt"

var createDevicePoolJSONExample = fmt.Sprintf(`{ "devicePool": %s }`, devicePoolJSONExample)

var devicePoolJSONExample = `{
	"arn": "arn:aws:devicefarm:us-west-2:946725712716:devicepool:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/538891ad-f279-4c3f-85fd-a10f1191e81d",
	"name": "name",
	"description": "description",
	"type": "PRIVATE",
	"rules": [
		{
			"attribute": "ARN",
			"operator": "IN",
			"value": "[\"arn:aws:devicefarm:us-west-2::device:CF6DC11E4C99430BA9A1BABAE5B45364\"]"
		}
	]
}`

var devicePoolExample = NewDevicePool(
	"arn:aws:devicefarm:us-west-2:946725712716:devicepool:a935698c-9ab7-4b02-aadf-365bf6dcdbd7/538891ad-f279-4c3f-85fd-a10f1191e81d",
	"name",
	"description",
	[]DevicePoolRule{
		{
			Attribute: "ARN",
			Operator:  "IN",
			Value:     `["arn:aws:devicefarm:us-west-2::device:CF6DC11E4C99430BA9A1BABAE5B45364"]`,
		},
	},
)

var listDevicePoolsJSONExample = fmt.Sprintf(`{
    "devicePools": [
        {
            "arn": "arn:aws:devicefarm:us-west-2::devicepool:082d10e5-d7d7-48a5-ba5c-b33d66efa1f5",
            "name": "Top Devices",
            "description": "Top devices",
            "type": "CURATED",
            "rules": [
                {
                    "attribute": "ARN",
                    "operator": "IN",
                    "value": "[\"arn:aws:devicefarm:us-west-2::device:5F1B162C265B4F34804B7D0DC2CDBE40\",\"arn:aws:devicefarm:us-west-2::device:1AEC6BFEDFA943299F801CB38A289E2E\",\"arn:aws:devicefarm:us-west-2::device:DEC41B3B48534980BC5B0432E2D34CA7\",\"arn:aws:devicefarm:us-west-2::device:4F74D943F7594EFF96957E238B3CA131\",\"arn:aws:devicefarm:us-west-2::device:9C236F5665C1466BB717C2517ACC3FE2\",\"arn:aws:devicefarm:us-west-2::device:E64D26FE27644A39A4BCEF009CDD8645\",\"arn:aws:devicefarm:us-west-2::device:2832D5722BEF4FF2B04498ECC4C1C2F6\",\"arn:aws:devicefarm:us-west-2::device:CEA80E8918814308A6275FEBC4310134\",\"arn:aws:devicefarm:us-west-2::device:5F20BBED05F74D6288D51236B0FB9895\",\"arn:aws:devicefarm:us-west-2::device:52C50B02A5154CC0AF653512445DB7B6\"]"
                }
            ]
        },
		%s
    ]
}`, devicePoolJSONExample)

var listDevicePoolsExample = []DevicePool{
	NewDevicePool(
		"arn:aws:devicefarm:us-west-2::devicepool:082d10e5-d7d7-48a5-ba5c-b33d66efa1f5",
		"Top Devices",
		"Top devices",
		[]DevicePoolRule{
			{
				Attribute: "ARN",
				Operator:  "IN",
				Value:     `["arn:aws:devicefarm:us-west-2::device:5F1B162C265B4F34804B7D0DC2CDBE40","arn:aws:devicefarm:us-west-2::device:1AEC6BFEDFA943299F801CB38A289E2E","arn:aws:devicefarm:us-west-2::device:DEC41B3B48534980BC5B0432E2D34CA7","arn:aws:devicefarm:us-west-2::device:4F74D943F7594EFF96957E238B3CA131","arn:aws:devicefarm:us-west-2::device:9C236F5665C1466BB717C2517ACC3FE2","arn:aws:devicefarm:us-west-2::device:E64D26FE27644A39A4BCEF009CDD8645","arn:aws:devicefarm:us-west-2::device:2832D5722BEF4FF2B04498ECC4C1C2F6","arn:aws:devicefarm:us-west-2::device:CEA80E8918814308A6275FEBC4310134","arn:aws:devicefarm:us-west-2::device:5F20BBED05F74D6288D51236B0FB9895","arn:aws:devicefarm:us-west-2::device:52C50B02A5154CC0AF653512445DB7B6"]`,
			},
		},
	),
	devicePoolExample,
}

var getDevicePoolExampleJSON = fmt.Sprintf(`{ "devicePool": %s }`, devicePoolJSONExample)
