package models

import (
	"time"
)

type TimeInterval struct {
	Start time.Time `json:"start" bson:"start"`
	End   time.Time `json:"end" bson:"end"`
}

func (self *TimeInterval) DateIsBefore(other time.Time) bool {
	return other.Before(self.Start)
}

func (self *TimeInterval) DateIsAfter(other time.Time) bool {
	return other.After(self.End)
}

func (self *TimeInterval) DateIsDuring(other time.Time) bool {
	return other.After(self.Start) && other.Before(self.End)
}

func (self *TimeInterval) OtherIntervalIsBefore(other TimeInterval) bool {
	return other.End.Before(self.Start)
}

func (self *TimeInterval) OtherIntervalIsAfter(other TimeInterval) bool {
	return other.Start.After(self.End)
}

func (self *TimeInterval) OtherIntervalIsDuring(other TimeInterval) bool {
	return other.Start.After(self.Start) && other.End.Before(self.End)
}

func (self *TimeInterval) OtherIntervalOverlaps(other TimeInterval) bool {
	return (other.Start.Before(self.Start) && other.End.After(self.Start)) ||
		(other.Start.Before(self.End) && other.End.After(self.End)) ||
		(other.Start.Before(self.Start) && other.End.After(self.End))
}