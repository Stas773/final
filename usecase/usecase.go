package usecase

import (
	"final/entities"
	"final/repository"
)

type BuilderStruct struct {
	interSMS      repository.SMSWork
	interMMS      repository.MMSWork
	interVoice    repository.VoiceWork
	interEmail    repository.EmailWork
	interBilling  repository.BillingWork
	interSupport  repository.SupportWork
	interIncident repository.IncidentWork
	interResult   repository.ResultWork
}

func (b *BuilderStruct) SMSReader() [][]entities.SMSData {
	return b.interSMS.SMSReader()
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
