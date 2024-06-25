package models

import "final/entities"

type ResultSet struct {
	SMS       [][]entities.SMSData              `json:"sms"`
	MMS       [][]entities.MMSData              `json:"mms"`
	VoiceCall []entities.VoiceData              `json:"voice_call"`
	Email     map[string][][]entities.EmailData `json:"email"`
	Billing   entities.BillingData              `json:"billing"`
	Support   []int                             `json:"support"`
	Incidents []entities.IncidentData           `json:"incident"`
}
