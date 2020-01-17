Executor
=======



What Components Should be an Executor
-------------------------------------

Executor is a component that encapsulates communications to out of the process or package.
For example, following components should be an Executor:

- communicating to any SDK
- communicating to any HTTP server
- communicating to any external commands

Executors should be used by platform depended workflows in `./internal/pkg/platform/...`.



Implementation Pattern
----------------------

Executors should be implemented the pattern describing below to make reusable and testable.

```go
// Defines an Executor Type. This type should be a function type.
// The argument such as ArgN should be actual value the Executor needs, and should not be an other Executor.
type MyExecutor func(Arg1, Arg2) (Ret1, Ret2)

// Defines an actual Executor implementation.
// The argument such as InjectedN may be an Executor that excapsulates complex workflow.
func NewMyExecutor(Injected1, Injected2) MyExecutor {
	return func(arg1 Arg1, arg2 Arg2) (Ret1, Ret2) {
		// ...
	}
}
```

This pattern is highly reusable because Executors is a composition of small Executors.

And it can be easily tested like following:

```go
func TestNewMyExecutor(t *testing.T) {
	// Prepares dependencies.
	injected1 := AnyInjected1()
	injected1.Something = "..."
	injected2 := AnyInjected2()
	injected2.Another = "..."

	// Creates the exec under test.
	myExecutor := NewMyExecutor(injected1, injected2)

	// Exercise the exec with arguments.
	ret1, ret2 := myExecutor("ARG1", "ARG2")

	if ret1 != "..." {
		t.Errorf("got %v", ret1)
		return
	}

	// ...
}
```
