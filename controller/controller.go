package controller

import (
	"final/entities"
	"final/usecase"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	usecase usecase.BuilderStruct
}

func (h *Handler) SMSReader(l *logrus.Logger) [][]entities.SMSData {
	return h.usecase.SMSReader(l)
}
func (h *Handler) MMSReader(l *logrus.Logger) ([][]entities.MMSData, error) {
	return h.usecase.MMSReader(l)
}
func (h *Handler) VoiceReader(l *logrus.Logger) []entities.VoiceData {
	return h.usecase.VoiceReader(l)
}
func (h *Handler) EmailReader(l *logrus.Logger) map[string][][]entities.EmailData {
	return h.usecase.EmailReader(l)
}
func (h *Handler) BillingReader(l *logrus.Logger) entities.BillingData {
	return h.usecase.BillingReader(l)
}
func (h *Handler) SupportReader(l *logrus.Logger) ([]int, error) {
	return h.usecase.SupportReader(l)
}
func (h *Handler) IncidentReader(l *logrus.Logger) ([]entities.IncidentData, error) {
	return h.usecase.IncidentReader(l)
}
func (h *Handler) ResultReader(l *logrus.Logger) (string, error) {
	return h.usecase.ResultReader(l)
}
