package gigasecond

// import path for the time package from the standard library
import (
	"math"
	"time"
)

const testVersion = 4

// AddGigasecond increments the supplied time by 10^9 seconds
func AddGigasecond(t time.Time) time.Time {
	gigaSecond := time.Second * time.Duration(int64(math.Pow(10, 9)))
	return t.Add(gigaSecond)
}
