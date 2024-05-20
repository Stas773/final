package resultsetmodels

import (
	"final/billing/billingmodels"
	"final/email/emailmodels"
	"final/incident/incidentmodels"
	"final/mms/mmsmodels"
	"final/sms/smsmodels"
	"final/voice/voicemodels"
)

type ResultSet struct {
	SMS       [][]smsmodels.SMSData                `json:"sms"`
	MMS       [][]mmsmodels.MMSData                `json:"mms"`
	VoiceCall []voicemodels.VoiceData              `json:"voice_call"`
	Email     map[string][][]emailmodels.EmailData `json:"email"`
	Billing   billingmodels.BillingData            `json:"billing"`
	Support   []int                                `json:"support"`
	Incidents []incidentmodels.IncidentData        `json:"incident"`
}
