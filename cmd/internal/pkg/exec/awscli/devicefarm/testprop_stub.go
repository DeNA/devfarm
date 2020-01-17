package devicefarm

func AnyTestProp() TestProp {
	return NewTestProp(
		TestTypeIsAppiumNode,
		"arn:aws:devicefarm:ANY_TEST_PACKAGE_UPLOAD",
		"arn:aws:devicefarm:ANY_TEST_SPEC_UPLOAD",
	)
}
