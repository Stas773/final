package incident

import (
	"encoding/json"
	"final/entities"
	"final/logger"
	"io"
	"net/http"
	"sort"
)

type IncidentStract struct {
}

func (is *IncidentStract) IncidentReader() ([]entities.IncidentData, error) {
	var result []entities.IncidentData
	resp, err := http.Get(incidentURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Logger.Error("can't connect, StatusCode:", resp.StatusCode)
		return nil, nil
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Status < result[j].Status
	})
	return result, nil
}
