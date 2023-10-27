package models

import "time"

type TimeStamp struct {
	FirstStock    time.Time
	RecentStock   time.Time
	NearestExpiry time.Time
}
