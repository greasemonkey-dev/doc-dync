package entities

import (
	"strings"
)

// FilterCriteria represents a filter with category and value
type FilterCriteria struct {
	Category string `json:"category"`
	Value    string `json:"value"`
}

// AvailabilityCriteria defines filter criteria specifically for availability
type AvailabilityCriteria struct {
	DateTime int64 `json:"dateTime"`
}

type ScoreCriteria struct {
	MinScore float64 `json:"minScore"`
}

// Filter interface defines the Match method for filtering logic
type Filter interface {
	Match(p Provider) bool
}

// AvailabilityFilter implements the Filter interface for availability filtering
type AvailabilityFilter struct {
	Criteria AvailabilityCriteria
}

func (f *AvailabilityFilter) Match(p Provider) bool {
	return isAvailable(p.AvailableDates, f.Criteria.DateTime)
}

func isAvailable(dates []AvailableDate, dateTime int64) bool {
	for _, d := range dates {
		if dateTime >= d.From && dateTime <= d.To {
			return true
			break
		}
	}
	return false
}

// NameFilter implements Filter for name filtering
type NameFilter struct {
	Criteria FilterCriteria
}

func (f *NameFilter) Match(p Provider) bool {
	return strings.EqualFold(strings.TrimSpace(p.Name), strings.TrimSpace(f.Criteria.Value))

}

type ScoreFilter struct {
	Criteria ScoreCriteria
}

func (f *ScoreFilter) Match(p Provider) bool {
	return f.Criteria.MinScore <= p.Score
}

// SpecialtyFilter implements Filter for specialty filtering
type SpecialtyFilter struct {
	Criteria FilterCriteria
}

func (f *SpecialtyFilter) Match(p Provider) bool {
	for _, s := range p.Specialties {
		if strings.EqualFold(strings.TrimSpace(s), strings.TrimSpace(f.Criteria.Value)) {
			return true
		}

	}
	return false
}
