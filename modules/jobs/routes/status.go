package controllers

import (
	"strings"

	"github.com/eject/cron"
	"github.com/eject/modules/jobs/app/jobs"
	"github.com/eject/eject"
)

type Jobs struct {
	*eject.Controller
}

func (c Jobs) Status() eject.Result {
	remoteAddress := c.Request.RemoteAddr
	if eject.Config.BoolDefault("jobs.acceptproxyaddress", false) {
		if proxiedAddress, isProxied := c.Request.Header["X-Forwarded-For"]; isProxied {
			remoteAddress = proxiedAddress[0]
		}
	}
	if !strings.HasPrefix(remoteAddress, "127.0.0.1") && !strings.HasPrefix(remoteAddress, "::1") {
		return c.Forbidden("%s is not local", remoteAddress)
	}
	entries := jobs.MainCron.Entries()
	return c.Render(entries)
}

func init() {
	eject.TemplateFuncs["castjob"] = func(job cron.Job) *jobs.Job {
		return job.(*jobs.Job)
	}
}
