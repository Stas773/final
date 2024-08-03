package controller

import (
	"final/entities"
	"final/usecase"

	"github.com/sirupsen/logrus"
)

var (
	SMSHandler      usecase.SMSWork      = &usecase.BuilderStruct{}
	MMSHandler      usecase.MMSWork      = &usecase.BuilderStruct{}
	VoiceHandler    usecase.VoiceWork    = &usecase.BuilderStruct{}
	EmailHandler    usecase.EmailWork    = &usecase.BuilderStruct{}
	BillingHandler  usecase.BillingWork  = &usecase.BuilderStruct{}
	SupportHandler  usecase.SupportWork  = &usecase.BuilderStruct{}
	IncidenrHandler usecase.IncidentWork = &usecase.BuilderStruct{}
)

func SMSReader(l *logrus.Logger) [][]entities.SMSData {
	data := SMSHandler.SMSReader(l)
	return data
}

func MMSReader(l *logrus.Logger) [][]entities.MMSData {
	data, err := MMSHandler.MMSReader(l)
	if err != nil {
		l.Error(err)
		return nil
	}
	if data == nil {
		l.Warn("MMS data is empty:", data)
		return nil
	}
	return data
}

func EmailReader(l *logrus.Logger) map[string][][]entities.EmailData {
	data := EmailHandler.EmailReader(l)
	return data
}

func VoiceReader(l *logrus.Logger) []entities.VoiceData {
	data := VoiceHandler.VoiceReader(l)
	return data
}

func BillingReader(l *logrus.Logger) entities.BillingData {
	data := BillingHandler.BillingReader(l)
	return data
}

func SupportReader(l *logrus.Logger) []int {
	data, err := SupportHandler.SupportReader(l)
	if err != nil {
		l.Error(err)
		return nil
	}
	if data == nil {
		l.Warn("Support data is empty:", data)
		return nil
	}
	return data
}

func IncidentReader(l *logrus.Logger) []entities.IncidentData {
	data, err := IncidenrHandler.IncidentReader(l)
	if err != nil {
		l.Error(err)
		return nil
	}
	if data == nil {
		l.Warn("Incident data is empty:", data)
		return nil
	}
	return data
}
