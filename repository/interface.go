package repository

import "final/entities"

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
