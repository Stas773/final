package email

import (
	"bufio"
	"final/entities"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/biter777/countries"
	"github.com/sirupsen/logrus"
)

const (
	Provider1  = "Gmail"
	Provider2  = "Yahoo"
	Provider3  = "Hotmail"
	Provider4  = "MSN"
	Provider5  = "Orange"
	Provider6  = "Comcast"
	Provider7  = "AOL"
	Provider8  = "Live"
	Provider9  = "RediffMail"
	Provider10 = "GMX"
	Provider11 = "Protonmail"
	Provider12 = "Yandex"
	Provider13 = "Mail.ru"
)

type EmailStruct struct {
}

func (es *EmailStruct) EmailReader(l *logrus.Logger) map[string][][]entities.EmailData {
	var dataStructs []entities.EmailData
	var data []string
	var emailData entities.EmailData
	emailMap := make(map[string][][]entities.EmailData)

	file, err := os.Open(fileName)
	if err != nil {
		l.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = strings.Split((line), ";")
		if len(data) == 3 {
			c := countries.ByName(data[0])
			if c.Alpha2() == data[0] && (data[1] == Provider1 || data[1] == Provider2 || data[1] == Provider3 || data[1] == Provider4 || data[1] == Provider5 || data[1] == Provider6 || data[1] == Provider7 || data[1] == Provider8 || data[1] == Provider9 || data[1] == Provider10 || data[1] == Provider11 || data[1] == Provider12 || data[1] == Provider13) {
				for i, v := range data {
					switch i {
					case 0:
						emailData.Country = v
					case 1:
						emailData.Provider = v
					case 2:
						emailData.DeliveryTime, err = strconv.Atoi(v)
						if err != nil {
							l.Error("can't convert string to int: ", err)
						}
					}
				}
				dataStructs = append(dataStructs, emailData)
			}
		}
	}

	grouped := make(map[string][]entities.EmailData)
	for _, d := range dataStructs {
		grouped[d.Country] = append(grouped[d.Country], d)
	}

	for country, emails := range grouped {
		sort.Slice(emails, func(i, j int) bool {
			return emails[i].DeliveryTime < emails[j].DeliveryTime
		})
		fastest := emails[:3]
		slowest := emails[len(emails)-3:]

		emailMap[country] = [][]entities.EmailData{fastest, slowest}
	}
	return emailMap
}
