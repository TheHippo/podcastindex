package podcastindex

import (
	"fmt"
	"strings"
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

func addFilter(name string, filter []string) string {
	if len(filter) == 0 {
		return ""
	}
	return fmt.Sprintf("&%s=%s", name, strings.Join(filter, ","))
}
