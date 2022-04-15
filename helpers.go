package validator

import "fmt"

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
