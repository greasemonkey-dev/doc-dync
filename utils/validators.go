package utils

import (
	"slices"
	"strings"
)

func IsValidScore(minScore float64) bool {
	if minScore >= 0 && minScore <= 10 {
		return true
	}
	return false
}

func IsValidSpecialties(requestedSpecialty string) bool {
	existingSpecialties := []string{"neuropathy", "physiologist", "cardiologist", "internist", "pain Assistance", "Neonatal"}
	// Convert both lists to lowercase for case-insensitive comparison
	requestedSpecialtyLower := strings.ToLower(requestedSpecialty)

	// Check if the user sport is present in the valid sports list
	return slices.Contains(existingSpecialties, requestedSpecialtyLower)
}
