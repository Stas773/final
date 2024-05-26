package resultsetmodels

import (
	"final/entities/billing/billingmodels"
	"final/entities/email/emailmodels"
	"final/entities/incident/incidentmodels"
	"final/entities/mms/mmsmodels"
	"final/entities/sms/smsmodels"
	"final/entities/voice/voicemodels"
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
