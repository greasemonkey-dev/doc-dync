package sync_api

import (
	"doc-sync/entities"
	"encoding/json"
	"log"
	"sort"
)

func FilterProviders(jsonData []byte, req entities.ProviderRequest) ([]entities.Provider, error) {
	//TODO: consider using map
	var providers []entities.Provider
	var filteredProviders []entities.Provider

	err := json.Unmarshal(jsonData, &providers)
	if err != nil {
		log.Printf("Error unmarshalling JSON:", err)
		return nil, err
	}
	availabilityFilter := &entities.AvailabilityFilter{
		Criteria: entities.AvailabilityCriteria{DateTime: req.Date},
	}

	specialtyFilter := &entities.SpecialtyFilter{
		Criteria: entities.FilterCriteria{Category: "specialty", Value: req.Specialty},
	}
	scoreFilter := &entities.ScoreFilter{
		Criteria: entities.ScoreCriteria{MinScore: req.MinScore},
	}

	// Inside FilterProviders
	ch := make(chan []entities.Provider, len(providers))

	for _, p := range providers {
		go func(provider entities.Provider) {
			if availabilityFilter.Match(provider) && specialtyFilter.Match(provider) && scoreFilter.Match(provider) {
				ch <- []entities.Provider{provider}
			} else {
				ch <- nil
			}
		}(p)
	}

	for range providers {
		result := <-ch
		if result != nil {
			filteredProviders = append(filteredProviders, result[0])
		}
	}

	if len(filteredProviders) > 0 {
		sortedProviders, err := sortProviders(filteredProviders)
		if err != nil {
			log.Printf("error while sorting", err)
			return nil, err
		}
		return sortedProviders, err
	} else {
		return filteredProviders, nil
	}
}

func sortProviders(providers []entities.Provider) ([]entities.Provider, error) {
	sort.Slice(providers, func(i, j int) bool {
		return providers[i].Score > providers[j].Score
	})
	return providers, nil
}
