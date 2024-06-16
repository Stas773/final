package mms

import (
	"encoding/json"
	"final/entities"
	"io"
	"net/http"
	"sort"

	"github.com/biter777/countries"
)

type MMSStract struct {
}

func (ms *MMSStract) MMSReader() ([][]entities.MMSData, error) {
	var result []entities.MMSData
	resp, err := http.Get(mmsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	filteredData := filterInvalidData(result)

	var sortedByCountry []entities.MMSData
	var sortedByProvider []entities.MMSData

	sortedByCountry = append(sortedByCountry, filteredData...)
	sort.Slice(sortedByCountry, func(i, j int) bool {
		return sortedByCountry[i].Country < sortedByCountry[j].Country
	})

	sortedByProvider = append(sortedByProvider, filteredData...)
	sort.Slice(sortedByProvider, func(i, j int) bool {
		return sortedByProvider[i].Provider < sortedByProvider[j].Provider
	})

	var sortedData [][]entities.MMSData
	sortedData = append(sortedData, sortedByCountry)
	sortedData = append(sortedData, sortedByProvider)

	return sortedData, nil
}

func filterInvalidData(data []entities.MMSData) []entities.MMSData {
	validCountries := make(map[string]bool)
	for _, country := range countries.All() {
		validCountries[country.Alpha2()] = true
	}

	validProviders := map[string]bool{
		"Topolo": true,
		"Rond":   true,
		"Kildy":  true,
	}

	var result []entities.MMSData
	for _, v := range data {
		if validCountries[v.Country] && validProviders[v.Provider] {
			all := countries.ByName(v.Country)
			v.Country = all.String()
			result = append(result, v)
		}
	}
	return result
}
