package jobs

import (
	"bitbucket.org/kenorld/eject-core"
	"bitbucket.org/kenorld/eject-cron"
)

const DefaultJobPoolSize = 10

var (
	// Singleton instance of the underlying job scheduler.
	MainCron *cron.Cron

	// This limits the number of jobs allowed to run concurrently.
	workPermits chan struct{}

	// Is a single job allowed to run concurrently with itself?
	selfConcurrent bool
)

func init() {
	MainCron = cron.New()
	eject.OnAppStart(func() {
		if size := eject.Config.GetIntDefault("jobs.pool", DefaultJobPoolSize); size > 0 {
			workPermits = make(chan struct{}, size)
		}
		selfConcurrent = eject.Config.GetBoolDefault("jobs.self_concurrent", false)
		MainCron.Start()
	})
}
