package devicefarm

func AnyExecutionConfiguration() ExecutionConfiguration {
	return NewExecutionConfiguration(
		JobTimeout(0),
		false,
		false,
		false,
		false,
	)
}
