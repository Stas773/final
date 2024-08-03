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
	MMSReader(l *logrus.Logger) ([][]entities.MMSData, error)
}
type VoiceWork interface {
	VoiceReader(l *logrus.Logger) []entities.VoiceData
}
type EmailWork interface {
	EmailReader(l *logrus.Logger) map[string][][]entities.EmailData
}
type BillingWork interface {
	BillingReader(l *logrus.Logger) entities.BillingData
}
type SupportWork interface {
	SupportReader(l *logrus.Logger) ([]int, error)
}
type IncidentWork interface {
	IncidentReader(l *logrus.Logger) ([]entities.IncidentData, error)
}
type ResultWork interface {
	ResultReader(l *logrus.Logger) (string, error)
}

func (b *BuilderStruct) SMSReader(l *logrus.Logger) [][]entities.SMSData {
	return b.interSMS.SMSReader(l)
}
func (b *BuilderStruct) MMSReader(l *logrus.Logger) ([][]entities.MMSData, error) {
	return b.interMMS.MMSReader(l)
}
func (b *BuilderStruct) VoiceReader(l *logrus.Logger) []entities.VoiceData {
	return b.interVoice.VoiceReader(l)
}
func (b *BuilderStruct) EmailReader(l *logrus.Logger) map[string][][]entities.EmailData {
	return b.interEmail.EmailReader(l)
}
func (b *BuilderStruct) BillingReader(l *logrus.Logger) entities.BillingData {
	return b.interBilling.BillingReader(l)
}
func (b *BuilderStruct) SupportReader(l *logrus.Logger) ([]int, error) {
	return b.interSupport.SupportReader(l)
}
func (b *BuilderStruct) IncidentReader(l *logrus.Logger) ([]entities.IncidentData, error) {
	return b.interIncident.IncidentReader(l)
}
func (b *BuilderStruct) ResultReader(l *logrus.Logger) (string, error) {
	return b.interResult.ResultReader(l)
}
