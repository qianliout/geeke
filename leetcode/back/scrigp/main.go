package main

import (
	"fmt"
	"regexp"
)

func main() {

	telephoneRegex := `^\d+$`
	phone := "+4000081323"

	matched, _ := regexp.MatchString(telephoneRegex, phone)
	fmt.Println(matched)
}
