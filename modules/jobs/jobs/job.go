package jobs

import (
	"reflect"
	"runtime/debug"
	"sync"
	"sync/atomic"

	"bitbucket.org/kenorld/eject-corn"
	"bitbucket.org/kenorld/eject-core"
	"github.com/Sirupsen/logrus"
)

type Job struct {
	Name    string
	inner   cron.Job
	status  uint32
	running sync.Mutex
}

const UNNAMED = "(unnamed)"

func New(job cron.Job) *Job {
	name := reflect.TypeOf(job).Name()
	if name == "Func" {
		name = UNNAMED
	}
	return &Job{
		Name:  name,
		inner: job,
	}
}

func (j *Job) Status() string {
	if atomic.LoadUint32(&j.status) > 0 {
		return "RUNNING"
	}
	return "IDLE"
}

func (j *Job) Run() {
	// If the job panics, just print a stack trace.
	// Don't let the whole process die.
	defer func() {
		if err := recover(); err != nil {
			if ejectError := eject.NewErrorFromPanic(err); ejectError != nil {
				logrus.WithFields(logrus.Fields{
				"error": err,
				"stack": ejectError.Stack
			}).Error("error in job")
			} else {
				logrus.WithFields(logrus.Fields{
				"error": err,
				"stack": string(debug.Stack())
			}).Error("error in job")
			}
		}
	}()

	if !selfConcurrent {
		j.running.Lock()
		defer j.running.Unlock()
	}

	if workPermits != nil {
		workPermits <- struct{}{}
		defer func() { <-workPermits }()
	}

	atomic.StoreUint32(&j.status, 1)
	defer atomic.StoreUint32(&j.status, 0)

	j.inner.Run()
}
