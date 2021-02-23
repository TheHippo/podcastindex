package podcastindex

import (
	"fmt"
	"time"
)

func addMax(max uint) string {
	if max != 0 {
		return fmt.Sprintf("&max=%d", max)
	}
	return ""
}

func addClean(clean bool) string {
	if clean {
		return "&clean"
	}
	return ""
}

func addTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return fmt.Sprintf("&since=%d", t.Unix())
}
