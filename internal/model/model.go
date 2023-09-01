package model

import (
	"database/sql"
	"time"
)

type HistoryLine struct {
	Slug       string       `db:"slug"`
	AddedAt    time.Time    `db:"added_at"`
	ExpireTime sql.NullTime `db:"time_of_expire"`
}

type ModifySegmentInfo struct {
	UserId        int64
	SlugsToAdd    []string
	SlugsToRemove []string
}

type SetExpireTimeInfo struct {
	UserId     int64     `db:"user_id"`
	Slug       string    `db:"slug"`
	ExpireTime time.Time `db:"time_of_expire"`
}

func (h *HistoryLine) ToStringArray() []string {
	var res []string
	res = append(res, h.GetSlugStr(), h.GetAddedAtStr(), h.GetExpireTimeStr())

	return res
}

func (h *HistoryLine) GetSlugStr() string {
	return h.Slug
}

func (h *HistoryLine) GetAddedAtStr() string {
	return h.AddedAt.String()
}

func (h *HistoryLine) GetExpireTimeStr() string {
	if h.ExpireTime.Valid {
		return h.ExpireTime.Time.String()
	}
	return " "
}
