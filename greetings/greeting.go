package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randFormat(), name)
	return message, nil
}

func HelloEveryOne(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randFormat() string {
	format := []string {
		"Hello %v, Welcome !",
		"Welcome %v, come on come on!",
		"Hi %v, Welcome to hortor !",
	}
	return format[rand.Intn(len(format))]
}


