package main

import "fmt"

func main() {
	S := "Sat"
	k := 23
	days := make(map[string]int)
	days["Mon"] = 0
	days["Tue"] = 1
	days["Wed"] = 2
	days["Thu"] = 3
	days["Fri"] = 4
	days["Sat"] = 5
	days["Sun"] = 6

	dayNo := make(map[int]string)
	dayNo[0] = "Mon"
	dayNo[1] = "Tue"
	dayNo[2] = "Wed"
	dayNo[3] = "Thu"
	dayNo[4] = "Fri"
	dayNo[5] = "Sat"
	dayNo[6] = "Sun"

	fmt.Printf("The day after %d days from %s is %s", k, S, dayNo[(days[S]+k)%7])
}
