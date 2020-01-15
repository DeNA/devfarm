package all

import (
	"github.com/dena/devfarm/internal/pkg/platforms"
	"sync"
)

func (ps Platforms) RunIOS(iosPlan platforms.IOSPlan) error {
	p, err := ps.GetPlatform(iosPlan.Platform)
	if err != nil {
		return err
	}

	runIOS := p.IOSRunner()
	return runIOS(iosPlan)
}

func (ps Platforms) RunAndroid(iosPlan platforms.AndroidPlan) error {
	p, err := ps.GetPlatform(iosPlan.Platform)
	if err != nil {
		return err
	}

	runAndroid := p.AndroidRunner()
	return runAndroid(iosPlan)
}

func (ps Platforms) RunAll(plans []platforms.EitherPlan) (ResultTable, error) {
	builder := NewResultTableBuilder()
	var wg sync.WaitGroup

	for platformID, plansForPlatform := range groupByPlatform(plans) {
		p, err := ps.GetPlatform(platformID)
		if err != nil {
			builder.AddErrors(platformID, err)
			continue
		}

		wg.Add(1)
		go func(p platforms.Platform, plansForPlatform []platforms.EitherPlan) {
			run := p.Runner()

			results, _ := run(plansForPlatform)
			builder.AddErrors(p.ID(), results...)

			wg.Done()
		}(p, plansForPlatform)
	}
	wg.Wait()

	table := builder.Build()
	return table, table.Err()
}
