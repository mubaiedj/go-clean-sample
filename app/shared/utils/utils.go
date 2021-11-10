package utils

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"strconv"
	"time"
)

func ConvertEntity(in, out interface{}) interface{} {
	str, _ := json.Marshal(in)
	err2 := json.Unmarshal(str, out)

	if err2 != nil {
		return nil
	}
	return out
}

func EntityToJson(entity interface{}) string {
	str, err := json.Marshal(entity)
	if err != nil {
		return "{}"
	}
	return string(str)
}

func EntityToJsonEscape(entity interface{}) string {
	str, err := json.Marshal(entity)

	buffer := new(bytes.Buffer)
	json.HTMLEscape(buffer, str)

	if err != nil {
		return "{}"
	}
	return string(str)
}

func JsonToEntity(jsonIn string, entity interface{}) {
	err := json.Unmarshal([]byte(jsonIn), entity)

	if err != nil {
		entity = nil
	}
}

func Guid() string {
	return uuid.New().String()
}

func ConvertStringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

/*
2006-01-02T15:04:05-0700
2020-02-27T07:54:46.536-03:00
2020-03-09T14:06:48.903-03:00

ANSIC       = "Mon Jan _2 15:04:05 2006"
UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
RFC822      = "02 Jan 06 15:04 MST"
RFC822Z     = "02 Jan 06 15:04 -0700"
RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700"
RFC3339     = "2006-01-02T15:04:05Z07:00"
RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
Kitchen     = "3:04PM"
// Handy time stamps.
Stamp      = "Jan _2 15:04:05"
StampMilli = "Jan _2 15:04:05.000"
StampMicro = "Jan _2 15:04:05.000000"
StampNano  = "Jan _2 15:04:05.000000000"
*/

func StringToTime(layout, str string) time.Time {
	t, err := time.Parse(layout, str)
	if err != nil {
		return time.Time{}
	}
	return t
}
