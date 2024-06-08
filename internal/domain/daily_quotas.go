package domain

import "time"

type DailyQuotasDto struct {
	QuotaID             int64
	AccountID           int64
	Date                time.Time
	TotalQuota          int64
	TotalQuotaUnlimited *int64
	SwipeCount          int
}
