package algorithm

import (
	"fmt"
	"log"
)

var timeZone = map[string]int {
	"UTC": 0*60*60,
	"EST": -5*60*60,
	"CST": -6*60*60,
	"MST": -7*60*60,
	"PST": -8*60*60,
}

func attend() {
	offset := timeZone["EST"]
	fmt.Println(offset)

	var seconds int
	var ok bool
	seconds, ok = timeZone["NIL"]
	fmt.Println(seconds)
	fmt.Println(ok)
}

func offset(tz string) int {
	if seconds, ok := timeZone[tz]; ok {
		return seconds
	}
	log.Println("unknown time zone:", tz)
	return 0
}