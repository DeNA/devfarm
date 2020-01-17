package all

import "github.com/dena/devfarm/cmd/internal/pkg/platforms"

func groupByPlatform(plans []platforms.EitherPlan) map[platforms.ID][]platforms.EitherPlan {
	group := make(map[platforms.ID][]platforms.EitherPlan)

	for _, plan := range plans {
		if groupedPlans, ok := group[plan.CommonPart.Platform]; ok {
			group[plan.CommonPart.Platform] = append(groupedPlans, plan)
		} else {
			group[plan.CommonPart.Platform] = []platforms.EitherPlan{plan}
		}
	}

	return group
}
