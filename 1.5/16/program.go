package main

import "fmt"

func main() {
	var (
		angle          int
		hours, minutes int
		total_minutes  float32
		koeff          float32 = 0.5 // 360 degrees is 24*60, one degree is 0.5 minutes
	)
	fmt.Scan(&angle)
	angle = angle % 360
	total_minutes = float32(angle) / koeff
	hours = int(total_minutes) / 60
	minutes = int(total_minutes) - hours*60
	fmt.Println("It is", hours, "hours", minutes, "minutes.")

}
