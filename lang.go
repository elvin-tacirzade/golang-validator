package validator

type Lang struct {
	Required string
	Email    string
	Image    string
	Numeric  string
	Str      string
	Url      string
	Min      map[string]string
	Max      map[string]string
	Same     string
}

func GetMessage(name, values string) Lang {
	eng := Lang{
		Required: "The " + name + " field is required.",
		Email:    "The " + name + " must be a valid email address.",
		Image:    "The " + name + " must be a file of type: " + values + ".",
		Numeric:  "The " + name + " must be a number.",
		Str:      "The " + name + " must be a string.",
		Url:      "The " + name + " must be a valid URL.",
		Min: map[string]string{
			"numeric": "The " + name + " must be at least " + values + ".",
			"file":    "The " + name + " must be at least " + values + " kilobytes.",
			"string":  "The " + name + " must be at least " + values + " characters.",
		},
		Max: map[string]string{
			"numeric": "The " + name + " must not be greater than " + values + ".",
			"file":    "The " + name + " must not be greater than " + values + " kilobytes.",
			"string":  "The " + name + " must not be greater than " + values + " characters.",
		},
		Same: "The " + name + " and " + values + " must match.",
	}
	return eng
}
