package support

import (
	"encoding/json"
	"final/entities/support/supportmodels"
	"final/logger"
	"io"
	"math"
	"net/http"
)

type SupportStract struct {
}

func (ms *SupportStract) SupportReader() ([]int, error) {
	var supportData []supportmodels.SupportData
	resp, err := http.Get("http://127.0.0.1:8383/support")
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
	err = json.Unmarshal(bodyBytes, &supportData)
	if err != nil {
		return nil, err
	}

	var result []int
	var timeToTicket float64
	var load int
	var tickets int
	for i, _ := range supportData {
		tickets += supportData[i].ActiveTickets
	}
	timeToTicket = math.Ceil(60.0 / 18.0 * float64(tickets))
	if timeToTicket < 9 {
		load = 1
	} else if timeToTicket > 16 {
		load = 3
	} else {
		load = 2
	}
	result = append(result, load, int(timeToTicket))
	return result, nil
}
