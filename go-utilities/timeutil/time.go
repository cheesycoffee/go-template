package timeutil

import (
	"time"
)

const (
	// TimeZoneAsia constanta
	TimeZoneAsia = "Asia/Jakarta"
	// TokenClaimKey const
	TokenClaimKey = "tokenClaim"

	// TimeFormatLogger const
	TimeFormatLogger = "2006/01/02 15:04:05"
)

// AsiaJakartaLocalTime location
var AsiaJakartaLocalTime *time.Location

func init() {
	AsiaJakartaLocalTime, _ = time.LoadLocation(TimeZoneAsia)
}
