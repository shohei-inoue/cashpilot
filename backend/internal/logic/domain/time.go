package domain

import "time"

// TimeFormat は API で使う日時フォーマット（RFC3339 互換）
const TimeFormat = "2006-01-02T15:04:05.999999Z07:00"

func formatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(TimeFormat)
}
