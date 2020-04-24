package main

import (
	"fmt"
	"time"
)

func main() {
	mydate, _ := time.Parse("20060102", "20200421")
	fmt.Println("time:", mydate.In(time.Local).Format("January 02, 2006 (MST)"), "-- specify Local time zone")
	fmt.Println("time:", mydate.Format("January 02, 2006 (MST)"), "-- defaults to UTC")
}
