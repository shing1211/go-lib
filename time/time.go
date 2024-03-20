package time

import (
	"fmt"
	"strconv"
	"time"
)

func Measure(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", name, elapsed)
	fmt.Println()
}

func ConvertUnixTime(unixTime string) time.Time {
	i, err := strconv.ParseInt(unixTime, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	tm := time.Unix(i, 0)
	return tm
}
