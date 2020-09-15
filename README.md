devfarm
=======

[![CircleCI](https://circleci.com/gh/DeNA/devfarm/tree/master.svg?style=svg)](https://circleci.com/gh/DeNA/devfarm/tree/master)

Tools to control iOS and Android mobile apps across several device farms such as AWS Device Farm.  The purpose of this project is portability among major device farms (including emulator/simulator farms) and be able to do Domain-Specific Tests across device farms.

The portability is important because there are a lot of device farms, so you must try these device farms indivisually to verify until satisfy your usecases.  This challenge is hard work but you can easily switch to other device cloud if you using this tool.

This tool supports only launching/watching/halting mobile apps on device farms.  These features can help domain-specific tests that sending/receiving directly domain events (e.g. jump a character, attack to someone, reading other character status, ...). Such tests can be noise-free and more efficient than UI-layer fuzzing or UI-layer schenario tests using like Appium or Calabash.



Usecases
--------

In our usecase, we bundled a domain specific random-walk agent into our apps.  This random-walk agent crash if something go wrong.  This tool could watch app crashes on device farms via CLI, so we could know it via CI services easily.  Additionaly, we could improve the coverage per unit time of the random-walk agent by increasing parallel devices numbers.

Or, if you want to do traditional tests using like JUnit or XCTest for domain specific-tests, you still can bundle the test harness into your apps.



Usage
-----

```console
$ devfarm run-ios --os-version 12.0 --device 'apple iphone xs' --instance-group example --platform aws-device-farm --ipa path/to/app.ipa --args-json '["-ARG1", "HELLO_DEVFARM"]' --lifetime-sec 900 &
platform                status
aws-device-farm         launching

$ devfarm status --instance-group example
platform                device                  os      state   note
aws-device-farm         apple iphone xs         ios     ACTIVE

$ devfarm halt --instance-group example
platform                status
aws-device-farm         halting

$ devfarm status --instance-group example
platform                device                  os      state           note
aws-device-farm         apple iphone xs         ios     INACTIVATING
```


### List Available Device and OS

```console
$ devfarm list-devices
platform                os              device                                  available       note
aws-device-farm         ios 9           apple iphone xs                         yes
...
aws-device-farm         android 9       google google pixel 3                    yes
...
```


### Launch by plan.yml

```console
$ devfarm run-all --help
Usage: [options] <plan.yml>
  -dry-run
        enables dry-run (WARNING: not stable yet)
  -verbose
        enables verbose logs
```

Syntax of plan.yml:

```yaml
instance_groups:
  <group>:
    # for iOS
    - platform: <platform>  # required
      ios: <version>        # required
      device: <device>      # required
      ipa: <filepath>       # required
      args: []              # optional
      lifetime_sec:         # required

    # for Android
    - platform: <platform>  # required
      android: <version>    # required
      device: <decice>      # required
      apk: <filepath>       # required
      app_id: <app_id>      # required
      intent_extras: []     # optional
      lifetime_sec:         # required
```

<details>
<summary>Example</summary>

```console
$ cat path/to/plan.yml
instance_groups:
  example:
    - platform: aws-device-farm
      ios: 12.0
      device: apple iphone xs
      ipa: path/to/app.ipa
      args:
        - -ARG1
        - VALUE1
      lifetime_sec: 900
    - platform: aws-device-farm
      android: 9
      device: google google pixel 3
      apk: path/to/app.apk
      app_id: com.example.app
      intent_extras:
        - -e
        - ARG1
        - VALUE1
      lifetime_sec: 900

$ devfarm run-all --dry-run path/to/plan.yml
platform         status     note
aws-device-farm  launching
aws-device-farm  launching
```
</details>


### Instance Group Dependent Commands

Instance Groups is an unit to halt or check status.
You can specify names what you want but it must be a non-empty and must not include `[^0-9a-zA-Z_]`.



#### Launch iOS Apps

```console
$ devfarm run-ios --help
Usage:
  -args-json string
        arguments that will be passed to the iOS app (via ProcessInfo#arguments) after decoding to plain arguments (default "[]")
  -device string
        device name listed by 'devfarm list-devices' (required)
  -dry-run
        enables dry-run (WARNING: not stable yet)
  -instance-group string
        instance group name (required)
  -ipa string
        ipa file to launch (required)
  -os-version string
        iOS version listed by 'devfarm list-devices' (required)
  -platform string
        platform name listed by 'devfarm list-devices' (required)
  -verbose
        enables verbose logs
```

<details>
<summary>Example</summary>

```console
$ devfarm run-ios \
    --instance-group example \
    --platform aws-device-farm \
    --device 'apple iphone xs' \
    --os-version 12.0 \
    --ipa 'path/to/app.ipa' \
    --args-json '["-ARG", "VALUE"]'
platform                status
aws-device-farm         launching
```
</details>



#### Launch Android Apps

```console
$ devfarm run-android --help
Usage:
  -apk string
        apk file to launch (required)
  -app-id string
        application ID (it often called as 'package name') to the app (required)
  -device string
        device name listed by 'devfarm list-devices' (required)
  -dry-run
        enables dry-run (WARNING: not stable yet)
  -instance-group string
        instance group name (required)
  -intent-extras-json string
        arguments that will be passed to the Android app (via Intent#getExtras) after decoding to plain arguments (default "[]")
  -os-version string
        Android version listed by 'devfarm list-devices' (required)
  -platform string
        platform name listed by 'devfarm list-devices' (required)
  -verbose
        enables verbose logs
```

<details>
<summary>Example</summary>

```console
$ devfarm forever-android \
    --instance-group example \
    --platform aws-device-farm \
    --device 'google google pixel 3' \
    --os-version 9 \
    --apk 'path/to/app.apk' \
    --app-id 'com.example.app' \
    --intent-extras-json '["-e", "ARG", "VALUE"]'
platform                status
aws-device-farm         launching
```
</details>


#### Status

```console
$ devfarm status --help
Usage:
  -dry-run
        enables dry-run (WARNING: not stable yet)
  -instance-group string
        instance group name to filter (optional)
  -verbose
        enables verbose logs
```

<details>
<summary>Example</summary>

```console
$ devfarm status
platform                device                  os      state   note
aws-device-farm         apple iphone xs         ios     ACTIVE
```

or

```
$ devfarm status --instance-group example
platform                device                  os      state   note
aws-device-farm         apple iphone xs         ios     ACTIVE
```
</details>


#### Stop Apps

```console
$ devfarm halt --help
Usage:
  -dry-run
        enables dry-run (WARNING: not stable yet)
  -instance-group string
        instance group name (required)
  -verbose
        enables verbose logs
```

<details>
<summary>Example</summary>

```console
$ devfarm halt --instance-group example
platform                status
aws-device-farm         halting
```
</details>


### Utility Commands
<details>
<summary>Commands</summary>

#### Version

```console
$ devfarm version
0.0.0
```


#### Check Authentication Status

```console
$ devfarm auth-status --help
Usage:
  -dry-run
        enables dry-run (WARNING: not stable yet)
  -verbose
        enables verbose logs
```

<details>
<summary>Example</summary>

```console
$ devfarm auth-status
platform                auth
aws-device-farm         success
```
</details>


#### Validate plan.yml

```console
$ devfarm validate --help
Usage: <plan.yml>
  -verbose
        enables verbose logs
```

<details>
<summary>Example</summary>

```console
$ cat path/to/plan.yml
instance_groups:
  example:
    - platform: any-platform
      ios: 12.0
      device: apple iphone xs
      ipa: path/to/app.ipa
      args:
        - -ARG1
        - VALUE1
    - platform: any-platform
      android: 9
      device: google google pixel 3
      apk: path/to/app.apk
      app_id: com.example.app
      intent_extras:
        - -e
        - ARG1
        - VALUE1

$ devfarm validate path/to/plan.yml
$ echo $?
0

$ devfarm validate --verbose path/to/plan.yml
{
  "instance_groups": {
    "example": [
      {
        "platform": "any-platform",
        "os": "ios",
        "ios": {
          "group_name": "example",
          "device": {
            "name": "apple iphone xs",
            "ios_version": "12.0"
          },
          "ipa": "path/to/app.ipa",
          "args": [
            "-ARG1",
            "VALUE1"
          ]
        },
        "android": {
          "group_name": "",
          "device": {
            "name": "",
            "android_version": ""
          },
          "apk": "",
          "app_id": "",
          "intent_extras": null
        }
      },
      {
        "platform": "any-platform",
        "os": "android",
        "ios": {
          "group_name": "",
          "device": {
            "name": "",
            "ios_version": ""
          },
          "ipa": "",
          "args": null
        },
        "android": {
          "group_name": "example",
          "device": {
            "name": "google google pixel 3",
            "android_version": "9"
          },
          "apk": "path/to/app.apk",
          "app_id": "com.example.app",
          "intent_extras": [
            "-e",
            "ARG1",
            "VALUE1"
          ]
        }
      }
    ]
  }
}

$ devfarm validate path/to/broken.yml
invalid iOS plan (at 1-th plan of instance group "example"):
    device: must not be empty
    ipa: must not be empty
invalid plan (at 2-th plan of instance group "example"):
    unsupported os: "unavailable"
invalid Android plan (at 1-th plan of instance group "other"):
    platform: must not be empty
    device: must not be empty
    apk: must not be empty
    app_id: must not be empty

$ echo $?
1
```
</details>


#### See Bundled Assets
```
$ devfarm ls-assets
```

<details>
<summary>Example</summary>

```console
$ devfarm ls-assets
assets/aws-device-farm/workflows/0-shared.bash
assets/aws-device-farm/workflows/1-install.bash
assets/aws-device-farm/workflows/2-pretest.bash
assets/aws-device-farm/workflows/3-test.bash
assets/aws-device-farm/workflows/4-posttest.bash
assets/devfarmagent/darwin-amd64/devfarmagent
assets/devfarmagent/devfarmagent.bash
assets/devfarmagent/linux-amd64/devfarmagent
assets/ios-deploy-agent/package-lock.json
assets/ios-deploy-agent/package.json
```
</details>


#### See Bundled Asset Content
```console
$ devfarm cat-asset <asset>
```

<details>
<summary>Example</summary>

```console
$ devfarm cat-asset assets/devfarmagent/darwin-amd64/devfarmagent | file -
/dev/stdin: Mach-O 64-bit executable x86_64
```
</details>
</details>



Install
-------

```console
$ go get -u github.com/dena/devfarm/cmd/devfarm
```

<details>
<summary>Using Docker</summary>

```console
$ docker pull <image_id>

$ cat .env
AWS_ACCESS_KEY_ID=***
AWS_SECRET_ACCESS_KEY=***

$ docker run --rm --env-file ./.env <image_id> auth-status

$ tree -a "$(pwd)"
.
├── .env
└── app
    ├── Example.apk
    ├── Example.ipa
    └── planfile.yml

1 directory, 4 files

$ docker run --rm --env-file ./.env -v "$(pwd)/app:/app" <image_id> run-all /app/planfile
```
</details>



Requirements
------------

- To use AWS Device Farm
    - [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-welcome.html)


Todo
----

- ~~Stabilize iOS apps made by Unity on AWS Device Farm (at now, 2/3 runs were failed)~~
    - We have started to do external monitoring, and we confirmed that it almost achived, but still about 1/30 runs were unstable
- Detect that the app become background
    - Detection for iOS have never been tested
    - Detection for Android is not supported yet, because the version of adb on AWS Device Farm is 1.0.32 (too old)
- Work `--dry-run` on all commands for testing
- Support Android Simulator on AWS EC2
