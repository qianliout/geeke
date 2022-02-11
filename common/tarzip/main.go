package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	t := parseTime("2021-03-10 09:24:27")

	fmt.Println(t.String())
}

func parseTime(t string) time.Time {
	t = strings.Replace(t, ": ", ":", -1)
	ts := strings.Split(t, ".")
	ti, err := time.ParseInLocation("2006-01-02 15:04:05", ts[0], time.Local)
	if err != nil {
		return time.Time{}
	}
	return ti.UTC()
}
