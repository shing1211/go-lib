package time

import (
	"fmt"
	"time"
)

func Measure(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", name, elapsed)
	fmt.Println()
}
