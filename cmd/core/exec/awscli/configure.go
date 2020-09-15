package awscli

import "github.com/dena/devfarm/cmd/core/exec"

type ConfigStatusGetter func() error

func NewConfigStatusGetter(awsCmd Executor, env exec.EnvGetter) ConfigStatusGetter {
	return func() error {
		// SEE: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html#envvars-list
		if env("AWS_ACCESS_KEY_ID") != "" && env("AWS_SECRET_ACCESS_KEY") != "" {
			return nil
		}

		_, err := awsCmd("configure", "get", "aws_access_key_id")
		if err != nil {
			return err
		}

		return nil
	}
}
