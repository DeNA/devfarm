package planfile

func StubValidatorFunc(planfile Planfile, err error) ValidateFunc {
	return func(UnsafePlanFile) (Planfile, error) {
		return planfile, err
	}
}
