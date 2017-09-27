package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	match, err := regexp.MatchString("p([a-z]+)ch", "peach")

	if err != nil {
		log.Fatal("match err.")
	}

	fmt.Println(match)

}
