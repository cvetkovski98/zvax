package model

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

type Slot struct {
	SlotID    *string   `json:"slotId,omitempty" redis:"slotId"`
	DateTime  time.Time `json:"dateTime" redis:"dateTime"`
	Location  string    `json:"location,omitempty" redis:"location"`
	Available bool      `json:"available,omitempty" redis:"available"`
}

func (slot *Slot) ToMap() map[string]string {
	return map[string]string{
		"slotId":    *slot.SlotID,
		"dateTime":  slot.DateTime.Format(time.RFC3339),
		"location":  slot.Location,
		"available": fmt.Sprintf("%t", slot.Available),
	}
}

func NewSlotFromMap(h map[string]string) (*Slot, error) {
	slotId, ok := h["slotId"]
	if !ok {
		return nil, errors.New("slotId is not in hash")
	}
	dateTime, err := time.Parse(time.RFC3339, h["dateTime"])
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse dateTime")
	}
	location, ok := h["location"]
	if !ok {
		return nil, errors.New("location is not in hash")
	}
	available, err := strconv.ParseBool(h["available"])
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse available")
	}
	return &Slot{
		SlotID:    &slotId,
		DateTime:  dateTime,
		Location:  location,
		Available: available,
	}, nil
}

func NewSlotRedisId(location string, dateTime time.Time) string {
	dStr := dateTime.Format("2006-01-02")
	tStr := dateTime.Format("15-04")
	return fmt.Sprintf("slot:%s:%s:%s", location, dStr, tStr)
}
