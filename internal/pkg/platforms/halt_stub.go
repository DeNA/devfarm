package platforms

func StubHalt(results Results) Halt {
	return func(InstanceGroupName, HaltBag) (Results, error) {
		return results, results.Err()
	}
}

func AnyHaltBag() HaltBag {
	return AnyBag()
}
