package tf

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var captchaSeed = rand.NewSource(int64(time.Now().Nanosecond()))

func CaptchaNumber(w int) string {
	var max int64 = 1
	for i := 0; i < w; i++ {
		max *= 10
	}
	r := captchaSeed.Int63() % max
	return fmt.Sprintf("%0"+strconv.Itoa(w)+"d", r)
}
