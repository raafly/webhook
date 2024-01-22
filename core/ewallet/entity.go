package ewallet

import (
	"time"
)

type ewallet struct {
	ID 			string
	Status      string
	CustomerID  string
	ReferenceID string
	Created		time.Time
	Update		time.Time
}