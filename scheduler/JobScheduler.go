package scheduler

import (
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

type JobScheduler struct {
	scheduler  *cron.Cron
	cronString string
}

func CreateJobScheduler(interval, offset time.Duration) *JobScheduler {
	parser := cron.NewParser(cron.Second | cron.Minute | cron.Hour)
	cronScheduler := cron.New(cron.WithParser(parser))
	cronString, err := createCronString(interval, offset)
	if err != nil {
		return nil
	}
	return &JobScheduler{scheduler: cronScheduler, cronString: cronString}
}

func (jobScheduler *JobScheduler) AddFunction(cmd func()) (cron.EntryID, error) {
	return jobScheduler.scheduler.AddFunc(jobScheduler.cronString, cmd)
}

func (jobScheduler *JobScheduler) Start() {
	jobScheduler.scheduler.Start()
}

func createCronString(interval, offset time.Duration) (string, error) {
	if offset >= interval {
		message := fmt.Sprintf("interval should be greater than offset, interval: %d, offset: %d", interval, offset)
		return "", errors.New(message)
	}
	if interval < time.Minute {
		return getSecondlyIntervalCronString(interval, offset), nil
	}
	if interval < time.Hour {
		return getMinutelyIntervalCronString(interval, offset), nil
	}
	return getHourlyIntervalCronString(interval, offset), nil
}

func getSecondlyIntervalCronString(interval, offset time.Duration) string {
	secondInterval := int(interval / time.Second)
	secondOffset := int(offset / time.Second)
	return fmt.Sprintf("%d/%d * *", secondOffset, secondInterval)
}

func getMinutelyIntervalCronString(interval, offset time.Duration) string {
	minuteInterval := int(interval / time.Minute)
	if offset < time.Minute {
		secondOffset := int(offset / time.Second)
		return fmt.Sprintf("%d */%d *", secondOffset, minuteInterval)
	}
	minuteOffset := int(offset / time.Minute)
	return fmt.Sprintf("0 %d/%d *", minuteOffset, minuteInterval)
}

func getHourlyIntervalCronString(interval, offset time.Duration) string {
	hourInterval := int(interval / time.Hour)
	if offset < time.Minute {
		secondOffset := int(offset / time.Second)
		return fmt.Sprintf("%d 0 */%d", secondOffset, hourInterval)
	}
	if offset < time.Hour {
		minuteOffset := int(offset / time.Minute)
		return fmt.Sprintf("0 %d */%d", minuteOffset, hourInterval)
	}
	hourOffset := int(offset / time.Hour)
	return fmt.Sprintf("0 0 %d/%d", hourOffset, hourInterval)
}
