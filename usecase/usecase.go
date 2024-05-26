package usecase

import (
	"final/entities/billing/billingmodels"
	"final/entities/email/emailmodels"
	"final/entities/incident/incidentmodels"
	"final/entities/mms/mmsmodels"
	"final/entities/sms/smsmodels"
	"final/entities/voice/voicemodels"
)

type SMSWork interface {
	SMSReader() [][]smsmodels.SMSData
}
type MMSWork interface {
	MMSReader() ([][]mmsmodels.MMSData, error)
}
type VoiceWork interface {
	VoiceReader() []voicemodels.VoiceData
}
type EmailWork interface {
	EmailReader() map[string][][]emailmodels.EmailData
}
type BillingWork interface {
	BillingReader() billingmodels.BillingData
}
type SupportWork interface {
	SupportReader() ([]int, error)
}
type IncidentWork interface {
	IncidentReader() ([]incidentmodels.IncidentData, error)
}
type ResultWork interface {
	ResultReader() (string, error)
}

func SMSReader(w SMSWork) [][]smsmodels.SMSData {
	return w.SMSReader()
}
func MMSReader(w MMSWork) ([][]mmsmodels.MMSData, error) {
	return w.MMSReader()
}
func VoiceReader(w VoiceWork) []voicemodels.VoiceData {
	return w.VoiceReader()
}
func EmailReader(w EmailWork) map[string][][]emailmodels.EmailData {
	return w.EmailReader()
}
func BillingReader(w BillingWork) billingmodels.BillingData {
	return w.BillingReader()
}
func SupportReader(w SupportWork) ([]int, error) {
	return w.SupportReader()
}
func IncidentReader(w IncidentWork) ([]incidentmodels.IncidentData, error) {
	return w.IncidentReader()
}
func ResultReader(w ResultWork) (string, error) {
	return w.ResultReader()
}
