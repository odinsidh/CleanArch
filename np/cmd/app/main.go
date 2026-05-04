package main

import (
	"fmt"

	"time"
)

func main() {
	a := time.Date(1900, time.April, 12, 12, 12, 12, 12, time.Local)
	xY := a.Year()
	xM := a.Month()
	xD := a.Day()

	fmt.Println(xY, xM, xD)

}
