package entities

type Provider struct {
	Name           string          `json:"name" `
	Score          float64         `json:"score"`
	Specialties    []string        `json:"specialties"`
	AvailableDates []AvailableDate `json:"availableDates"`
}
type AvailableDate struct {
	From int64 `json:"from"`
	To   int64 `json:"to"`
}
type ProviderRequest struct {
	Specialty string  `json:"specialty" binding:"required"`
	Date      int64   `json:"date" binding:"required"`
	MinScore  float64 `json:"minScore" binding:"required"`
}
