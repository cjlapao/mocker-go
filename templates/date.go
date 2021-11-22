package templates

import (
	"time"

	"github.com/cjlapao/mocker-go/mocker"
)

type Date struct {
}

func (d Date) Between(initialDate string, finalDate string) string {
	startDate, err := time.Parse(time.RFC3339, initialDate)
	if err != nil {
		startDate = time.Now()
	}
	endDate, err := time.Parse(time.RFC3339, finalDate)
	if err != nil {
		endDate = time.Now()
	}

	return mocker.New().Date().Between(startDate, endDate).Format(time.RFC3339)
}

func (d Date) Past(years int) string {
	return mocker.New().Date().Past(years).Format(time.RFC3339)
}

func (d Date) Recent(days int) string {
	return mocker.New().Date().Recent(days).Format(time.RFC3339)
}

func (d Date) Soon(days int) string {
	return mocker.New().Date().Soon(days).Format(time.RFC3339)
}

func (d Date) Future(years int) string {
	return mocker.New().Date().Future(years).Format(time.RFC3339)
}

func (d Date) Month() string {
	return mocker.New().Date().Month()
}

func (d Date) Weekday() string {
	return mocker.New().Date().Weekday()
}
