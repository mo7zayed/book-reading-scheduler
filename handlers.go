package main

import (
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
)

func init() {
	RegisterCustomValidationRules()
}

func getWeekdays(daysString string) []int {
	var weekdays []int

	days := strings.Split(daysString, ",")

	for _, day := range days {
		d, _ := strconv.Atoi(day)
		weekdays = append(weekdays, d)
	}

	return weekdays
}

func getDates(StartDate string, addMonthes int) map[string]time.Weekday {
	dates := make(map[string]time.Weekday)

	start, _ := time.Parse(DateFormat, StartDate)
	end := start.AddDate(0, addMonthes, 0)

	for rd := RangeDate(start, end); ; {
		date := rd()
		if date.IsZero() {
			break
		}

		dates[date.Format(DateFormat)] = date.Weekday()
	}

	return dates
}

// HandleHome handles home page request
func HandleHome(c echo.Context) error {
	// Validate the incomming request
	validationErrors := govalidator.New(govalidator.Options{
		Request: c.Request(),
		Rules: govalidator.MapData{
			"start_date":      []string{"date", "date:yyyy-mm-dd"},
			"days":            []string{"weekdays"},
			"sessions_number": []string{"min:1", "max:100"},
		},
		RequiredDefault: true, // all the field to be pass the rules
	}).Validate()

	if len(validationErrors) > 0 {
		return RespondValidationErrors(c, validationErrors)
	}

	payload := new(Payload)
	if err := c.Bind(payload); err != nil {
		return Respond(c, map[string]interface{}{
			"message": "Error processing payload",
		}, http.StatusUnprocessableEntity, false)
	}

	// do logic
	weekDays := getWeekdays(payload.Days)

	neededWeeksForAChapter := float64(payload.SessionsNumber) / float64(len(weekDays))

	totalWeeksNeeded := int(math.Round(float64(neededWeeksForAChapter * float64(ChaptersNumber))))

	dates := getDates(
		payload.StartDate,
		totalWeeksNeeded/4,
	)

	var neededDays []string

	for date, wd := range dates {
		if InIntSclice(int(wd), weekDays) {
			neededDays = append(neededDays, date)
		}
	}

	return Respond(c, map[string]interface{}{
		"dates": neededDays,
	}, http.StatusOK, true)
}
