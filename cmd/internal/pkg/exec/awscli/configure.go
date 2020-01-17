package awscli

type ConfigStatusGetter func() error

func NewConfigStatusGetter(awsCmd Executor) ConfigStatusGetter {
	return func() error {
		_, err := awsCmd("configure", "get", "aws_access_key_id")
		if err != nil {
			return err
		}

		return nil
	}
}
