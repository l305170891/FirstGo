package main

import (
	"fmt"
	"regexp"
	"log"
)

func main() {
	match, err := regexp.MatchString("p([a-z]+)ch", "peach")

	if err != nil{
		log.Fatal("match err.")
	}

	fmt.Println(match)

}
