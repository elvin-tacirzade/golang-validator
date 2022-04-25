package validator

import (
	"net/http"
	"strings"
)

func New(r *http.Request, data map[string][]string) []string {
	var messages []string
	for key, values := range data {
		msg := GetMessage(key, "")
		for _, value := range values {
			//required
			if value == "required" {
				check := Required(r.FormValue(key))
				messages = AppendMessage(check, msg.Required, messages)
			}
			//email
			if value == "email" {
				check := Email(r.FormValue(key))
				messages = AppendMessage(check, msg.Email, messages)
			}
			//image
			divide := strings.Split(value, ":")
			if divide[0] == "image" {
				msg = GetMessage(key, divide[1])
				check := Image(r, key, divide[1])
				messages = AppendMessage(check, msg.Image, messages)
			}
			//numeric
			if value == "numeric" {
				check := Numeric(r.FormValue(key))
				messages = AppendMessage(check, msg.Numeric, messages)
			}
			//string
			if value == "string" {
				check := !Numeric(r.FormValue(key))
				messages = AppendMessage(check, msg.Str, messages)
			}
			//url
			if value == "url" {
				check := Url(r.FormValue(key))
				messages = AppendMessage(check, msg.Url, messages)
			}
			//min-max
			//divide = strings.Split(value, ":")
			if divide[0] == "min" || divide[0] == "max" {
				check, msgMinMax := MinMax(r, key, divide[1], divide[0], values)
				messages = AppendMessage(check, msgMinMax, messages)
			}
			//same
			if divide[0] == "same" {
				msg = GetMessage(key, divide[1])
				check := Same(r.FormValue(key), r.FormValue(divide[1]))
				messages = AppendMessage(check, msg.Same, messages)
			}
		}
	}
	return messages
}
