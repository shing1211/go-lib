package string

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	log "github.com/sirupsen/logrus"
)

func StringToInt64(s string) (int64, error) {
	numero, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		log.Warn(err)
		return 0, err
	}
	return numero, err
}

func StringTouInt64(s string) (uint64, error) {
	numero, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		log.Warn(err)
		return 0, err
	}
	return numero, err
}

func StringToFloat64(s string) (float64, error) {
	numero, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Warn(err)
		return 0, err
	}
	return numero, err
}

func ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	}
	return ""
}

func ToStringKey(values ...interface{}) string {
	results := make([]string, len(values))

	for idx, value := range values {
		if valuer, ok := value.(driver.Valuer); ok {
			value, _ = valuer.Value()
		}

		switch v := value.(type) {
		case string:
			results[idx] = v
		case []byte:
			results[idx] = string(v)
		case uint:
			results[idx] = strconv.FormatUint(uint64(v), 10)
		default:
			results[idx] = fmt.Sprint(reflect.Indirect(reflect.ValueOf(v)).Interface())
		}
	}

	return strings.Join(results, "_")
}

func IsValidDBNameChar(c rune) bool {
	return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '.' && c != '*' && c != '_' && c != '$' && c != '@'
}
