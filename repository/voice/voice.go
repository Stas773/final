package voice

import (
	"bufio"
	"final/entities"
	"os"
	"strconv"
	"strings"

	"github.com/biter777/countries"
	"github.com/sirupsen/logrus"
)

const (
	Provider1 = "TransparentCalls"
	Provider2 = "E-Voice"
	Provider3 = "JustPhone"
)

type VoiceStruct struct {
}

func (vs *VoiceStruct) VoiceReader(l *logrus.Logger) []entities.VoiceData {
	var dataStructs []entities.VoiceData
	var data []string
	var voiceData entities.VoiceData

	file, err := os.Open(fileName)
	if err != nil {
		l.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = strings.Split((line), ";")
		if len(data) == 8 {
			c := countries.ByName(data[0])
			if c.Alpha2() == data[0] && (data[3] == Provider1 || data[3] == Provider2 || data[3] == Provider3) {
				for i, v := range data {
					switch i {
					case 0:
						voiceData.Country = v
					case 1:
						voiceData.Bandwidth = v
					case 2:
						voiceData.ResponseTime = v
					case 3:
						voiceData.Provider = v
					case 4:
						vFloat, err := strconv.ParseFloat(v, 32)
						if err != nil {
							l.Error("can't convert string to float: ", err)
						}
						voiceData.ConnectionStability = float32(vFloat)
					case 5:
						voiceData.TTFB, err = strconv.Atoi(v)
						if err != nil {
							l.Error("can't convert string to int: ", err)
						}
					case 6:
						voiceData.VoicePurity, err = strconv.Atoi(v)
						if err != nil {
							l.Error("can't convert string to int: ", err)
						}
					case 7:
						voiceData.MedianOfCallsTime, err = strconv.Atoi(v)
						if err != nil {
							l.Error("can't convert string to int: ", err)
						}
					}
				}
				dataStructs = append(dataStructs, voiceData)
			}
		}
	}
	return dataStructs
}
