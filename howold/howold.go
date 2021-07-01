package main

import (
	"example.com/greetings"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	log.SetPrefix("SomePrefix:")
	log.SetFlags(0)
	for sum := 1 ; sum < 10; sum++ {
		//common()
	}
	hellos()
}

func hellos() {
	names := []string{"zhangsan", "lisi", "wangwu"}
	message, err := greetings.HelloEveryOne(names)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(message)
}

func common() {
	message, err := greetings.Hello("yao")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("name:" + message)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string {
		"Hi welcome %v",
		"Hello %v, I will come back",
		"Welcome to hortor %v, are you ok?",
	}
	return formats[rand.Intn(len(formats))]
}