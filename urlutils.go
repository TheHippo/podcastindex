package podcastindex

import "fmt"

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
