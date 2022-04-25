package validator

import (
	"fmt"
	"strings"
)

func AppendMessage(check bool, msg string, messages []string) []string {
	if check {
		messages = append(messages, msg)
	}
	return messages
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func InSlice(s []string, v string) bool {
	for _, value := range s {
		d := strings.Split(value, ":")
		if v == d[0] {
			return true
		}
	}
	return false
}
