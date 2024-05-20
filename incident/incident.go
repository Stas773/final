package incident

import (
	"encoding/json"
	"final/incident/incidentmodels"
	"final/logger"
	"io"
	"net/http"
	"sort"
)

type IncidentStract struct {
}

func (is *IncidentStract) IncidentReader() ([]incidentmodels.IncidentData, error) {
	var result []incidentmodels.IncidentData
	resp, err := http.Get("http://127.0.0.1:8383/accendent")
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
