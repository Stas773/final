package usecase

import (
	"final/entities"
)

type SMSWork interface {
	SMSReader() [][]entities.SMSData
}
type MMSWork interface {
	MMSReader() ([][]entities.MMSData, error)
}
type VoiceWork interface {
	VoiceReader() []entities.VoiceData
}
type EmailWork interface {
	EmailReader() map[string][][]entities.EmailData
}
type BillingWork interface {
	BillingReader() entities.BillingData
}
type SupportWork interface {
	SupportReader() ([]int, error)
}
type IncidentWork interface {
	IncidentReader() ([]entities.IncidentData, error)
}
type ResultWork interface {
	ResultReader() (string, error)
}

func SMSReader(w SMSWork) [][]entities.SMSData {
	return w.SMSReader()
}
func MMSReader(w MMSWork) ([][]entities.MMSData, error) {
	return w.MMSReader()
}
func VoiceReader(w VoiceWork) []entities.VoiceData {
	return w.VoiceReader()
}
func EmailReader(w EmailWork) map[string][][]entities.EmailData {
	return w.EmailReader()
}
func BillingReader(w BillingWork) entities.BillingData {
	return w.BillingReader()
}
func SupportReader(w SupportWork) ([]int, error) {
	return w.SupportReader()
}
func IncidentReader(w IncidentWork) ([]entities.IncidentData, error) {
	return w.IncidentReader()
}
func ResultReader(w ResultWork) (string, error) {
	return w.ResultReader()
}
