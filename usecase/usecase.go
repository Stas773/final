package usecase

import (
	"final/entities"

	"github.com/sirupsen/logrus"
)

type BuilderStruct struct {
	interSMS      SMSWork
	interMMS      MMSWork
	interVoice    VoiceWork
	interEmail    EmailWork
	interBilling  BillingWork
	interSupport  SupportWork
	interIncident IncidentWork
	interResult   ResultWork
}

type SMSWork interface {
	SMSReader(l *logrus.Logger) [][]entities.SMSData
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

func (b *BuilderStruct) SMSReader(l *logrus.Logger) [][]entities.SMSData {
	return b.interSMS.SMSReader(l)
}
func (b *BuilderStruct) MMSReader() ([][]entities.MMSData, error) {
	return b.interMMS.MMSReader()
}
func (b *BuilderStruct) VoiceReader() []entities.VoiceData {
	return b.interVoice.VoiceReader()
}
func (b *BuilderStruct) EmailReader() map[string][][]entities.EmailData {
	return b.interEmail.EmailReader()
}
func (b *BuilderStruct) BillingReader() entities.BillingData {
	return b.interBilling.BillingReader()
}
func (b *BuilderStruct) SupportReader() ([]int, error) {
	return b.interSupport.SupportReader()
}
func (b *BuilderStruct) IncidentReader() ([]entities.IncidentData, error) {
	return b.interIncident.IncidentReader()
}
func (b *BuilderStruct) ResultReader() (string, error) {
	return b.interResult.ResultReader()
}
