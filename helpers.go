package validator

import (
	"fmt"
	"strings"
)

func appendMessage(check bool, msg string, messages []string) []string {
	if check {
		messages = append(messages, msg)
	}
	return messages
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func inSlice(s []string, v string) bool {
	for _, value := range s {
		d := strings.Split(value, ":")
		if v == d[0] {
			return true
		}
	}
	return false
}
