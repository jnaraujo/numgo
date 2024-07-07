package internal

import "log"

func Assert(expected bool, msg string) {
	if !expected {
		log.Fatal(msg)
	}
}
