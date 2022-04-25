package validator

import (
	"math"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func Required(value string) bool {
	if value != "" {
		return false
	}
	return true
}

func Email(value string) bool {
	match, _ := regexp.MatchString(emailRegex, value)
	if match {
		return false
	}
	return true
}

func Image(r *http.Request, key, mimes string) bool {
	_, header, err := r.FormFile(key)
	if err != nil {
		return true
	}
	ext := filepath.Ext(header.Filename)[1:]
	mimesSlice := strings.Split(mimes, ",")
	for i := 0; i < len(mimesSlice); i++ {
		if mimesSlice[i] == ext {
			return false
		}
	}
	return true
}

func Numeric(value string) bool {
	_, err := strconv.Atoi(value)
	_, err2 := strconv.ParseFloat(value, 64)
	if err == nil || err2 == nil {
		return false
	}
	return true
}

func Url(value string) bool {
	match, _ := regexp.MatchString(urlRegex, value)
	if match {
		return false
	}
	return true
}

func MinMax(r *http.Request, key, Value, operator string, values []string) (bool, string) {
	msg := GetMessage(key, Value)
	IntValue, IntValueErr := strconv.Atoi(Value)
	CheckError(IntValueErr)
	FloatValue, FloatValueErr := strconv.ParseFloat(Value, 64)
	CheckError(FloatValueErr)
	_, header, errFile := r.FormFile(key)
	if errFile == nil {
		kilobyte := float64(header.Size) / 1024
		kilobyte = math.Trunc(kilobyte*10) / 10
		if operator == "min" {
			if kilobyte >= FloatValue {
				return false, ""
			}
			return true, msg.Min["file"]
		} else {
			if kilobyte <= FloatValue {
				return false, ""
			}
			return true, msg.Max["file"]
		}
	} else {
		value := r.FormValue(key)
		intValue, intValueErr := strconv.Atoi(value)
		floatValue, floatValueErr := strconv.ParseFloat(value, 64)
		if (intValueErr == nil || floatValueErr == nil) && InSlice(values, "numeric") {
			if operator == "min" {
				if intValue >= IntValue || floatValue >= FloatValue {
					return false, ""
				}
				return true, msg.Min["numeric"]
			} else {
				if intValue <= IntValue || floatValue <= FloatValue {
					return false, ""
				}
				return true, msg.Max["numeric"]
			}
		} else if value == "" {
			return false, ""
		} else {
			if operator == "min" {
				if len(value) >= IntValue {
					return false, ""
				}
				return true, msg.Min["string"]
			} else {
				if len(value) <= IntValue {
					return false, ""
				}
				return true, msg.Max["string"]
			}
		}
	}
}

func Same(value, sameValue string) bool {
	if value == sameValue {
		return false
	}
	return true
}
