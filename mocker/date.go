package mocker

import (
	"math"
	"time"
)

var _dateGenerator *DateGenerator

type DateGenerator struct {
	Mocker *Mocker
}

func NewDateGenerator(mocker *Mocker) *DateGenerator {
	if _dateGenerator != nil {
		return _dateGenerator
	}

	_dateGenerator = &DateGenerator{}
	_dateGenerator.Mocker = mocker

	return _dateGenerator
}

func (d *DateGenerator) Between(initialDate time.Time, finalDate time.Time) time.Time {
	if finalDate.After(initialDate) {
		randomYear := d.Mocker.IntBetween(initialDate.Year(), finalDate.Year())
		randomMonth := d.Mocker.IntBetween(int(initialDate.Month()), int(finalDate.Month()))
		randomDay := d.Mocker.IntBetween(int(initialDate.Day()), int(finalDate.Day()))
		randomHour := d.Mocker.IntBetween(int(initialDate.Hour()), int(finalDate.Hour()))
		randomMinute := d.Mocker.IntBetween(int(initialDate.Minute()), int(finalDate.Minute()))
		randomSecond := d.Mocker.IntBetween(int(initialDate.Second()), int(finalDate.Second()))
		randomNanosecond := d.Mocker.IntBetween(int(initialDate.Nanosecond()), int(finalDate.Nanosecond()))
		result := time.Date(randomYear, time.Month(randomMonth), randomDay, randomHour, randomMinute, randomSecond, randomNanosecond, time.UTC)
		println("Year: " + result.Format(time.RFC3339))
		return result
	}

	return time.Now()
}

func (d *DateGenerator) Past(years int) time.Time {
	initialYear := time.Now().Year() - 1
	initialDate := time.Date(initialYear, 1, 1, 0, 0, 0, 0, time.UTC)
	finalYear := time.Now().Year() - (years + 1)
	finalDate := time.Date(finalYear, 12, 31, 23, 59, 59, 999, time.UTC)

	result := d.Between(finalDate, initialDate)

	return result
}

func (d *DateGenerator) Recent(days int) time.Time {
	if days <= 0 {
		return time.Now()
	}

	randomDay := d.Mocker.IntBetween(0, days)
	duration := time.Duration(randomDay * 24 * int(time.Hour))
	numberHours := math.Abs(float64(duration)) * -1

	finalDate := time.Now().Add(time.Duration(numberHours))

	return finalDate
}

func (d *DateGenerator) Soon(days int) time.Time {
	if days <= 0 {
		return time.Now()
	}

	randomDay := d.Mocker.IntBetween(0, days)
	duration := time.Duration(randomDay * 24 * int(time.Hour))

	finalDate := time.Now().Add(duration)

	return finalDate
}

func (d *DateGenerator) Future(years int) time.Time {
	initialYear := time.Now().Year() + 1
	initialDate := time.Date(initialYear, 1, 1, 0, 0, 0, 0, time.UTC)
	finalYear := time.Now().Year() + (years + 1)
	finalDate := time.Date(finalYear, 12, 31, 23, 59, 59, 999, time.UTC)

	result := d.Between(initialDate, finalDate)

	return result
}

func (d *DateGenerator) Month() string {
	randomMonthInt := d.Mocker.IntBetween(1, 12)

	month := time.Month(randomMonthInt)
	return month.String()
}

func (d *DateGenerator) Weekday() string {
	randomWeekday := d.Mocker.IntBetween(0, 6)

	println(time.Weekday(0).String())
	weekDay := time.Weekday(randomWeekday)
	return weekDay.String()
}
