package devicefarm

import "encoding/json"

type TestType string

func (t TestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

const (
	// BUILTIN_FUZZ: The built-in fuzz type.
	// BUILTIN_EXPLORER: For Android, an app explorer that will traverse an Android app, interacting with it and capturing screenshots at the same time.
	// APPIUM_JAVA_JUNIT: The Appium Java JUnit type.
	// APPIUM_JAVA_TESTNG: The Appium Java TestNG type.
	// APPIUM_PYTHON: The Appium Python type.
	// APPIUM_NODE: The Appium Node.js type.
	// APPIUM_RUBY: The Appium Ruby type.
	// APPIUM_WEB_JAVA_JUNIT: The Appium Java JUnit type for web apps.
	// APPIUM_WEB_JAVA_TESTNG: The Appium Java TestNG type for web apps.
	// APPIUM_WEB_PYTHON: The Appium Python type for web apps.
	// APPIUM_WEB_NODE: The Appium Node.js type for web apps.
	// APPIUM_WEB_RUBY: The Appium Ruby type for web apps.
	// CALABASH: The Calabash type.
	// INSTRUMENTATION: The Instrumentation type.
	// UIAUTOMATION: The uiautomation type.
	// UIAUTOMATOR: The uiautomator type.
	// XCTEST: The Xcode test type.
	// XCTEST_UI: The Xcode UI test type.
	TestTypeIsAppiumNode TestType = "APPIUM_NODE"
)

type TestProp struct {
	TestType       TestType  `json:"type"`
	TestPackageARN UploadARN `json:"testPackageArn"`
	TestSpecARN    UploadARN `json:"testSpecArn"`
}

func NewTestProp(testType TestType, testPackageARN UploadARN, testSpecARN UploadARN) TestProp {
	return TestProp{
		TestType:       testType,
		TestPackageARN: testPackageARN,
		TestSpecARN:    testSpecARN,
	}
}
