// Package sqlquery handles parsing SQL and converting to a dialect for Tidalwave.
package sqlquery

import (
	"strings"

	"github.com/busbud/tidalwave/logger"
	"github.com/dustinblackman/moment"
	"github.com/jinzhu/copier"
)

const queryDateFormat = "YYYY-MM-DDTHH:mm:ss"

// DateParam stores date query information.
type DateParam struct {
	Date     string
	DateTime *moment.Moment
	Operator string
	TimeUsed bool
	Type     string
}

func createDateParam(date, operator string) []DateParam {
	dateParam := DateParam{Operator: operator, TimeUsed: true}
	date = stripQuotes(date)
	if len(date) > 0 && len(date) <= 10 {
		dateParam.TimeUsed = false
		if operator == "<=" {
			date += "T23:59:59"
		} else {
			date += "T00:00:00"
		}
	}
	dateParam.Date = date
	dateParam.Type = "start"
	if operator == "<" || operator == "<=" {
		dateParam.Type = "end"
	}

	dateParam.DateTime = moment.New().Moment(queryDateFormat, date)

	if operator == "=" && !dateParam.TimeUsed {
		returnDateParam := DateParam{}
		err := copier.Copy(&returnDateParam, &dateParam)
		if err != nil {
			logger.Log.Fatal(err)
		}

		dateParam.Operator = ">="
		returnDateParam.Operator = "<="
		returnDateParam.Type = "end"

		endDate := strings.Split(date, "T")[0] + "T23:59:59"
		returnDateParam.Date = endDate
		returnDateParam.DateTime = moment.New().Moment(queryDateFormat, endDate)

		return []DateParam{dateParam, returnDateParam}
	}

	return []DateParam{dateParam}
}
