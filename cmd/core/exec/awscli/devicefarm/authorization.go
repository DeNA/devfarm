package devicefarm

type AuthorizationStatusChecker func() error

func NewAuthorizationStatusChecker(deviceFarmCmd Executor) AuthorizationStatusChecker {
	return func() error {
		_, err := deviceFarmCmd("list-devices")
		return err
	}
}
