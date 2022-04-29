package validator

import (
	"math"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func required(value string) bool {
	if value != "" {
		return false
	}
	return true
}

func email(value string) bool {
	match, _ := regexp.MatchString(emailRegex, value)
	if match {
		return false
	}
	return true
}

func image(r *http.Request, key, mimes string) bool {
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

func numeric(value string) bool {
	_, err := strconv.Atoi(value)
	_, err2 := strconv.ParseFloat(value, 64)
	if err == nil || err2 == nil {
		return false
	}
	return true
}

func url(value string) bool {
	match, _ := regexp.MatchString(urlRegex, value)
	if match {
		return false
	}
	return true
}

func minMax(r *http.Request, key, Value, operator string, values []string) (bool, string) {
	msg := getMessage(key, Value)
	IntValue, IntValueErr := strconv.Atoi(Value)
	checkError(IntValueErr)
	FloatValue, FloatValueErr := strconv.ParseFloat(Value, 64)
	checkError(FloatValueErr)
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
		if (intValueErr == nil || floatValueErr == nil) && inSlice(values, "numeric") {
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

func same(value, sameValue string) bool {
	if value == sameValue {
		return false
	}
	return true
}
