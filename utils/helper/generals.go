package helper

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
	"reflect"
	"strconv"
	"time"
)

type Pagination struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	Total       int `json:"total"`
	TotalPage   int `json:"total_page"`
}

func ReturnInvalidResponse(httpcode int, data, meta interface{}, messageClient, messageServer string) error {
	responseBody := map[string]interface{}{
		"data": data,
		"meta": meta,
		"status": map[string]interface{}{
			"code":           httpcode,
			"message_client": messageClient,
			"message_server": messageServer,
		},
	}

	return echo.NewHTTPError(httpcode, responseBody)
}

func ReturnResponse(c echo.Context, httpcode int, data, meta interface{}, messageClient, messageServer string) error {
	responseBody := map[string]interface{}{
		"data": data,
		"meta": meta,
		"status": map[string]interface{}{
			"code":           httpcode,
			"message_client": messageClient,
			"message_server": messageServer,
		},
	}

	return c.JSON(httpcode, responseBody)
}

func ValidateRequestFormData(c echo.Context, rules govalidator.MapData) (i interface{}) {
	opts := govalidator.Options{
		Request: c.Request(),
		Rules:   rules,
	}

	v := govalidator.New(opts)
	mappedError := v.Validate()
	if len(mappedError) > 0 {
		i = mappedError
	}

	return i
}

func StartDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func EndDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 0, 0, t.Location())
}

// ToInt return casts of an interface to an int type.
func ToInt(i interface{}) int {
	v, _ := toInt(i)
	return v
}

// toInt return casts of an interface to an int type.
func toInt(i interface{}) (int, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int:
		return s, nil
	case int64:
		return int(s), nil
	case int32:
		return int(s), nil
	case int16:
		return int(s), nil
	case int8:
		return int(s), nil
	case uint:
		return int(s), nil
	case uint64:
		return int(s), nil
	case uint32:
		return int(s), nil
	case uint16:
		return int(s), nil
	case uint8:
		return int(s), nil
	case float64:
		return int(s), nil
	case float32:
		return int(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	}
}

func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// ToStringDate return casts of an string time intto a time type and error if failed.
func ToStringDate(i string) (time.Time, error) {
	v, err := stringToDate(i)
	return v, err
}

func stringToDate(s string) (time.Time, error) {
	return parseDateWith(s, []string{
		time.RFC3339,
		"2006-01-02T15:04:05", // iso8601 without timezone
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC850,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		"2006-01-02 15:04:05.999999999 -0700 MST", // Time.String()
		"2006-01-02",
		"02 Jan 2006",
		"2006-01-02T15:04:05-0700", // RFC3339 without timezone hh:mm colon
		"2006-01-02 15:04:05 -07:00",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05Z07:00", // RFC3339 without T
		"2006-01-02 15:04:05Z0700",  // RFC3339 without T or timezone hh:mm colon
		"2006-01-02 15:04:05",
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	})
}

func parseDateWith(s string, dates []string) (d time.Time, e error) {
	for _, dateType := range dates {
		if d, e = time.Parse(dateType, s); e == nil {
			return
		}
	}
	return d, fmt.Errorf("unable to parse date: %s", s)
}

func SumInt(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func SumFloat64(array []float64) float64 {
	result := 0.0
	for _, v := range array {
		result += v
	}
	return result
}
