package platforms

type PlanValidatorBag interface{}

type PlanValidator func(bag PlanValidatorBag, plan EitherPlan) error
