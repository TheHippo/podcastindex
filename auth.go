package podcastindex

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func generateAuthorizationHeader(key, secret string, now time.Time) string {
	raw := fmt.Sprintf("%s%s%d", key, secret, now.Unix())
	return fmt.Sprintf("%x", sha1.Sum([]byte(raw)))
}
