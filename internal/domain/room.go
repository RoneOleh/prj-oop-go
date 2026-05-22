package domain

import "time"

type Room struct {
	Id          uint64
	\
	Name        string
	Description *string,
	City        string
	Address     string
	Lat         float64
	Lon         float64
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}
