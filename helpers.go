package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
)

// Respond Function Is To Handel Http Returns With A Json.
func Respond(c echo.Context, payload map[string]interface{}, status int, success bool) error {
	return c.JSON(status, map[string]interface{}{
		"status":  status,
		"success": success,
		"payload": payload,
	})
}

// RespondValidationErrors Function Is To Handel Http Returns With A Json.
func RespondValidationErrors(c echo.Context, errors map[string][]string) error {
	return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
		"status":  http.StatusUnprocessableEntity,
		"success": false,
		"payload": map[string]string{
			"message": "Validation Errors Please Check Your Inputs",
		},
		"errors": errors,
	})
}

// RangeDate returns a date range function over start date to end date inclusive.
// After the end of the range, the range function returns a zero date,
// date.IsZero() is true.
func RangeDate(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}

// InIntSclice ...
func InIntSclice(val int, array []int) bool {
	for _, v := range array {
		if v == val {
			return true
		}
	}
	return false
}

// RegisterCustomValidationRules ...
func RegisterCustomValidationRules() {
	govalidator.AddCustomRule("weekdays", func(field string, rule string, message string, value interface{}) error {
		days := strings.Split(value.(string), ",")

		if len(days) == 0 {
			return fmt.Errorf("The %s weekdays invalid expected values between 0-6. EX: 0,1,2", field)
		}

		for _, day := range days {
			d, _ := strconv.Atoi(day)
			if d < 0 || d > 6 {
				return fmt.Errorf("The %s weekday number is invalid expected values between 0-6. EX: 0,1,2", field)
			}
		}

		return nil
	})
}
