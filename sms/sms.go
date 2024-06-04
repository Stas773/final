package sms

import (
	"bufio"
	"final/entities"
	"final/logger"
	"os"
	"sort"
	"strings"

	"github.com/biter777/countries"
)

const (
	Provider1 = "Topolo"
	Provider2 = "Rond"
	Provider3 = "Kildy"
)

type SMSStruct struct {
}

func (ss *SMSStruct) SMSReader() [][]entities.SMSData {
	var dataStructs []entities.SMSData
	var data []string
	var smsData entities.SMSData

	file, err := os.Open("simulator/sms.data")
	if err != nil {
		logger.Logger.Panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = strings.Split((line), ";")
		if len(data) == 4 {
			c := countries.ByName(data[0])
			if c.Alpha2() == data[0] && (data[3] == Provider1 || data[3] == Provider2 || data[3] == Provider3) {
				for i, v := range data {
					switch i {
					case 0:
						smsData.Country = c.String()
					case 1:
						smsData.Bandwidth = v
					case 2:
						smsData.ResponseTime = v
					case 3:
						smsData.Provider = v
					}
				}
				dataStructs = append(dataStructs, smsData)
			}
		}
	}
	var sortedByCountry []entities.SMSData
	var sortedByProvider []entities.SMSData

	sortedByCountry = append(sortedByCountry, dataStructs...)
	sort.Slice(sortedByCountry, func(i, j int) bool {
		return sortedByCountry[i].Country < sortedByCountry[j].Country
	})

	sortedByProvider = append(sortedByProvider, dataStructs...)
	sort.Slice(sortedByProvider, func(i, j int) bool {
		return sortedByProvider[i].Provider < sortedByProvider[j].Provider
	})

	var sortedData [][]entities.SMSData
	sortedData = append(sortedData, sortedByCountry)
	sortedData = append(sortedData, sortedByProvider)
	return sortedData
}
