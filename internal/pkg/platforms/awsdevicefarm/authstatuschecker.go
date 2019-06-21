package awsdevicefarm

import (
	"github.com/dena/devfarm/internal/pkg/executor/awscli"
	"github.com/dena/devfarm/internal/pkg/executor/awscli/devicefarm"
)

type authStatusChecker func() error

func newAuthStatusChecker(
	isInstalled awscli.InstalltionStatusGetter,
	version awscli.VersionGetter,
	isConfigured awscli.ConfigStatusGetter,
	checkAuthorization devicefarm.AuthorizationStatusChecker,
) authStatusChecker {
	return func() error {
		if err := isInstalled(); err != nil {
			return &AuthStatusReason{
				message:   "AWS CLI is not installed. please install it via https://aws.amazon.com/cli/",
				debugInfo: err.Error(),
			}
		}
		if _, err := version(); err != nil {
			return &AuthStatusReason{
				message:   "cannot execute aws command (--verbose may help you)",
				debugInfo: err.Error(),
			}
		}
		if err := isConfigured(); err != nil {
			return &AuthStatusReason{
				message:   "aws cli is not configured (--verbose may help you)",
				debugInfo: err.Error(),
			}
		}
		if err := checkAuthorization(); err != nil {
			return &AuthStatusReason{
				message:   "failure response was returned (--verbose may help you)",
				debugInfo: err.Error(),
			}
		}
		return nil
	}
}

type AuthStatusReason struct {
	message   string
	debugInfo string
}

func (e *AuthStatusReason) Error() string {
	return e.message
}
