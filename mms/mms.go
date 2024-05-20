package mms

import (
	"encoding/json"
	"final/mms/mmsmodels"
	"io"
	"net/http"
	"sort"

	"github.com/biter777/countries"
)

type MMSStract struct {
}

func (ms *MMSStract) MMSReader() ([][]mmsmodels.MMSData, error) {
	var result []mmsmodels.MMSData
	resp, err := http.Get("http://127.0.0.1:8383/mms")
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

	var sortedByCountry []mmsmodels.MMSData
	var sortedByProvider []mmsmodels.MMSData

	sortedByCountry = append(sortedByCountry, filteredData...)
	sort.Slice(sortedByCountry, func(i, j int) bool {
		return sortedByCountry[i].Country < sortedByCountry[j].Country
	})

	sortedByProvider = append(sortedByProvider, filteredData...)
	sort.Slice(sortedByProvider, func(i, j int) bool {
		return sortedByProvider[i].Provider < sortedByProvider[j].Provider
	})

	var sortedData [][]mmsmodels.MMSData
	sortedData = append(sortedData, sortedByCountry)
	sortedData = append(sortedData, sortedByProvider)

	return sortedData, nil
}

func filterInvalidData(data []mmsmodels.MMSData) []mmsmodels.MMSData {
	validCountries := make(map[string]bool)
	for _, country := range countries.All() {
		validCountries[country.Alpha2()] = true
	}

	validProviders := map[string]bool{
		"Topolo": true,
		"Rond":   true,
		"Kildy":  true,
	}

	var result []mmsmodels.MMSData
	for _, v := range data {
		if validCountries[v.Country] && validProviders[v.Provider] {
			all := countries.ByName(v.Country)
			v.Country = all.String()
			result = append(result, v)
		}
	}
	return result
}
