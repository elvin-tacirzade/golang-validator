package validator

import (
	"net/http"
	"strings"
)

func New(r *http.Request, data map[string][]string) []string {
	var messages []string
	for key, values := range data {
		msg := getMessage(key, "")
		for _, value := range values {
			//required
			if value == "required" {
				check := required(r.FormValue(key))
				messages = appendMessage(check, msg.Required, messages)
			}
			//email
			if value == "email" {
				check := email(r.FormValue(key))
				messages = appendMessage(check, msg.Email, messages)
			}
			//image
			divide := strings.Split(value, ":")
			if divide[0] == "image" {
				msg = getMessage(key, divide[1])
				check := image(r, key, divide[1])
				messages = appendMessage(check, msg.Image, messages)
			}
			//numeric
			if value == "numeric" {
				check := numeric(r.FormValue(key))
				messages = appendMessage(check, msg.Numeric, messages)
			}
			//string
			if value == "string" {
				check := !numeric(r.FormValue(key))
				messages = appendMessage(check, msg.Str, messages)
			}
			//url
			if value == "url" {
				check := url(r.FormValue(key))
				messages = appendMessage(check, msg.Url, messages)
			}
			//min-max
			//divide = strings.Split(value, ":")
			if divide[0] == "min" || divide[0] == "max" {
				check, msgMinMax := minMax(r, key, divide[1], divide[0], values)
				messages = appendMessage(check, msgMinMax, messages)
			}
			//same
			if divide[0] == "same" {
				msg = getMessage(key, divide[1])
				check := same(r.FormValue(key), r.FormValue(divide[1]))
				messages = appendMessage(check, msg.Same, messages)
			}
		}
	}
	return messages
}
